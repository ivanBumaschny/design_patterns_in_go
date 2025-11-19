package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	// although it'll build the email, it wont expose the critical parts
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain an @")
	}
	b.email.from = from
	return b
}
func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should contain an @")
	}
	b.email.to = to
	return b
}
func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}
func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	fmt.Println(email)
}

// builder parameter -> a function that applies to the builder
type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@foo.com").
			Subject("Meeting").
			Body("Hello! Would you want to meet?")
	})
}
