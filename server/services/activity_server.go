package services

import (
	"context"
	"fmt"
	"server/configs"
	"server/models"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var activityCollection *mongo.Collection = configs.GetCollection(configs.DB, "activitys")
var validate = validator.New()

type activityServer struct {
}

func NewActivityServer() ActivityServer {
	return activityServer{}
}

func (activityServer) mustEmbedUnimplementedActivityServer() {}

func (activityServer) CreateActivity(ctx context.Context, req *Activity) (*Response, error) {
	// var activity models.Activity

	// if validationErr := validate.Struct(&activity); validationErr != nil {
	// 	fmt.Println(validationErr)
	// 	fmt.Println("")
	// 	return nil, validationErr
	// }

	fmt.Println("create activity.")
	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.Date)
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	newUser := models.Activity{
		Name:           req.Name,
		Description:    req.Description,
		ImageProfile:   req.ImageProfile,
		Type:           req.Type,
		OwnerId:        req.OwnerId,
		Location:       req.Location,
		MaxParticipant: int(req.MaxParticipant),
		Participant:    req.Participant,
		Date:           time,
		Duration:       req.Duration,
		ChatId:         req.ChatId,
	}

	result, err := activityCollection.InsertOne(ctx, newUser)
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	ID := fmt.Sprintf("%v", result.InsertedID)
	res := Response{
		Status:  200,
		Message: ID,
	}

	fmt.Println("Complete.")
	fmt.Println("")
	return &res, nil
}

func (activityServer) GetActivitys(context.Context, *Empty) (*ActivityList, error) {
	fmt.Println("Get All activity.")

	var activitys []*Activity

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := activityCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	defer results.Close(ctx)
	//reading from the db in an optimal way
	for results.Next(ctx) {
		var req models.Activity
		if err = results.Decode(&req); err != nil {
			fmt.Println(err)
			fmt.Println("")
			return nil, err
		}

		var one = Activity{
			ActivityId:     req.ID.Hex(),
			Name:           req.Name,
			Description:    req.Description,
			ImageProfile:   req.ImageProfile,
			Type:           req.Type,
			OwnerId:        req.OwnerId,
			Location:       req.Location,
			MaxParticipant: int64(req.MaxParticipant),
			Participant:    req.Participant,
			Date:           req.Date.String(),
			Duration:       req.Duration,
			ChatId:         req.ChatId,
		}

		activitys = append(activitys, &one)
	}

	var data = ActivityList{
		Data: activitys,
	}
	fmt.Println("Complete.")
	return &data, nil

}

func (activityServer) GetActivity(ctx context.Context, req *ActivityId) (*Activity, error) {

	activityId := req.Id
	var activity models.Activity

	objId, _ := primitive.ObjectIDFromHex(activityId)

	err := activityCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&activity)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	var data = Activity{
		ActivityId:     activity.ID.Hex(),
		Name:           activity.Name,
		Description:    activity.Description,
		ImageProfile:   activity.ImageProfile,
		Type:           activity.Type,
		OwnerId:        activity.OwnerId,
		Location:       activity.Location,
		MaxParticipant: int64(activity.MaxParticipant),
		Participant:    activity.Participant,
		Date:           activity.Date.String(),
		Duration:       activity.Duration,
		ChatId:         activity.ChatId,
	}

	return &data, nil
}

func (activityServer) EditActivity(ctx context.Context, req *Activity) (*Response, error) {

	activityId := req.ActivityId

	objId, _ := primitive.ObjectIDFromHex(activityId)

	layout := "2006-01-02T15:04:05.000Z"
	time, err := time.Parse(layout, req.Date)
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	update := bson.M{
		"name":           req.Name,
		"description":    req.Description,
		"imageprofile":   req.ImageProfile,
		"type":           req.Type,
		"ownerid":        req.OwnerId,
		"location":       req.Location,
		"maxparticipant": int(req.MaxParticipant),
		"participant":    req.Participant,
		"date":           time,
		"duration":       req.Duration,
		"chatId":         req.ChatId,
	}

	result, err := activityCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	//get updated user details
	var updatedActivity models.Activity
	if result.MatchedCount == 1 {
		err := activityCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedActivity)

		if err != nil {
			fmt.Println(err)
			fmt.Println("")
			return nil, err
		}
	}

	res := Response{
		Status:  200,
		Message: "Success save activity information",
	}
	return &res, nil
}

func (activityServer) DeleteActivity(ctx context.Context, req *ActivityId) (*Response, error) {

	activityId := req.Id
	objId, _ := primitive.ObjectIDFromHex(activityId)

	result, err := activityCollection.DeleteOne(ctx, bson.M{"_id": objId})

	if err != nil {
		return nil, err
	}

	if result.DeletedCount < 1 {
		res := Response{
			Status:  404,
			Message: "User with id " + activityId + " not found",
		}
		return &res, nil
	}

	res := Response{
		Status:  200,
		Message: "User successfully deleted!",
	}
	return &res, nil

}
