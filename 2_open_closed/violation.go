package openclosed

type Email struct {
	Type    string
	Message string
}

func (e Email) SendrHtml() {

}

func (e Email) SendTxt() {

}

func SendMessage(email Email) {
	if email.Type == "Html" {
		email.SendrHtml()
	}
	if email.Type == "TXt" {
		email.SendTxt()
	}
}

func CreateEmail() {
	htmlEmail := Email{
		Type:    "Html",
		Message: "<h1> Hello </h1>",
	}
	txtEmail := Email{
		Type:    "Txt",
		Message: "Hello",
	}
	SendMessage(htmlEmail)
	SendMessage(txtEmail)
}
