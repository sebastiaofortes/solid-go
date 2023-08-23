package openclosed

import "testing"

func TestEmail(t *testing.T) {
	htmlEmail := new(HtmlEmail)
	htmlEmail.Message = "<h1> Hello </h1>"

	txtEmail := new(TxtEmail)
	txtEmail.Message = "Hello"

	htmlCSSEmail := new(HtmlCSSEmail)
	htmlCSSEmail.Message = "<h1> Hello </h1>"

	SendMessage(txtEmail)
	SendMessage(htmlEmail)
	SendMessage(htmlCSSEmail)

}
