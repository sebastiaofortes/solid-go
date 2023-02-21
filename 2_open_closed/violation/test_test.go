package openclosed

func TestEmail() {
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
