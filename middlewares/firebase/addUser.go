package firebase_middleware

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

// AddUserToFirebase Method to add a User to Firestore
func AddUserToFirebase(name string, isPaidUser bool, options []string, finalResult[]string, email string, phoneNumber int, didStartChatbot bool) (string, error) {

	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding data to the DB
	doc, _, err := client.Collection("Users").Add(ctx, map[string]interface{}{
		"First Name": name,
		"Email":      email,
		"Phone":      phoneNumber,
		"DidStartChatbot": didStartChatbot,
		"IsPaidUser": isPaidUser,
		"Options":    options,
		"FinalCareerOptions": finalResult,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, nil

}
