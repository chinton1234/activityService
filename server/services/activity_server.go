package services

import (
	"context"
	"fmt"
	"server/configs"
	"server/models"

	"server/RPC"
	"time"

	// "github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var activityCollection *mongo.Collection = configs.GetCollection(configs.DB, "activitys")

// var validate = validator.New()

type activityServer struct {
}

func NewActivityServer() ActivityServer {
	return activityServer{}
}

func (activityServer) mustEmbedUnimplementedActivityServer() {}

func (activityServer) CreateActivity(ctx context.Context, req *ActivityForm) (*Response, error) {
	// var activity models.Activity

	// if validationErr := validate.Struct(&activity); validationErr != nil {
	// 	fmt.Println(validationErr)
	// 	fmt.Println("")
	// 	return nil, validationErr
	// }

	fmt.Println("create activity.")

	newAct := models.ActCreate{
		Name:           req.Name,
		Description:    req.Description,
		ActivityType:   req.ActivityType,
		ImageProfile:   req.ImageProfile,
		OwnerId:        req.OwnerId,
		Location:       req.Location,
		MaxParticipant: int(req.MaxParticipant),
		Date:           req.Date,
		Duration:       req.Duration,
		ChatId:         "",
	}

	owner := req.OwnerId

	result1, err := activityCollection.InsertOne(ctx, newAct)
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	data := result1.InsertedID.(primitive.ObjectID).Hex()

	var matchingfunction RPC.Export
	text := "create " + data + " " + owner
	fmt.Println(text)
	mId, err := matchingfunction.Matching(text)
	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	objID, err := primitive.ObjectIDFromHex(mId)
	if err != nil {
		panic(err)
	}

	update := bson.M{
		"matchingId": objID,
	}

	result, err := activityCollection.UpdateOne(ctx, bson.M{"_id": result1.InsertedID.(primitive.ObjectID)}, bson.M{"$set": update})

	if err != nil {
		fmt.Println(err)
		fmt.Println("")
		return nil, err
	}

	//get updated user details
	var updatedActivity models.Activity
	if result.MatchedCount == 1 {
		err := activityCollection.FindOne(ctx, bson.M{"_id": result1.InsertedID.(primitive.ObjectID)}).Decode(&updatedActivity)

		if err != nil {
			fmt.Println(err)
			fmt.Println("")
			return nil, err
		}
	}

	res := Response{
		Status:  200,
		Message: data,
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

		participantId := req.Participant.Hex()

		var one = Activity{
			ActivityId:     req.ID.Hex(),
			Name:           req.Name,
			Description:    req.Description,
			ActivityType:   req.ActivityType,
			ImageProfile:   req.ImageProfile,
			OwnerId:        req.OwnerId,
			Location:       req.Location,
			MaxParticipant: int64(req.MaxParticipant),
			Participant:    participantId,
			Date:           req.Date,
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

	participantId := activity.Participant.Hex()

	var data = Activity{
		ActivityId:     activity.ID.Hex(),
		Name:           activity.Name,
		Description:    activity.Description,
		ActivityType:   activity.ActivityType,
		ImageProfile:   activity.ImageProfile,
		OwnerId:        activity.OwnerId,
		Location:       activity.Location,
		MaxParticipant: int64(activity.MaxParticipant),
		Participant:    participantId,
		Date:           activity.Date,
		Duration:       activity.Duration,
		ChatId:         activity.ChatId,
	}

	return &data, nil
}

func (activityServer) EditActivity(ctx context.Context, req *ActivityEdit) (*Response, error) {

	activityId := req.ActivityId

	objId, _ := primitive.ObjectIDFromHex(activityId)

	update := bson.M{
		"name":           req.Name,
		"description":    req.Description,
		"activityType":   req.ActivityType,
		"imageProfile":   req.ImageProfile,
		"ownerid":        req.OwnerId,
		"location":       req.Location,
		"maxparticipant": int(req.MaxParticipant),
		"date":           req.Date,
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

	var activity models.Activity
	err := activityCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&activity)
	if err != nil {
		return nil, err
	}

	matchingId := activity.Participant.Hex()

	text := "delete " + matchingId
	fmt.Println(text)
	var matchingfunction RPC.Export
	data, err := matchingfunction.Matching(text)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

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
