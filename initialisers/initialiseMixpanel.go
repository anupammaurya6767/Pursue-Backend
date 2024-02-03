package initialisers

import (
	"context"
	"fmt"
	"os"

	"github.com/mixpanel/mixpanel-go"
)

// InitialiseMixpanel initialises the Mixpanel API Client and a background context
func InitialiseMixpanel() (context.Context, *mixpanel.ApiClient) {

	projectToken := os.Getenv("MIXPANEL_PROJECT_TOKEN")
	fmt.Println(projectToken)

	// Initialising the Mixpanel API Client and a background context
	ctx := context.Background()
	client := mixpanel.NewApiClient(projectToken)

	return ctx, client
}
