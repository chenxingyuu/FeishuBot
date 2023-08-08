package constant

type LakeBotType uint8

const (
	CustomBot LakeBotType = iota + 1
	ApplicationBot
)

func (t LakeBotType) String() string {
	switch t {
	case CustomBot:
		return "自定义机器人"
	case ApplicationBot:
		return "应用机器人"
	default:
		return "ErrorLakeBotType"
	}
}

type WebhookType uint8

const (
	GithubWebhook WebhookType = iota + 1
	GiteeWebhook
)

func (t WebhookType) String() string {
	switch t {
	case GithubWebhook:
		return "GithubWebhook"
	case GiteeWebhook:
		return "GiteeWebhook"
	default:
		return "ErrorBotType"
	}
}
