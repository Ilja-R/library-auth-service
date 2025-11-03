package broker

type Message struct {
	Recipient string
	Subject   string
	Body      string
}
