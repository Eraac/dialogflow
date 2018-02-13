package dialogflow

const (
	DataSlack    = "slack_message"
	DataFacebook = "facebook_message"
	DataKik      = "kik_message"
	DataTelegram = "telegram_message"

	PlatformDialogflow = ""
	PlatformGoogle     = "google"
	PlatformFacebook   = "facebook"
	PlatformKik        = "kik"
	PlatformLine       = "line"
	PlatformSkype      = "skype"
	PlatformSlack      = "slack"
	PlatformTelegram   = "telegram"
	PlatformViber      = "viber"

	TypeText          = 0
	TypeCard          = 1
	TypeQuickReply    = 2
	TypeImage         = 3
	TypeCustomPayload = 4

	TypeGoogleSimpleResponse = "simple_response"
	TypeGoogleBasicCard      = "basic_card"
	TypeGoogleListCard       = "list_card"
	TypeGoogleSuggestionChip = "suggestion_chips"
	TypeGoogleCarouselCard   = "carousel_card"
	TypeGoogleLinkOut        = "link_out_chip"
	TypeGoogleCustomPayload  = "custom_payload"

	ContextAskPermission = "ask_permission"

	ParameterEventCallback = "event_callback"
	ParameterPermissionAsked = "permission_asked"
)
