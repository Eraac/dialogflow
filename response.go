package dialogflow

type (
	Response struct {
		Speech      string `json:"speech"`
		DisplayText string `json:"displayText"`
	}
)
