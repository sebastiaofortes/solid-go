package openclosed

type Email interface {
	Send()
}

type TxtEmail struct {
	Message string
}

func (e TxtEmail) Send() {

}

type HtmlEmail struct {
	Message string
}

func (e HtmlEmail) Send() {

}

type HtmlCSSEmail struct {
	HtmlEmail
}

func SendMessage(email Email) {
	email.Send()
}