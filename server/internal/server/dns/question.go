package dns

import (
	"strconv"
	"strings"
)

type Question struct {
	QName  string // converted domain name
	QType  string // query type
	QClass string // query class
}

func (q *Question) encodeName(name string) string {
	domainParts := strings.Split(name, ".")
	qname := ""
	for _, part := range domainParts {
		newDomainPart := string(byte(len(part))) + part
		qname += newDomainPart
	}
	return qname + "\x00"
}

func (q *Question) decodeName(name string) string {
	encodedDomainSlice := strings.Split(string(name[len(name)-4]), "")
	domain := ""
	temp := 0 //start of last string

	for i := 0; i < len(name)-4; i++ {
		if encodedDomainSlice[i] == `\` {
			domain += strings.Join(encodedDomainSlice[i-temp:i], "")
			temp = i + 3
		}
	}
	return domain
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
