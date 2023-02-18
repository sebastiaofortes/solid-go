package openclosed

type InterfacetEmail interface {
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

func CorrectSendMessage(email InterfacetEmail) {
	email.Send()
}

func CorrectCreateEmail() {
	htmlEmail := new(HtmlEmail)
	htmlEmail.Message = "<h1> Hello </h1>"

	txtEmail := new(TxtEmail)
	txtEmail.Message = "Hello"

	SubHtml := new(SubHtml)
	SubHtml.Message = "<h1> Hello </h1>"

	CorrectSendMessage(txtEmail)
	CorrectSendMessage(htmlEmail)
	CorrectSendMessage(SubHtml)

}
