package dialogflow

const (
	GoogleAudioOutput        = "actions.capability.AUDIO_OUTPUT"
	GoogleScreenOutput       = "actions.capability.SCREEN_OUTPUT"
	GoogleMediaResponseAudio = "actions.capability.MEDIA_RESPONSE_AUDIO"
)

type (
	DataGoogle struct {
		IsInSandbox bool          `json:"isInSandbox"`
		Surface     SurfaceGoogle `json:"surface"`
		// Inputs
		User              UserGoogle         `json:"user"`
		Conversation      ConversationGoogle `json:"conversation"`
		AvailableSurfaces SurfaceGoogle      `json:"available_surfaces"`
	}

	SurfaceGoogle struct {
		Capabilities []CapabilityGoogle `json:"capabilities"`
	}

	CapabilityGoogle struct {
		Name string `json:"name"`
	}

	UserGoogle struct {
		LastSeen string `json:"lastSeen"`
		Locale   string `json:"locale"`
		UserID   string `json:"userId"`
	}

	ConversationGoogle struct {
		ConversationID    string `json:"conversationId"`
		Type              string `json:"type"`
		ConversationToken string `json:"conversationToken"`
	}
)

func (d DataGoogle) HasCapabilities(capability string) bool {
	for _, c := range d.AvailableSurfaces.Capabilities {
		if c.Name == capability {
			return true
		}
	}

	return false
}
