package notifymodel

import "github.com/WolffunService/theta-shared-common/enums/serverenum"

type NotifyPlatform string

const (
	PlatformEmail    NotifyPlatform = "email"
	PlatformSlack    NotifyPlatform = "slack"
	PlatformTelegram NotifyPlatform = "telegram"
)

type (
	NotifyMessage struct {
		Environment serverenum.Environment `json:"environment"`
		Platform    NotifyPlatform         `json:"platform"`
		Channel     string                 `json:"channel"`
		Title       string                 `json:"title"`
		Content     string                 `json:"content"`
		DynamicData map[string]interface{} `json:"dynamicData"`
	}
)
