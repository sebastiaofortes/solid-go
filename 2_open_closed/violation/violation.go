package openclosed

type Email struct {
	Type    string
	// All struct atributes
}

func (e Email) SendrHtml() {
	//send email in html format
}

func (e Email) SendTxt() {
	//send email in txt format
}

func SendMessage(email Email) {
	if email.Type == "Html" {
		email.SendrHtml()
	}
	if email.Type == "TXt" {
		email.SendTxt()
	}
}