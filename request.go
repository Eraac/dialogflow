package dialogflow

type (
	Request struct {
		ID        string `json:"id"`
		Timestamp string `json:"timestamp"`
		Lang      string `json:"lang"`
		Result    Result `json:"result"`
		Status    Status `json:"status"`
		SessionID string `json:"sessionId"`
	}

	Result struct {
		Source           string      `json:"source"`
		ResolvedQuery    string      `json:"resolvedQuery"`
		Speech           string      `json:"speech"`
		Action           string      `json:"action"`
		ActionIncomplete bool        `json:"actionIncomplete"`
		Parameters       Parameters  `json:"parameters"`
		Contexts         Contexts    `json:"contexts"`
		Metadata         Metadata    `json:"metadata"`
		Fulfillment      Fulfillment `json:"fulfillment"`
		Score            float64     `json:"score"`
	}

	Parameters map[string]interface{}
	Contexts   []Context

	Context struct {
		Name       string     `json:"name"`
		Parameters Parameters `json:"parameters"`
		Lifespan   int        `json:"lifespan"`
	}

	Metadata struct {
		IntentID                  string `json:"intentId"`
		WebhookUsed               string `json:"webhookUsed"`
		WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
		IntentName                string `json:"intentName"`
	}

	Fulfillment struct {
		Speech   string    `json:"speech"`
		Messages []Message `json:"messages"`
	}

	Message struct {
		Type     int    `json:"type"`
		Platform string `json:"platform"`

		// Duplicate field between XxxMessage
		// If you put tag in XxxMessage go doesn't marshal it (because other is empty)
		ImageURL string `json:"imageUrl,omitempty"`
		Title    string `json:"title,omitempty"`

		TextMessage
		ImageMessage
		CardMessage
		QuickReplyMessage
		CustomMessage
	}

	Status struct {
		Code      int    `json:"code"`
		ErrorType string `json:"errorType"`
	}
)
