package firebase_middleware 

import (
	"log"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// Make a function to count the total Number of Paid Users
func ListOfUsers() []map[string]interface{} {
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	iter := client.Collection("Users").Documents(ctx)
	var Users []map[string]interface{}

	for {
		data, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Users = append(Users, data.Data())
	}

	return Users
}