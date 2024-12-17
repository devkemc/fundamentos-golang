package emails

import (
	"context"
	"fmt"
	"time"
)

type emailServiceSimulator struct {
}

func (e *emailServiceSimulator) SendEmail(ctx context.Context, email Email) error {
	fmt.Printf("Sending email %v\n", email)
	time.Sleep(time.Second * 10)
	return nil
}

func NewEmailServiceSimulator() EmailService {
	return &emailServiceSimulator{}
}
