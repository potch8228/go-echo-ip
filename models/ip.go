package models

import (
	"net"
)

type Ip struct {
	V4 string `json:"v4"`
	V6 string `json:"v6"`
}

type NotValidIpError struct{}

func (e *NotValidIpError) Error() string {
	return "not valid IP"
}

func MakeIp(addr string) (*Ip, error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return nil, &NotValidIpError{}
	}

	if ip.DefaultMask() != nil {
		return &Ip{V4: addr, V6: ""}, nil
	}

	return &Ip{V4: "", V6: addr}, nil
}
