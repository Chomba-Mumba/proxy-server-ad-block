package dns

import (
	"bytes"
	"testing"
)

func TestEncodeQName(t *testing.T) {
	q := &Question{}

	got := q.encodeQName("test.dev")
	want := []byte{
		0x04, 't', 'e', 's', 't',
		0x03, 'd', 'e', 'v',
		0x00,
	}

	if !bytes.Equal(got, want) {
		t.Fatalf("unexpected bytes:\n got: %v\nwant: %v", got, want)
	}
}

func TestToBytes(t *testing.T) {
	q := &Question{}

	err := q.newQuestion("test.dev", 1, 1)
	if err != nil {
		t.Fatalf("failed to create new question: %v", err)
	}

	want := []byte{
		0x04, 't', 'e', 's', 't',
		0x03, 'd', 'e', 'v',
		0x00,
		0x00, 0x01,
		0x00, 0x01,
	}

	got, err := q.ToBytes()
	if err != nil {
		t.Fatalf("error converting question: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Fatalf("unexpected bytes:\n got: %v\nwant: %v", got, want)
	}
}
