package firebase_middleware 

import (
	"log"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// MIDDLEWARE: ListOfPaidUsers- Method to fetch the list of Paid Users
func ListOfPaidUsers() []map[string]interface{} {
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
		if data.Data()["Paid"] == "false" {
			Users = append(Users, data.Data())
		}
	}

	return Users
}
