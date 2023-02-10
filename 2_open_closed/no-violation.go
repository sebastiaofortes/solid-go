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
	TxtEmail	
}

func (e HtmlEmail) Send() {

}

func CorrectSendMessage(email InterfacetEmail) {
	email.Send()
}

func CorrectCreateEmail() {
	htmlEmail := new(HtmlEmail)
	htmlEmail.Message = "<h1> Hello </h1>"


	txtEmail := new(TxtEmail)
	txtEmail.Message = "Hello"

	CorrectSendMessage(txtEmail)
	CorrectSendMessage(htmlEmail)
	
}
