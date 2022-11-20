package models

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActCreate struct{
	Name 				string `json:"activityName" validate:"required"`
	Description 		string  `json: "description"`
	ActivityType		[]string `json: "activityType"`
	ImageProfile 		string `json: "imageProfile"`
	OwnerId 			string `json: "ownerId" validate:"required"`
	Location			string `json: "location" validate:"required"`
	MaxParticipant		int `json: "maxParticipant" validate:"required"`
	Participant			primitive.ObjectID `bson:"matchingId" json:"matchingId,omitempty"`
	Date				string `json: "date"`
	Duration			float32 `json: "duration"`
	ChatId   			string `json: "chatId"`
} 




type Activity struct{
	ID    				primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name 				string `json:"activityName" validate:"required"`
	Description 		string  `json: "description"`
	ActivityType		[]string `json: "activityType"`
	ImageProfile 		string `json: "imageProfile"`
	OwnerId 			string `json: "ownerId" validate:"required"`
	Location			string `json: "location" validate:"required"`
	MaxParticipant		int `json: "maxParticipant" validate:"required"`
	Participant			primitive.ObjectID `bson:"matchingId" json:"matchingId,omitempty"`
	Date				string `json: "date"`
	Duration			float32 `json: "duration"`
	ChatId   			string `json: "chatId"`
} 
// activityName   description   imageProfile   activityType   ownerId    location  maxParticipant   participant(list) date   duration  chatId    



