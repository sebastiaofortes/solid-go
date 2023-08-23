package openclosed

type Email interface {
	Send()
}

type TxtEmail struct {
	// All struct atributes
}

func (e TxtEmail) Send() {

}

type HtmlEmail struct {
	// All struct atributes
}

func (e HtmlEmail) Send() {

}

type HtmlCSSEmail struct {
	HtmlEmail
}

func SendMessage(email Email) {
	email.Send()
}
