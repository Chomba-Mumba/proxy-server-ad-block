package dns

import (
	"testing"
)

func TestEncodeName(t *testing.T) {
	question := Question{}
	domain := "test.dev"
	encodedNameWant := `\x03test\x03dev\x00`
	encodedNameRes, err := question.encodeName(domain)
	if encodedNameWant != encodedNameRes {
		t.Errorf(`qname = %q, %v, want match for %#q, nil`, encodedNameRes, err, encodedNameWant)
	}
}

func TestDecodeName(t *testing.T) {
	question := Question{}
	domainWant := "test.dev"
	encodedDomain := `\x03test\x03dev\x00`
	domainRes, err := question.decodeName(encodedDomain)
	if domainWant != domainRes {
		t.Errorf(`qname = %q, %v, want match for %#q, nil`, domainRes, err, domainWant)
	}

}
