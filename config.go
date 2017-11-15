package dialogflow

// Config for the router
type Config struct {
	Debug bool
	Token string
}

// IsDebug return true if router should display debug info
func (r *Router) IsDebug() bool {
	return r.config.Debug
}
