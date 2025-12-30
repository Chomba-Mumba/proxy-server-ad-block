package dns

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type Question struct {
	QName  []byte // byte slice representation of a string
	QType  uint16 // query type
	QClass uint16 // query class
}

// convert question to its binary representation
func (q *Question) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	_, err := buf.Write([]byte(q.QName))
	if err != nil {
		return nil, fmt.Errorf("error  in creating buffer: %v", err)
	}

	err = binary.Write(buf, binary.BigEndian, q.QType)
	if err != nil {
		return nil, fmt.Errorf("error in converting data to binary representation: %v", err)
	}

	err = binary.Write(buf, binary.BigEndian, q.QClass)
	if err != nil {
		return nil, fmt.Errorf("error in converting data to binary representation: %v", err)
	}

	return buf.Bytes(), nil
}

func encodeQName(domain string) []byte {
	domainParts := strings.Split(domain, ".")
	var qname []byte
	for _, part := range domainParts {
		qname = append(qname, byte(len(part)))
		qname = append(qname, []byte(part)...)

	}
	qname = append(qname, 0x00) // root terminator

	return qname
}

func decodeQName(qname []byte) string {
	return ""
}
