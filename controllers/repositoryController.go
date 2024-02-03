package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/chatbot"
	"log"
	"net/http"
)

// ReadRepositoryController: Controller to fetch all the repository Details
func ReadRepositoryController(ctx *gin.Context) {
	// Reading the parameter from the URl
	id := ctx.Param("id")
	fmt.Println(id)

	// Making a call to the backend function
	data, err := chatbot.ReadRepository(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Returning the final response as a string
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

// WriteToRepository: Controller to feed in repository details to the database
func WriteToRepositoryController(ctx *gin.Context) {
	// Reading the request body
	var requestBody struct {
		ID                *string
		DatabaseTable     *string
		Parameters        []string
		CareerSuggestions []string
	}

	err := ctx.Bind(&requestBody)
	if err != nil {
		log.Fatal(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Call the function
	id, err := chatbot.AddUserChoicesInRepository(*requestBody.ID, *requestBody.DatabaseTable, requestBody.Parameters, requestBody.CareerSuggestions)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Returning the final response Id of the object created
	ctx.String(http.StatusOK, id)
}

// EditRepositoryController - Controller to edit Repository Options
func EditRepositoryController(ctx *gin.Context) {
	var requestBody struct {
		ID                *string
		DatabaseTable     *string
		Parameters        []string
		CareerSuggestions []string
	}
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to Fetch RequestBody: "+err.Error())
		return
	}

	// Passing these parameters in the middleware
	resp, err := chatbot.EditRepository(*requestBody.ID, *requestBody.DatabaseTable, requestBody.Parameters, requestBody.CareerSuggestions)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to edit repository: "+err.Error())
		return
	}

	ctx.String(http.StatusOK, resp)
}

// DeleteRepositoryController - Controller to delete an entry from the repository
func DeleteRepositoryController(ctx *gin.Context) {
	// Fetching the request Body
	var requestBody struct {
		ID            *string
		DatabaseTable *string
	}
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to Fetch RequestBody: "+err.Error())
		return
	}

	// Passing the parameters to the middleware
	resp, err := chatbot.DeleteRepository(*requestBody.ID, *requestBody.DatabaseTable)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to delete from repository: "+err.Error())
		return
	}

	ctx.String(http.StatusOK, resp)
}


// RepositoryController - Controller to fetch the final list of career Options
func FetchFinalCareerOptionsController(ctx *gin.Context) {
	requestParam := ctx.Param("id")
	fmt.Println(requestParam)

	// Fetching the request body
	var requestBody struct {
		DatabaseTable *string
	}
	ctx.Bind(&requestBody)

	// Making a call to the backend function
	data, err := chatbot.FetchFinalCareerOptions(requestParam, *requestBody.DatabaseTable)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Returning the final response as a string
	ctx.JSON(
		http.StatusOK,
		data,
	)
}
