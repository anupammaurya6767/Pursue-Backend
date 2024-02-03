package firebase_middleware 

import (
	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func UpdateStageOfUser(stage int, userId string) error {

	// Initialising the Context and Client
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Updating the value of stage in Firestore
	_, err := client.Collection("Users").Doc(userId).Update(ctx, []firestore.Update{
		{
			Path:  "Stage",
			Value: stage,
		},
	})
	if err != nil {
		return err
	}

	// Returning nil if everything goes fine
	return nil
}
