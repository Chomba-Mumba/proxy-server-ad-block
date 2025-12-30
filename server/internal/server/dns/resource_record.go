package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ResourceRecord struct {
	Name        string //domain name
	Type        uint16 //record type
	Class       uint16 //record class
	TTL         uint32 //time to live
	RDLength    uint16 //len of resource data
	RData       []byte //resource data
	RDataParsed string //parsed resource data
}

// initialise record fields from a byte slice
func (r *ResourceRecord) NewResourceRecordFromBytes(data []byte, messageBufs ...*bytes.Buffer) error {
	buf := bytes.NewBuffer(data)
	var messageBuf *bytes.Buffer
	if messageBufs != nil {
		messageBuf = messageBufs[0]
	}

	name, err := appendFromBufferUntilNull(buf)
	if err != nil {
		return fmt.Errorf("failed to decode name: %v", err)
	}
	nameLength := len(name) - 1

	decodedName, err := r.decodeName(string(name), messageBuf)
	if err != nil {
		fmt.Println("failed to decode the name %v", err)
	}

	typ := binary.BigEndian.Uint16(data[nameLength : nameLength+2])
	class := binary.BigEndian.Uint16(data[nameLength+2 : nameLength+4])
	ttl := binary.BigEndian.Uint32(data[nameLength+4 : nameLength+8])
	rdLength := binary.BigEndian.Uint16(data[nameLength+8 : nameLength+10])
	rdData := data[nameLength+10 : nameLength+10+int(rdLength)]
	rDataParsed, _ := parseRData(typ, rdData, messageBuf)

	r.Name = decodedName
	r.Type = typ
	r.Class = class
	r.TTL = ttl
	r.RDLength = rdLength
	r.RData = rdData
	r.RDataParsed = rDataParsed

	return nil
}

func (r *ResourceRecord) decodeName(name string, messageBufs ...*bytes.Buffer) (string, error) {
	encoded := []byte(name)
	var result bytes.Buffer
	var messageBuf *bytes.Buffer
	if messageBufs != nil {
		messageBuf = messageBufs[0]
	}

	for i := 0; i < len(encoded); i++ {
		length := int(encoded[i])
		if length == 0 {
			break
		}

		if encoded[i]>>6 == 0b11 && messageBuf != nil {
			b := encoded[i+1]
			offset := int(b & 0b11111111)
			messageBytes := messageBuf.Bytes()
			messageBytes = messageBytes[offset:]

			name, err := appendFromBufferUntilNull(bytes.NewBuffer(messageBytes))
			if err != nil {
				return "", fmt.Errorf("Error in reading name: %v", err)
			}
			n, _ := r.decodeName(string(name))

			name = []byte(n)
			length = len(name)
			if result.Len() > 0 {
				result.WriteByte('.')
			}
			result.Write(name)
			i += length
			break
		}
	}
	return "", nil
}

func parseRData(typ uint16, rdData []byte, messageBuf *bytes.Buffer) (string, error) {
	return "", nil
}

func appendFromBufferUntilNull(buffer *bytes.Buffer) ([]byte, error) {
	var res []byte
	for {
		b, err := buffer.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("failed to read from buffer: %v", err)
		}
		if b == 0x00 {
			break
		}
		res = append(res, b)
	}
	return res, nil
}
