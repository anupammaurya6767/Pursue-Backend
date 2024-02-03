package controllers

import (
	"github.com/gin-gonic/gin"
	MixPanel "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/mixpanel"
	"net/http"
)

// TrackMixpanelController - Method to interact with Mixpanel
func TrackMixpanelController(ctx *gin.Context) {

	// Capturing the request body
	var MixpanelEvent struct {
		Event  *string
		Object *string
	}
	err := ctx.Bind(&MixpanelEvent)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Tracking an event using Mixpanel
	resp, err := MixPanel.TrackMixpanelEvent(*MixpanelEvent.Event, *MixpanelEvent.Object)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	// Sending a 200 response if everything goes right
	ctx.String(http.StatusOK, resp)

}
