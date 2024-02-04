package controllers

import (
	"github.com/gin-gonic/gin"
	MixPanel "github.com/kunxl-gg/Amrit-Career-Counsellor.git/middlewares/mixpanel"
	"net/http"
)

// TrackMixpanelController - Method to interact with Mixpanel
func TrackMixpanelController(ctx *gin.Context) {
	// Capturing the request body
	var mixpanelEvent struct {
		Event  *string
		Object *string
	}
	err := ctx.Bind(&mixpanelEvent)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to parse request body: "+err.Error())
		return
	}

	// Tracking an event using Mixpanel
	resp, err := MixPanel.TrackMixpanelEvent(*mixpanelEvent.Event, *mixpanelEvent.Object)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to track Mixpanel event: "+err.Error())
		return
	}

	// Sending a 200 response if everything goes right
	ctx.String(http.StatusOK, resp)
}
