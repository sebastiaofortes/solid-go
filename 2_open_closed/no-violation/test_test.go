package openclosed

func TestEmail() {
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
