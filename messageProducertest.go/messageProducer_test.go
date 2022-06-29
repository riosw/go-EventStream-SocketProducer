package messageproducer

import (
	"testing"

	mgor "github.com/riosw/go-EventStream-SocketProducer/messageProducer"
)

func TestSimpleMessage(t *testing.T) {
	generator := mgor.SimpleStringMessageGenerator{
		StringMessage: "car",
	}

	message := generator.EmitMessage().Content

	if message != "car" {
		t.Fatalf("Did not receive message 'car'")
	}

}
