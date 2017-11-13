package dialogflow

type (
	Response struct {
		Speech      string `json:"speech,omitempty"`
		DisplayText string `json:"displayText,omitempty"`
		// Data          Data           `json:"data,omitempty"`
		Messages      []Message      `json:"messages,omitempty"`
		ContextOut    Contexts       `json:"contextOut,omitempty"`
		Source        string         `json:"source,omitempty"`
		FollowUpEvent *FollowUpEvent `json:"followupEvent,omitempty"`
	}

	// Data map[string]interface{}

	FollowUpEvent struct {
		Name string     `json:"name"`
		Data Parameters `json:"data"`
	}

	TextMessage struct {
		Speech string `json:"speech,omitempty"`
	}

	ImageMessage struct {
		ImageURL string
	}

	CardMessage struct {
		Buttons  []Button `json:"buttons,omitempty"`
		ImageURL string
		Subtitle string `json:"subtitle,omitempty"`
		Title    string
	}

	Button struct {
		Text     string `json:"text,omitempty"`
		PostBack string `json:"postback,omitempty"`
	}

	QuickReplyMessage struct {
		Replies []string `json:"replies,omitempty"`
		Title   string
	}

	CustomMessage struct {
		Payload interface{} `json:"payload,omitempty"`
	}
)

func NewResponse() *Response {
	return &Response{}
}
