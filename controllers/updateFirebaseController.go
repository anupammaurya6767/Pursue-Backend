package controllers

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	firebase_middleware "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/firebase"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// AddUserController - Method to add User to Firestore
func AddUserController(ctx *gin.Context) {

	// Initialising a Global variable for the Firebase User
	var user types.FirebaseUser

	// Binding the incoming JSON to the User variable
	err := ctx.Bind(&user)
	if err != nil {
		log.Fatal("There was an error in reading the Incoming Json", err)
	}

	// Adding User to Firestore
	userId, err := firebase_middleware.AddUserToFirebase(*user.Name, user.IsPaidUser, user.Options, user.FinalCareerOptions, *user.Email, user.PhoneNumber, user.DidStartChatbot)
	if err != nil {
		log.Fatal("There was an error in adding user to firestore", err)
	}

	// Returning the final userID
	ctx.String(http.StatusOK, string(userId))

}

// DeleteUserController - Method to delete User from Firestore
func DeleteUserController(ctx *gin.Context) {

	// Fetching the UserId from the URL path
	userId := ctx.Param("userId")

	// Deleting the User from Firestore
	err := firebase_middleware.DeleteUser(userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Deleted User ", userId)

}

// UpdateStageOfUserController - Method to update value of stage in Firestore
func UpdateStageOfUserController(ctx *gin.Context) {

	// Fetching the UserId from the URL path
	var userId string = ctx.Param("userId")

	var stage struct {
		Stage int
	}
	ctx.Bind(&stage)
	fmt.Println(stage)

	// Updating Stage for User
	err := firebase_middleware.UpdateStageOfUser(stage.Stage, userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Stage for User ", userId)

}

// UpdateOptionsController - Each User selects multiple options before getting to the final result, This is the method to update the options selected by the user
func UpdateOptionsController(ctx *gin.Context) {

	// Fetching the UserId from the URL path
	var userId string = ctx.Param("userId")
	fmt.Println(userId)

	var options struct {
		NewOptionSelected *string
	}
	ctx.Bind(&options)

	// Updating Stage for User
	err := firebase_middleware.UpdateSelectedOption(userId, *options.NewOptionSelected)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Options for User ", userId)
}

// UpdateFinalResultController - Method to update value of finalResult in Firestore
func AddCareerDescriptionToFirebaseController(ctx *gin.Context) {

	// Initialising a Global variable for the Firebase User
	var careerDescription types.FirebaseCareerOption

	// Binding the incoming JSON to the User variable
	err := ctx.Bind(&careerDescription)
	if err != nil {
		log.Fatal("There was an error in reading the Incoming Json", err)
	}

	// Adding User to Firestore
	careerId, err := firebase_middleware.AddCareerDescriptionToFirebase(*careerDescription.Name, *careerDescription.Description, careerDescription.TopColleges, careerDescription.AverageSalaries, careerDescription.CareerPathSteps, careerDescription.Courses, careerDescription.Skills)
	if err != nil {
		log.Fatal("There was an error in adding user to firestore", err)
	}

	// Returning the final userID
	ctx.String(http.StatusOK, string(careerId))
}
