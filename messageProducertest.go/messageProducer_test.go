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

func TestOrderedSelection(t *testing.T) {
	generator := messageproducer.OrderedSelection{
		OrderedValues: []string{"car", "truck", "bike"},
	}

	message := generator.EmitMessage()
	if message != "car" {
		t.Fatalf("Did not receive message 'car'")
	}

	message = generator.EmitMessage()
	if message != "truck" {
		t.Fatalf("Did not receive message 'truck'")
	}

	message = generator.EmitMessage()
	if message != "bike" {
		t.Fatalf("Did not receive message 'bike'")
	}

	message = generator.EmitMessage()
	if message != "car" {
		t.Fatalf("Did not receive message 'car'")
	}
}
