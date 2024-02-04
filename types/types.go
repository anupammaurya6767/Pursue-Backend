package types

import (
	"github.com/gin-gonic/gin"
)

// Types for the Admin Routes and User Routes
var AdminRoutes *gin.RouterGroup
var UserRoutes *gin.RouterGroup

// Types for Firebase Schema
type FirebaseUser struct {
	UserID             *string
	OrderID            *string
	CustomerID         *string
	Name               *string
	Email              *string
	PhoneNumber        int
	DidStartChatbot    bool
	IsPaidUser         bool
	Options            []string
	FinalCareerOptions []string
}


type FirebaseAdminUser struct {
	Name     *string
	Email    *string
	Password *string
}

type FirebaseCareerOption struct {
	Name            *string            `json:"name"`
	Description     *string            `json:"description"`
	CareerPathSteps []string           `json:"careerPathSteps"`
	AverageSalaries []map[string][]int `json:"averageSalaries"`
	TopColleges     []TopCollge        `json:"topColleges"`
	Skills          []Skills           `json:"skills"`
	Courses         []Courses          `json:"courses"`
}

type Courses struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	RedirectURl *string `json:"redirectUrl"`
}

type Skills struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type TopCollge struct {
	Name           *string `json:"name"`
	Image          *string `json:"image"`
	RedirectionUrl *string `json:"redirectionUrl"`
}

// Types for the Payment Gateway
type JuspayCustomer struct {
	CustomerId   *string
	EmailAddress *string
	FirstName    *string
	LastName     *string
	MobileNumber int
	CountryCode  int
}

// Type to Store Questions for Chatbot
type ChatBotQuestion struct {
	ID        *string
	Section   *string
	Questions *string
	Options   []string
}
