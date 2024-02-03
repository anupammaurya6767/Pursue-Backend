package firebase_middleware

import (
	"fmt"
	"reflect"

	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func UpdateSelectedOption(userId string, newOptionSelected string) error  {

	// Initialising the Context and Client
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Fetch the current value of options
	var userDetails map[string]interface{}	
	doc, err := client.Collection("Users").Doc(userId).Get(ctx)
	if err != nil {
		return err
	}
	doc.DataTo(&userDetails)


	// Check if options field exists and is of expected type
	fmt.Println(reflect.TypeOf(userDetails["Options"]))
	currentOptions , ok := userDetails["Options"].([]interface{})
	if !ok {
		return fmt.Errorf("Expected type []interface{} for Options field, got %T", userDetails["Options"])
	}


	// Append the new option to the options array
	currentOptions = append(currentOptions, newOptionSelected)
	fmt.Println(currentOptions)
	fmt.Println(newOptionSelected)


	// Updating the value of stage in Firestore
	_, err = client.Collection("Users").Doc(userId).Update(ctx, []firestore.Update{
		{
			Path:  "Options",
			Value: currentOptions,
		},
	})
	if err != nil {
		return err
	}

	// Returning nil if everything goes fine
	return nil
}
