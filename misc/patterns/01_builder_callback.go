package main

type email struct {
	from    string
	to      string
	subject string
	body    string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
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

}

type buildFn func(*EmailBuilder)

func SendEmail(action buildFn) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func Builder_Callback() {
	SendEmail(func(b *EmailBuilder){
		b.From("test@mail.com").To("my@mail.com").Subject("Meeting").Body("Hello!")
	})
}
