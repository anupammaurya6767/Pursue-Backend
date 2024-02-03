package firebase_middleware

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// CountChatBotSession - Method to count the total number of chat bot sessions started by the users
func CountTotalChatBotSessions() (uint, error) {

	// Initialising the global variables for client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Getting the itreator for the db
	iter := client.Collection("Users").Documents(ctx)
	var totalChatBotSessionsCount uint = 0

	// Looping using the iterator 
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return 0, err
		}

		// Getting the data from the document
		chatBotSession := doc.Data()["DidStartChatBot"].(bool)

		if chatBotSession == true {
			totalChatBotSessionsCount++
		}
	}

	// If everything goes well, return the total count
	return totalChatBotSessionsCount, nil
}
