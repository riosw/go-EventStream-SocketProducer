package messageproducer

type MessageProducer interface {
	EmitMessage()
}

type message struct {
	Content string
}

type SimpleStringMessageGenerator struct {
	StringMessage string
}

func (generator *SimpleStringMessageGenerator) EmitMessage() message {
	return message{Content: generator.StringMessage}
}
