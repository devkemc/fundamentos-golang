package emails

import "context"

type EmailService interface {
	SendEmail(ctx context.Context, email Email) error
}
