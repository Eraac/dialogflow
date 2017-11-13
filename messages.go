package dialogflow

func (r *Response) AddText(res TextMessage, platforms ...string) {
	for _, p := range platforms {
		r.Messages = append(r.Messages, Message{
			Type:        TypeText,
			Platform:    p,
			TextMessage: res,
		})
	}
}

func (r *Response) AddImage(res ImageMessage, platforms ...string) {
	for _, p := range platforms {
		r.Messages = append(r.Messages, Message{
			Type:         TypeImage,
			Platform:     p,
			ImageURL:     res.ImageURL,
			ImageMessage: res,
		})
	}
}

func (r *Response) AddCard(res CardMessage, platforms ...string) {
	for _, p := range platforms {
		r.Messages = append(r.Messages, Message{
			Type:        TypeCard,
			Platform:    p,
			ImageURL:    res.ImageURL,
			Title:       res.Title,
			CardMessage: res,
		})
	}
}

func (r *Response) AddQuickReply(res QuickReplyMessage, platforms ...string) {
	for _, p := range platforms {
		r.Messages = append(r.Messages, Message{
			Type:              TypeQuickReply,
			Platform:          p,
			Title:             res.Title,
			QuickReplyMessage: res,
		})
	}
}

func (r *Response) AddCustom(res CustomMessage, platforms ...string) {
	for _, p := range platforms {
		r.Messages = append(r.Messages, Message{
			Type:          TypeCustomPayload,
			Platform:      p,
			CustomMessage: res,
		})
	}
}
