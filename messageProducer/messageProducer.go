package messageproducer

type MessageProducer[M message] interface {
	EmitMessage() M
}

// Allow message to be this way to allow more types in the future
// TODO : add json type
type message interface {
	[]byte | []rune | string
}

// A simple generator
type SimpleString struct {
	StringMessage string
}

func (generator *SimpleString) EmitMessage() string {
	return generator.StringMessage
}

type MessageSender interface {
	SendMessage() error
}

// We can also make a more complex generator like this

type OrderedSelection struct {
	OrderedValues  []string
	currentPointer int
}

func (generator *OrderedSelection) EmitMessage() string {
	msg := generator.OrderedValues[generator.currentPointer]
	generator.currentPointer = (generator.currentPointer + 1) % len(generator.OrderedValues)
	return msg
}
