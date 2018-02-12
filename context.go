package dialogflow

import "github.com/sirupsen/logrus"

type (
	Context struct {
		Request Request
		Logger logrus.Logger
	}
)

func (ctx *Context) Source() string {
	return ctx.Request.Source()
}
