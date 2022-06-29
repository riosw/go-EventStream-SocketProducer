package messageproducer

import (
	"testing"

	messageproducer "github.com/riosw/go-EventStream-SocketProducer/messageProducer"
)

func TestSimpleMessage(t *testing.T) {
	generator := messageproducer.SimpleString{
		StringMessage: "car",
	}

	message := generator.EmitMessage()

	if message != "car" {
		t.Fatalf("Did not receive message 'car'")
	}
}
