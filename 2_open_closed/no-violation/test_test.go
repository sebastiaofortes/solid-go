package openclosed

import "testing"

func TestEmail(t *testing.T) {
	htmlEmail := HtmlEmail{}

	txtEmail := TxtEmail{}

	htmlCSSEmail := HtmlCSSEmail{}

	SendMessage(txtEmail)
	SendMessage(htmlEmail)
	SendMessage(htmlCSSEmail)

}
