package chatbot

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"google.golang.org/api/iterator"
)

func ReadQuestion() (map[string][]interface{}, error) {
	// Initialising client and context for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Reading all the documents from the collection
	iter := client.Collection("ChatbotQuestions").Documents(ctx)
	var questions = make(map[string][]interface{})

	for {
		// Getting the next document
		data, err := iter.Next()

		// Checking if we have reached the end of the list
		if err == iterator.Done {
			break
		}

		// Returning the error if there is any
		if err != nil {
			return nil, err
		}

		questionData := data.Data()
		questionData["ID"] = data.Ref.ID

		// Checking if the section exists in the map
		section := data.Data()["Section"].(string)
		_, exists := questions[section]
		if !exists {
			questions[section] = []interface{}{
				questionData,
			}
		} else {
			questions[section] = append(questions[section], questionData)
		}
	}

	// Returning the Questions List
	return questions, nil
}

func AddQuestion(ID string, Section string, Question string, Option []string) (string, error) {
	// Initialising client and context for firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Adding Question to Firebase
	_, err := client.Collection("ChatbotQuestions").Doc(ID).Set(ctx, map[string]interface{}{
		"Section":  Section,
		"Question": Question,
		"Options":  Option,
	})

	if err != nil {
		return "", err
	}

	// If everything goes fine send a success Message
	return ID, nil
}

// Edit a Question
func EditQuestion(ID string, Section string, Question string, Option []string) (string, error) {
	// Initialising firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Editing an element with the id provided
	_, err := client.Collection("ChatbotQuestions").Doc(ID).Update(ctx, []firestore.Update{
		{
			Path:  "Question",
			Value: Question,
		},
		{
			Path:  "Options",
			Value: Option,
		},
		{
			Path:  "Section",
			Value: Section,
		},
	})
	if err != nil {
		return "", err
	}

	return "Edited Data Successfully", nil
}

func DeleteQuestion(ID string) (string, error) {
	// Initialising firebase
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Checking if the value already exists or not
	doc, err := client.Collection("ChatbotQuestions").Doc(ID).Get(ctx)
	if !doc.Exists() {
		return "", fmt.Errorf("The element doesn't exist")
	}

	// Deleting an element with id provided
	_, err = client.Collection("ChatbotQuestions").Doc(ID).Delete(ctx)
	if err != nil {
		return "", err
	}

	// Returning the final Success Message
	return "Deleted Item Successfully", nil
}
