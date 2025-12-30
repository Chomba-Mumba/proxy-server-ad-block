package dns

import (
	"bytes"
	"testing"
)

func TestEncodeQName(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		want     []byte
		gotError error
	}{
		{
			name:   "valid A record",
			domain: "test.dev",
			want: []byte{
				0x04, 't', 'e', 's', 't',
				0x03, 'd', 'e', 'v',
				0x00,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeQName(tt.domain)
			if !bytes.Equal(got, tt.want) {
				t.Fatalf("unexpected bytes :\n got: %v\nwant: %v", got, tt.want)
			}
		})

	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		qtype    uint16
		qclass   uint16
		gotError error
		want     []byte
	}{
		{
			name:     "valid DNS A question",
			domain:   "test.dev",
			qtype:    1,
			qclass:   1,
			gotError: nil,
			want: []byte{
				0x04, 't', 'e', 's', 't',
				0x03, 'd', 'e', 'v',
				0x00,
				0x00, 0x01,
				0x00, 0x01,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Question{
				QName:  encodeQName(tt.domain),
				QType:  tt.qtype,
				QClass: tt.qclass,
			}

			got, err := q.ToBytes()
			if err != nil {
				t.Fatalf("error converting question: %v", err)
			}

			if !bytes.Equal(got, tt.want) {
				t.Fatalf("unexpected bytes:\n got: %v\nwant: %v", got, tt.want)
			}
		})
	}
}

func TestNewDNSMessage(t *testing.T) {
	tests := []struct {
		name          string
		questionSlice []*Question
		gotError      error
		want          []byte
	}{
		{
			name:     "valid DNS questions",
			gotError: nil,
			questionSlice: []*Question{
				&Question{
					QName:  encodeQName("test.dev"), // A record
					QClass: 1,
					QType:  1,
				},
			},
			want: []byte{
				0x04, 't', 'e', 's', 't',
				0x03, 'd', 'e', 'v',
				0x00,
				0x00, 0x01,
				0x00, 0x01,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestResolveDNS(t *testing.T) {

}
