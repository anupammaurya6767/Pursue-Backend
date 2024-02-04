package firebase_middleware

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// AddUserToFirebase Method to add a User to Firestore
func AddUserToFirebase(user types.FirebaseUser) (string, error) {
	// Initialising the client and context to interact with Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding data to the DB
	doc, _, err := client.Collection("Users").Add(ctx, map[string]interface{}{
		"Name":               user.Name,
		"Email":              user.Email,
		"PhoneNumber":        user.PhoneNumber,
		"DidStartChatbot":    user.DidStartChatbot,
		"IsPaidUser":         user.IsPaidUser,
		"Options":            user.Options,
		"FinalCareerOptions": user.FinalCareerOptions,
		"UserID":             user.UserID,
		"OrderID":            user.OrderID,
		"CustomerID":         user.CustomerID,
	})
	if err != nil {
		return "", err
	}

	return doc.ID, nil
}
