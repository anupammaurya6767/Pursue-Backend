package chatbot

import (
	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

// ReadRepository: reads all the entries from the given repository
func ReadRepository(DatabaseTitle string) ([]map[string]interface{}, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Making an iterator for the given Table
	iter := client.Collection(DatabaseTitle).Documents(ctx)
	repositoryEntries := make([]map[string]interface{}, 0)

	for {
		data, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		repositoryData := make(map[string]interface{})
		originalRepositoryData := data.Data()

		for key, values := range originalRepositoryData {
			repositoryData[key] = values
		}
		repositoryData["ID"] = data.Ref.ID

		repositoryEntries = append(repositoryEntries, repositoryData)
	}

	// Returning all the repository entries
	return repositoryEntries, nil

}

func AddUserChoicesInRepository(ID string, DatabaseTitle string, Parameters []string, CareerOptions []string) (string, error) {
	// Initialising ctx and client for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding Entry to Firebase
	_, err := client.Collection(DatabaseTitle).Doc(ID).Set(ctx, map[string]interface{}{
		"Parameters":    Parameters,
		"CareerOptions": CareerOptions,
	})
	if err != nil {
		return "", err
	}

	return "Added an Entry to Repository", err
}

func DeleteRepository(ID string, DatabaseTable string) (string, error) {
	// Initialising firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Deleting the entry from firebase
	_, err := client.Collection(DatabaseTable).Doc(ID).Delete(ctx)
	if err != nil {
		return "", err
	}

	return "Deleted Item From Repository Successfully", nil
}


func FetchFinalCareerOptions(requestParam string, DatabaseTable string) (map[string]interface{}, error) {
	// Initialising firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Fetching the request object
	requestObject, err := client.Collection(DatabaseTable).Doc(requestParam).Get(ctx)
	if err != nil {
		return nil, err
	}

	// Reading the parameters from the request object
	return requestObject.Data(), err
}

func EditRepository(ID string, DatabaseTable string, Parametes []string, CareerOptions []string) (string, error) {
	// Initialising Firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	_, err := client.Collection(DatabaseTable).Doc(ID).Update(ctx, []firestore.Update{
		{
			Path:  "Parameters",
			Value: Parametes,
		},
		{
			Path:  "CareerOptions",
			Value: CareerOptions,
		},
	})
	if err != nil {
		return "", err
	}

	return "Successfully Edited Repository Table", nil
}
