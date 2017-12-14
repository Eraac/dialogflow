package dialogflow

import "github.com/sirupsen/logrus"

// Config for the router
type Config struct {
	Debug  bool
	Token  string
	Logger logrus.FieldLogger
}

// IsDebug return true if router should display debug info
func (r *Router) IsDebug() bool {
	return r.config.Debug
}
