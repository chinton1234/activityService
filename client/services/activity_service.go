package services

import (
	"context"
	"fmt"
)

type ActivityService interface {
}

type activityService struct {
	activityClient activityClient
}

func NewActivityService(activityClient activityClient) activityService {
	return activityService{activityClient}
}

func (base activityService) CreateActivity(req Activity) error {
	// req := Activity{
	// 	Name: name,
	// Description: 	"Find friends to eat dinner together.",
	// imageProfile:	"www.img.co/dummy.img",
	// activityType:			["restaurant","GG"],
	// ownerId: 		"a92dikjsiao92nfaoiw",
	// location:		"Siam park",
	// maxParticipant:	15,
	// participant:   ["a92dikjsiao92nfaoiw"],
	// date:			"2510-04-04T07:00:00.000Z",
	// duration:		"3 hours",
	// chatId:			"09jdao92ndnawndsak"
	// }

	res, err := base.activityClient.CreateActivity(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service: Create Activity\n")
	fmt.Printf("Request : %v\n",req.Name)
	fmt.Printf("Response: %v %v\n",res.Status,res.Message)
	return nil
}

// func (c *activityClient) GetActivitys(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Activity_GetActivitysClient, error) {

// }

// type Activity_GetActivitysClient interface {
// 	Recv() (*Activity, error)
// 	grpc.ClientStream
// }

// type activityGetActivitysClient struct {
// 	grpc.ClientStream
// }

// func (x *activityGetActivitysClient) Recv() (*Activity, error) {
// 	m := new(Activity)
// 	if err := x.ClientStream.RecvMsg(m); err != nBackground()
// 		return nil, err
// 	}
// 	return m, nil
// }

// func (c *activityClient) GetActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Activity, error) {

// }

// func (c *activityClient) EditActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {

// }

// func (c *activityClient) DeleteActivity(ctx context.Context, in *ActivityId, opts ...grpc.CallOption) (*Response, error) {

// }