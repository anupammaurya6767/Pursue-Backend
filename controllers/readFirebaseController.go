// This File has the Controllers for all the Read Operations to Firebase

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	firebase_middleware "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/firebase"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
)

// CountUserController Function to fetch Total User Count
func CountUserController(ctx *gin.Context) {
	count := firebase_middleware.CountTotalUsers()
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// CountPaidUserController Method to fetch the total Count of Paid Users
func CountPaidUserController(ctx *gin.Context) {
	count, err := firebase_middleware.CountPaidUsers()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.String(http.StatusOK, strconv.Itoa(count))
}

// CountTotalChatbotSessionsController Method to fetch the total Count of Chatbot Users
func CountTotalChatbotSessionsController(ctx *gin.Context) {
	count, err := firebase_middleware.CountTotalChatBotSessions()
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	}
	ctx.String(http.StatusOK, strconv.Itoa(int(count)))
}

// ListOfUsersController Method to fetch Total User  List
func ListOfUsersController(ctx *gin.Context) {
	users := firebase_middleware.ListOfUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}

// ListOfPaidUsersController Method to fetch the list of Paid Users
func ListOfPaidUsersController(ctx *gin.Context) {
	users := firebase_middleware.ListOfPaidUsers()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": users,
		})
}

// GetUserController Method to fetch a single User
func GetUserController(ctx *gin.Context) {

	// Fetching the UserId from the path parameter
	userId := ctx.Param("userID")

	// Getting the user using the userID
	doc, err := firebase_middleware.GetUserFromFirebase(userId)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	// Returning Final Json when everything goes fine
	ctx.JSON(
		http.StatusOK,
		doc.Data(),
	)
}

// GetCareerOptionsController Method to fetch the Career Options of a User
func GetCareerOptionsController(ctx *gin.Context) {

	description := firebase_middleware.ReadCareerDescriptionFromFirebase()
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"result": description,
		},
	)
}

// GetCareerOptionController - Method to fetch one specific Career Option from the List of Options
func GetCareerOptionController(ctx *gin.Context) {

	// Fetching the Career Title from the Request Body
	var CareerTitle struct {
		Title *string
	}
	ctx.Bind(&CareerTitle)

	// Getting the Career Option from the List of Options
	careerOption := firebase_middleware.RetrieveCareerDescription(*CareerTitle.Title)

	// Returning the Career Option
	ctx.JSON(
		http.StatusOK,
		careerOption,
	)
}

// EditCareerDescriptionController - Method to Edit a Career Option
func EditCareerDescriptionController(ctx *gin.Context) {
	var requestBody types.FirebaseCareerOption
	var requestParam string = ctx.Param("id")

	// Editing the Career Option
	resp, err := firebase_middleware.EditCareerDescription(requestParam, requestBody)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to edit career option: "+err.Error())
		return
	}

	// Returning the Response
	ctx.String(http.StatusOK, resp)
}

// DeleteCareerDescriptionController - Method to Delete a Career Option
func DeleteCareerDescriptionController(ctx *gin.Context) {
	// Fetching the request parameter
	var requestParam string = ctx.Param("id")

	// Deleting the Career Option
	resp, err := firebase_middleware.DeleteCareerDescription(requestParam)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to delete career option: "+err.Error())
		return
	}

	// Returning the Response
	ctx.String(http.StatusOK, resp)
}
