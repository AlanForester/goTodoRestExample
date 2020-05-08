package model

import (
	"fmt"
	"net"
	"time"
)

type Proxy struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

func (p *Proxy) Check() bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", p.IP, p.Port), 10*time.Second)
	if err == nil {
		defer conn.Close()
		return true
	}
	return false
}
