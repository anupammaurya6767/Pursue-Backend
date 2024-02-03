package firebase_middleware

import (
	"log"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// ReadCareerDescriptionFromFirebase - Retreve all the Career Descriptions from Firebase
func ReadCareerDescriptionFromFirebase() []map[string]interface{} {
	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading data form the db
	iter := client.Collection("CareerDescription").Documents(ctx)

	// Array of Career Descriptions
	var careerDescriptions []map[string]interface{}

	// Looping through the iterator
	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		variable := data.Data()

		careerDescriptions = append(careerDescriptions, variable)
	}

	return careerDescriptions
}
