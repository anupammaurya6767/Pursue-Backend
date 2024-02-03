package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/chatbot"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/types"
	"log"
	"net/http"
)

func AddQuestionController(ctx *gin.Context) {

	// Reading the request body
	var requestBody types.ChatBotQuestion
	err := ctx.Bind(&requestBody)
	if err != nil {
		log.Fatal(err)
	}

	// Making a request to the middleware
	id, err := chatbot.AddQuestion(*requestBody.ID, *requestBody.Section, *requestBody.Questions, requestBody.Options)

	// Returning a Response with Question ID
	ctx.String(http.StatusOK, "Added ChatBotQuestion to the Database with id: ", id)
}

func ReadQuestionController(ctx *gin.Context) {
	// Call the middleware function
	data, err := chatbot.ReadQuestion()
	if err != nil {
		log.Fatal(err)
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Return a response
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

// Method to Edit Question
func EditQuestionController(ctx *gin.Context) {
	// Binding the requestBody to the variable
	var requestBody types.ChatBotQuestion
	ctx.Bind(&requestBody)
	fmt.Println(*requestBody.ID)

	// Pass the requestBody into some sort of middleware
	resp, err := chatbot.EditQuestion(*requestBody.ID, *requestBody.Section, *requestBody.Questions, requestBody.Options)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Return a Response
	ctx.String(http.StatusOK, resp)
}

// Method to Delete Question
func DeleteQuestionController(ctx *gin.Context) {
	// Fetching the request ID from the path parameters
	id := ctx.Param("id")
	fmt.Println(id)

	// Pass the requestBody into some sort of middleware
	resp, err := chatbot.DeleteQuestion(id)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Return a Response
	ctx.String(http.StatusOK, resp)
}
