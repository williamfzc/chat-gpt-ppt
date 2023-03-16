package chat_gpt_ppt

type Client interface {
	SetToken(token string)
	Prepare(topics []string) error
	FillTopic(topic string) (*Topic, error)
}

type ClientType = string

const (
	ClientGpt35 = "GPT35"
)

func GetClient(clientType ClientType) Client {
	switch clientType {
	case ClientGpt35:
		return NewGpt35Client()
	}
	return nil
}
