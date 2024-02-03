package firebase_middleware 

import (
	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

// MIDDELWARE: GetUser - Get the User from Firestore using the userID
func GetUserFromFirebase(userID string) (*firestore.DocumentSnapshot, error) {
	// Client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Get the User from Firestore
	doc, err := client.Collection("Users").Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}

	// Get the Data from the Document
	return doc, nil
}
