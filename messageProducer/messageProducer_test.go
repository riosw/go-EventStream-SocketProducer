package messageproducer

import (
	"fmt"
	"testing"
)

func TestSimpleMessage(t *testing.T) {
	generator := SimpleStringMessageGenerator{
		StringMessage: "car",
	}

	message := generator.EmitMessage().Content

	fmt.Println(message)

	if message != "car" {
		t.Fatalf("Did not receive message 'car'")
	}

}
