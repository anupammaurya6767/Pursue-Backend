package firebase_middleware 

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
	"log"
)

func CountTotalUsers() int {
	// Initialising the global variables for client and context
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Getting the itreator for the db
	iter := client.Collection("Users").Documents(ctx)

	// Looping through the iterator
	var count int = 0
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			return count
		}
		if err != nil {
			log.Fatal(err)
		}
		count++
	}
}
