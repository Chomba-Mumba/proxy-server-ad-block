package dns

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

type Client struct {
	ip   net.IP
	port int
}

func newClient(ipAddress string, port int) (*Client, error) {
	i := net.ParseIP(ipAddress)
	if i == nil {
		return nil, fmt.Errorf("%v is an invalid ip address", ipAddress)
	}

	return &Client{
		ip:   i,
		port: port,
	}, nil
}

func (c *Client) ipType() (string, error) {
	ipAddress := c.ip.To4()
	if ipAddress != nil {
		return "ipv4", nil
	}
	return "ipv6", nil
}

// send message to given IP address and Port and return response
func (c *Client) Query(message []byte) ([]byte, error) {
	//create udp connection
	ipType, err := c.ipType()
	var addr string
	if err != nil {
		return nil, fmt.Errorf("failed to get the IP %v", err)
	}

	switch ipType {
	case "ipv4":
		addr = fmt.Sprintf("%s:%d", c.ip, c.port)
	case "ipv6":
		addr = fmt.Sprintf("[%s]:%d", c.ip, c.port)
	}

	conn, err := net.Dial("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("error making connection to DNS Server: %v", err)
	}

	defer conn.Close()

	//timeout for connection
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	_, err = conn.Write(message)
	if err != nil {
		return nil, fmt.Errorf("failed to send the DNS message: %v", err)
	}

	//receive response
	buf := make([]byte, 1024)

	//read response
	n, err := conn.Read(buf)

	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	response := buf[:n]

	//check if request matches the repsonse ID
	if !bytes.Equal(message[:2], response[:2]) {
		return nil, fmt.Errorf("the response ID does not match the request ID")
	}

	return response, nil
}
