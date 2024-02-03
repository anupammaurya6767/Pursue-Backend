package revenue

import (
	"fmt"
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
)

func ChangeRevenueModel() {
	ctx, client := initialisers.InitialiseFirebase()
	defer client.Close()

	// Checking if the revenue Model is Empty or not
	document, err := client.Collection("RevenueModel").Documents(ctx).GetAll()
	if err != nil {
		fmt.Println(err)
	}

	if len(document) == 0 {
		doc, _, err := client.Collection("RevenueModel").Add(ctx, map[string]interface{}{
			"IsRevenueEnabled": false,
			"Amount":           20,
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(doc.ID)
	} else {

	}
}
