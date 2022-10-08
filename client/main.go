package main

import (
	"log"

	"client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	creds := insecure.NewCredentials()
	cc,err := grpc.Dial("localhost:8080",grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	
	defer cc.Close()
	
	activityClient := services.NewActivityClient(cc)
	activityService := services.NewActivityService(activityClient)


	data := services.Activity{
		Name: "GG",
		Description: 	"Find friends to eat dinner together.",
		ImageProfile:	"www.img.co/dummy.img",
		Type:			[]string {"restaurant","GG"},
		OwnerId: 		"a92dikjsiao92nfaoiw",
		Location:		"Siam park",
		MaxParticipant:	15,
		Participant:   	[]string {"a92dikjsiao92nfaoiw"},
		Date:			"2510-04-04T07:00:00.000Z",
		Duration:		3,
		ChatId:			"09jdao92ndnawndsak",
	}


	err = activityService.CreateActivity(data)

	if err != nil {
		log.Fatal(err)
	}
	
}
