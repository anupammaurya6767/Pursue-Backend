package firebase_middleware

import (
	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

func EditCareerDescription(id string, careerDescription types.FirebaseCareerOption) (string, error) {
	// Initialising the Context and Client
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Updating the value of stage in Firestore
	_, err := client.Collection("CareerDescription").Doc(id).Update(ctx, []firestore.Update{
		{
			Path:  "Name",
			Value: careerDescription.Name,
		},
		{
			Path: "Description",
			Value: careerDescription.Description,
		},
		{
			Path: "Skills",
			Value: careerDescription.Skills,
		},
		{
			Path: "AverageSalaries",
			Value: careerDescription.AverageSalaries,
		},
		{
			Path: "Courses",
			Value: careerDescription.Courses,
		},
		{
			Path: "TopColleges",
			Value: careerDescription.TopColleges,
		},
		{
			Path: "CareerPathSteps",
			Value: careerDescription.CareerPathSteps,
		},
	})
	if err != nil {
		return "", err
	}

	// Returning nil if everything goes fine
	return "Career Description with " + id + " is Updated Successfully", nil
}


func DeleteCareerDescription(id string) (string, error) {
	// Initialising the Context and Client
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Deleting the value from Firestore
	_, err := client.Collection("CareerDescription").Doc(id).Delete(ctx)
	if err != nil {
		return "", err
	}

	// Returning a success message if everything goes fine
	return "Career Description with " + id + " is Deleted Successfully", nil
}
