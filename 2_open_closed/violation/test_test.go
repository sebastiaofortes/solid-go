package openclosed

import "testing"

func TestEmail(t *testing.T) {
	htmlEmail := Email{
		Type:    "Html",
	}
	txtEmail := Email{
		Type:    "Txt",
	}
	SendMessage(htmlEmail)
	SendMessage(txtEmail)
}
