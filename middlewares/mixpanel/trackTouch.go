package MixPanel

import (
	"github.com/kunxl-gg/Amrit-Career-Counsellor.git/initialisers"
	"github.com/mixpanel/mixpanel-go"
)

func TrackMixpanelEvent(Event string, Object string) (string, error) {
	// Initialiser Mixpanel
	ctx, client := initialisers.InitialiseMixpanel()

	err := client.Track(ctx, []*mixpanel.Event{
		client.NewEvent(Event, "", map[string]interface{}{
			"Clicked Object": Object,
		}),
	})
	if err != nil {
		return "", err
	}

	return "Successfully Added Event", nil
}
