package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func checkConnection(ip string) bool {
	conn, err := net.Dial("ip4:icmp", ip)
	if err != nil {
		return false
	}
	defer conn.Close()

	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO"),
		},
	}

	msgBytes, err := message.Marshal(nil)
	if err != nil {
		return false
	}

	_, err = conn.Write(msgBytes)
	if err != nil {
		return false
	}

	reply := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	_, err = conn.Read(reply)
	return err == nil
}

func main() {
	ip := "8.8.8.8"

	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	for {
		if checkConnection(ip) {
			fmt.Printf("OK %s.\n", ip)
		} else {
			fmt.Printf("DEAD %s.\n", ip)
		}
		time.Sleep(250 * time.Millisecond)
	}
}
