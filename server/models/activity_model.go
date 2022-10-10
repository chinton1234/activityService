package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Activity struct{
	ID    				primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name 				string `json:"activityName" validate:"required"`
	Description 		string  `json: "description"`
	ImageProfile 		string `json: "imageProfile"`
	Type 				[]string `json: "activityType"`
	OwnerId 			string `json: "ownerId" validate:"required"`
	Location			string `json: "location" validate:"required"`
	MaxParticipant		int `json: "maxParticipant" validate:"required"`
	Participant			[]string `json: "participant"`
	Date				time.Time `json: "date"`
	Duration			int64 `json: "duration"`
	ChatId   			string `json: "chatId"`
} 
// activityName   description   imageProfile   activityType   ownerId    location  maxParticipant   participant(list) date   duration  chatId    