package firebase_middleware

import (
	"log"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// RetrieveCareerDescription - Retrieve career description from firebase
func RetrieveCareerDescription(careerTitle string) map[string]interface{} {

	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading data form the db
	iter := client.Collection("CareerDescription").Documents(ctx)

	// Looping through the iterator
	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var title string = data.Data()["Name"].(string)

		if(title == careerTitle) {
			return data.Data()
		}
	}

	return nil 
}
