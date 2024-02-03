package firebase_middleware 

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func DeleteUser(userId string) error {

	// Initialise Firestore Client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Delete the User using the Firebase sdk
	_, err := client.Collection("Users").Doc(userId).Delete(ctx)
	if err != nil {
		return err
	}

	// Return nil if everything is fine
	return nil

}
