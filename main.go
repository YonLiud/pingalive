package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func checkConnection(ip, msg string) bool {
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
			Data: []byte(msg),
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

func parseFlags() (string, int, string, string) {
	var ip string
	var interval int
	var msg string
	var spinnerStyle string

	flag.StringVar(&ip, "ip", "8.8.8.8", "Custom IP address to ping")
	flag.IntVar(&interval, "interval", 250, "Custom interval (in ms) between checks")
	flag.StringVar(&msg, "msg", "ping", "Custom message to send in the ICMP request")
	flag.StringVar(&spinnerStyle, "spinner", "default", "Choose spinner style: default, or clock")
	flag.Parse()

	return ip, interval, msg, spinnerStyle
}

// ? Contributors! Add your spinner style here! make it unique and fun!

func getSpinnerStyle(spinnerStyle string) []string {
	switch spinnerStyle {
	case "clock":
		return []string{"|", "/", "-", "\\"}
	default:
		return []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	}
}

func printStatus(isAlive bool, spinner []string, i int, ip string) {
	const (
		Reset = "\033[0m"
		White = "\033[37m"
		Green = "\033[32m"
		Red   = "\033[31m"
		Cyan  = "\033[36m"
		RedBG = "\033[41m"
	)

	if isAlive {
		fmt.Printf("\r%s %sOK%s %s%s.    ", spinner[i], Green, Reset, ip, Reset)
	} else {
		fmt.Printf("\r%s %s%sDEAD%s %s%s. ", spinner[i], RedBG, White, Reset, ip, Reset)
	}
}

func main() {
	ip, interval, msg, spinnerStyle := parseFlags()
	spinner := getSpinnerStyle(spinnerStyle)

	i := 0

	for {
		if checkConnection(ip, msg) {
			printStatus(true, spinner, i, ip)
		} else {
			printStatus(false, spinner, i, ip)
		}

		i = (i + 1) % len(spinner)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
