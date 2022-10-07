package services

import context "context"

type activityServer struct {
}

func NewActivityServer() ActivityServer {
	return activityServer{}
}

func (activityServer) mustEmbedUnimplementedActivityServer() {}

func (activityServer) CreateActivity(ctx context.Context, req *Activity) (*Response, error) {

	res := Response{
		Status:  12,
		Message: "GG",
	}
	return &res, nil
}

func (activityServer) GetActivitys(*Empty, Activity_GetActivitysServer) error {
	return nil
}

func (activityServer) GetActivity(ctx context.Context, req *ActivityId) (*Activity, error) {

	res := Activity{}
	return &res, nil
}

func (activityServer) EditActivity(ctx context.Context, req *Activity) (*Response, error) {
	res := Response{
		Status:  12,
		Message: "GG",
	}
	return &res, nil
}

func (activityServer) DeleteActivity(ctx context.Context, req *ActivityId) (*Response, error) {
	res := Response{
		Status:  12,
		Message: "GG",
	}
	return &res, nil
}
