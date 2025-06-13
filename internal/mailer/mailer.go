package mailer

const (
	maxRetries = 3
)

type Client interface {
	Send(recipient string, data *GmailData) error
}
