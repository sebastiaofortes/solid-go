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

type SubHtml struct {
	HtmlEmail
}

func SendMessage(email Email) {
	email.Send()
}

func CreateEmail() {
	htmlEmail := new(HtmlEmail)
	htmlEmail.Message = "<h1> Hello </h1>"

	txtEmail := new(TxtEmail)
	txtEmail.Message = "Hello"

	SubHtml := new(SubHtml)
	SubHtml.Message = "<h1> Hello </h1>"

	SendMessage(txtEmail)
	SendMessage(htmlEmail)
	SendMessage(SubHtml)

}
