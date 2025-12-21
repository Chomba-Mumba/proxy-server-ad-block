package resolver

import (
	"bytes"
	"encoding/binary"
)

type Message struct {
}

type Header struct {
	ID      uint16
	Flags   uint16 // control flags for dns question
	QDCOUNT uint16 // number of entries in question section
	ANCOUNT uint16 // number of resource records in the answer section
	NSCOUNT uint16 // number of name server resource records in authority seciton
	ARCOUNT uint16 // number of resource records in the additional records field
}

type HeaderFlag struct {
	QR     bool  // message is query or resp
	OPCODE uint8 // specifies kind of query
	AA     bool  // responding server is authoritative
	TC     bool  // message was truncated
	RD     bool  // recursion desired
	RA     bool  // recursion available in name server
	Z      uint8 // unused atm, reserved for future use
	RCode  uint8 // response code
}

// generate 16 bit flag value from individual components
func (hf *HeaderFlag) GenerateFlag() uint16 {
	qr := uint16(boolToInt(hf.QR))
	opcode := uint16(hf.OPCODE)
	aa := uint16(boolToInt(hf.AA))
	tc := uint16(boolToInt(hf.TC))
	rd := uint16(boolToInt(hf.RD))
	ra := uint16(boolToInt(hf.RA))
	z := uint16(hf.Z)
	rcode := uint16(hf.RCode)

	return uint16(qr<<15 | opcode<<11 | aa<<10 | tc<<9 | rd<<8 | ra<<7 | z<<4 | rcode)
}

// convert header to byte representation
func (h *Header) ToBytes() []byte {
	buf := new(bytes.Buffer)

	// add header fields in order
	binary.Write(buf, binary.BigEndian, h.ID)
	binary.Write(buf, binary.BigEndian, h.Flags)
	binary.Write(buf, binary.BigEndian, h.QDCOUNT)
	binary.Write(buf, binary.BigEndian, h.ANCOUNT)
	binary.Write(buf, binary.BigEndian, h.NSCOUNT)
	binary.Write(buf, binary.BigEndian, h.ARCOUNT)

	return buf.Bytes()
}

func boolToInt(b bool) uint8 {
	var out uint8
	if b {
		out = 1
	}
	return out
}
