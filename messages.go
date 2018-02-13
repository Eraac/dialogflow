package dialogflow

func (r *Response) AddText(res TextMessage, platform string) {
	r.Messages = append(r.Messages, Message{
		Type:        TypeText,
		Platform:    platform,
		TextMessage: res,
	})
}

func (r *Response) AddImage(res ImageMessage, platform string) {
	r.Messages = append(r.Messages, Message{
		Type:         TypeImage,
		Platform:     platform,
		ImageURL:     res.ImageURL,
		ImageMessage: res,
	})
}

func (r *Response) AddCard(res CardMessage, platform string) {
	r.Messages = append(r.Messages, Message{
		Type:        TypeCard,
		Platform:    platform,
		ImageURL:    res.ImageURL,
		Title:       res.Title,
		CardMessage: res,
	})
}

func (r *Response) AddQuickReply(res QuickReplyMessage, platform string) {
	r.Messages = append(r.Messages, Message{
		Type:              TypeQuickReply,
		Platform:          platform,
		Title:             res.Title,
		QuickReplyMessage: res,
	})
}

func (r *Response) AddCustom(res CustomMessage, platform string) {
	r.Messages = append(r.Messages, Message{
		Type:          TypeCustomPayload,
		Platform:      platform,
		CustomMessage: res,
	})
}

func (r *Response) AddGoogleSimpleResponse(res GoogleSimpleResponse) {
	r.Messages = append(r.Messages, GoogleMessage{
		Type:                 TypeGoogleSimpleResponse,
		Platform:             PlatformGoogle,
		GoogleSimpleResponse: res,
	})
}

func (r *Response) AddGoogleCustomPayload(res interface{}) {
	r.Messages = append(r.Messages, GoogleMessage{
		Type:     TypeGoogleCustomPayload,
		Platform: PlatformGoogle,
		GooglePayload: GooglePayload{
			Payload: map[string]interface{}{
				PlatformGoogle: res,
			},
		},
	})
}

func (r *Response) AskForPermission(callback, reason string, permissions ...string) {
	r.Speech = GooglePlaceholderForPermission

	r.Data = Data{PlatformGoogle: DataResponseGoogle{
		ExpectUserResponse: true,
		IsSSML:             false,
		NoInputPrompts:     []interface{}{},
		SystemIntent: SystemIntent{
			Intent: "actions.intent.PERMISSION",
			Data: AskPermission{
				Type:        "type.googleapis.com/google.actions.v2.PermissionValueSpec",
				OptContext:  reason,
				Permissions: permissions,
			},
		},
	}}

	r.ContextOut = []DFContext{
		{Name: ContextAskPermission, Lifespan: 5, Parameters: Parameters{
			ParameterEventCallback:   callback,
			ParameterPermissionAsked: permissions,
		}},
	}
}
