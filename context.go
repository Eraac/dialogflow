package dialogflow

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type (
	Context struct {
		Request Request
		Logger  logrus.FieldLogger
	}
)

func (ctx *Context) GetUserID() string {
	switch ctx.Source() {
	case PlatformTelegram:
		return ctx.GetUserIDByKey("telegram_chat_id")
	case PlatformFacebook:
		return ctx.GetUserIDByKey("facebook_sender_id")
	case PlatformGoogle:
		d := ctx.GetGoogleData()
		return d.User.UserID
	}

	return ""
}

func (ctx *Context) GetUserIDByKey(key string) string {
	c, err := ctx.Request.Result.Contexts.Find("generic")

	if err != nil {
		ctx.Logger.Error(err)
		return ""
	}

	str, err := c.Parameters.GetString(key)

	if err != nil {
		ctx.Logger.Error(err)
		return ""
	}

	return str
}

func (ctx *Context) Source() string {
	return ctx.Request.OriginalRequest.Source
}

func (ctx Context) GetGoogleData() DataRequestGoogle {
	d := DataRequestGoogle{}

	if err := json.Unmarshal(ctx.Request.OriginalRequest.Data, &d); err != nil {
		ctx.Logger.Error(err)
	}

	return d
}
