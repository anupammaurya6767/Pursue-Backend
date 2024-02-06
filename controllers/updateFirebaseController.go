package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	firebase_middleware "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/firebase"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// AddUserController - Method to add User to Firestore
func AddUserController(ctx *gin.Context) {
	// Initializing a global variable for the Firebase User
	var user types.FirebaseUser

	// Binding the incoming JSON to the User variable
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Println("There was an error in reading the Incoming JSON:", err)
		ctx.String(http.StatusBadRequest, "Failed to read incoming JSON: "+err.Error())
		return
	}

	// Adding User to Firestore
	id, err := firebase_middleware.AddUserToFirebase(user)
	if err != nil {
		log.Println("There was an error in adding user to Firestore:", err)
		ctx.String(http.StatusInternalServerError, "Failed to add user to Firestore: "+err.Error())
		return
	}

	// Returning the final userID
	ctx.String(http.StatusOK, id)
}

// DeleteUserController - Method to delete User from Firestore
func DeleteUserController(ctx *gin.Context) {
	// Fetching the UserID from the URL path
	userId := ctx.Param("userId")

	// Deleting the User from Firestore
	err := firebase_middleware.DeleteUser(userId)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to delete user from Firestore: "+err.Error())
		return
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Deleted User "+userId)
}

// UpdateStageOfUserController - Method to update the value of stage in Firestore
func UpdateStageOfUserController(ctx *gin.Context) {
	// Fetching the UserID from the URL path
	userId := ctx.Param("userId")

	var stage struct {
		Stage int
	}
	ctx.BindJSON(&stage)

	// Updating Stage for User
	err := firebase_middleware.UpdateStageOfUser(stage.Stage, userId)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to update stage for user: "+err.Error())
		return
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Stage for User "+userId)
}

// UpdateOptionsController - Each User selects multiple options before getting to the final result. This is the method to update the options selected by the user
func UpdateOptionsController(ctx *gin.Context) {
	// Fetching the UserID from the URL path
	userId := ctx.Param("userId")

	var options struct {
		NewOptionSelected *string `json:"newOptionSelected"`
	}
	ctx.BindJSON(&options)

	// Updating Options for User
	err := firebase_middleware.UpdateSelectedOption(userId, options.NewOptionSelected)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to update options for user: "+err.Error())
		return
	}

	// Sending 200 response if everything goes fine
	ctx.String(http.StatusOK, "Updated Options for User "+userId)
}

// AddCareerDescriptionToFirebaseController - Method to add career description to Firestore
func AddCareerDescriptionToFirebaseController(ctx *gin.Context) {
	// Initializing a global variable for the Firebase Career Description
	var careerDescription types.FirebaseCareerOption

	// Binding the incoming JSON to the Career Description variable
	err := ctx.BindJSON(&careerDescription)
	if err != nil {
		log.Println("There was an error in reading the Incoming JSON:", err)
		ctx.String(http.StatusBadRequest, "Failed to read incoming JSON: "+err.Error())
		return
	}

	// Adding Career Description to Firestore
	careerId, err := firebase_middleware.AddCareerDescriptionToFirebase(careerDescription)
	if err != nil {
		log.Println("There was an error in adding career description to Firestore:", err)
		ctx.String(http.StatusInternalServerError, "Failed to add career description to Firestore: "+err.Error())
		return
	}

	// Returning the final careerID
	ctx.String(http.StatusOK, careerId)
}
