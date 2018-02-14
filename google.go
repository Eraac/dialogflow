package dialogflow

const (
	GoogleCapabilityAudioOutput        = "actions.capability.AUDIO_OUTPUT"
	GoogleCapabilityScreenOutput       = "actions.capability.SCREEN_OUTPUT"
	GoogleCapabilityMediaResponseAudio = "actions.capability.MEDIA_RESPONSE_AUDIO"
	GoogleCapabilityWebBrowser         = "actions.capability.WEB_BROWSER"

	GoogleInputTypeVoice = "VOICE"

	GoogleActionsIntentText       = "actions.intent.TEXT"
	GoogleActionsIntentPermission = "actions.intent.PERMISSION"

	GooglePermissionName            = "NAME"
	GooglePermissionCoarseLocation  = "DEVICE_COARSE_LOCATION"
	GooglePermissionPreciseLocation = "DEVICE_PRECISE_LOCATION"

	GooglePlaceholderForPermission = "PLACEHOLDER_FOR_PERMISSION"
)

type (
	// === From request ===

	DataRequestGoogle struct {
		IsInSandbox       bool               `json:"isInSandbox"`
		Surface           SurfaceGoogle      `json:"surface"`
		Inputs            []InputGoogle      `json:"inputs"`
		User              UserGoogle         `json:"user"`
		Device            DeviceGoogle       `json:"device"`
		Conversation      ConversationGoogle `json:"conversation"`
		AvailableSurfaces SurfaceGoogle      `json:"available_surfaces"`
	}

	SurfaceGoogle struct {
		Capabilities []CapabilityGoogle `json:"capabilities"`
	}

	CapabilityGoogle struct {
		Name string `json:"name"`
	}

	InputGoogle struct {
		Intent    string           `json:"intent"`
		RawInputs []RawInputGoogle `json:"raw_inputs"`
		Arguments []ArgumentGoogle `json:"arguments"`
	}

	RawInputGoogle struct {
		Query     string `json:"query"`
		InputType string `json:"inputType"`
	}

	ArgumentGoogle struct {
		RawText   string `json:"rawText"`
		TextValue string `json:"textValue"`
		Name      string `json:"name"`
	}

	UserGoogle struct {
		LastSeen    string   `json:"lastSeen"`
		AccessToken string   `json:"accessToken"`
		Permissions []string `json:"permissions"`
		Locale      string   `json:"locale"`
		UserID      string   `json:"userId"`
	}

	DeviceGoogle struct {
		Location LocationGoogle `json:"location"`
	}

	LocationGoogle struct {
		Coordinates CoordinatesGoogle `json:"coordinates"`
	}

	CoordinatesGoogle struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	ConversationGoogle struct {
		ConversationID    string `json:"conversationId"`
		Type              string `json:"type"`
		ConversationToken string `json:"conversationToken"`
	}

	// === For response ===

	DataResponseGoogle struct {
		ExpectUserResponse bool          `json:"expect_user_response"`
		IsSSML             bool          `json:"is_ssml"`
		NoInputPrompts     []interface{} `json:"no_input_prompts"`
		SystemIntent       SystemIntent  `json:"system_intent"`
	}

	SystemIntent struct {
		Intent string      `json:"intent"`
		Data   interface{} `json:"data"`
	}

	AskPermission struct {
		Type        string   `json:"@type"`
		OptContext  string   `json:"opt_context"`
		Permissions []string `json:"permissions"`
	}
)

func (d DataRequestGoogle) HasCapabilities(capability string) bool {
	for _, c := range d.AvailableSurfaces.Capabilities {
		if c.Name == capability {
			return true
		}
	}

	return false
}

func (d DataRequestGoogle) IsGranted(permissions ...string) bool {
	permissionFound := false

	for _, permission := range permissions {
		permissionFound = false

		for _, p := range d.User.Permissions {
			if p == permission {
				permissionFound = true
				break
			}
		}

		if !permissionFound {
			return false
		}
	}

	return permissionFound
}
