package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spf13/pflag"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	Reset = "\033[0m"
	White = "\033[37m"
	Green = "\033[32m"
	Red   = "\033[31m"
	Cyan  = "\033[36m"
	RedBG = "\033[41m"
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

func parseFlags() ([]string, int, string, string) {
	var ips []string
	var interval int
	var msg string
	var spinnerStyle string

	pflag.StringSliceVar(&ips, "ips", []string{"8.8.8.8"}, "Custom IP addresses to ping (comma-separated)")
	pflag.IntVar(&interval, "interval", 250, "Custom interval (in ms) between checks")
	pflag.StringVar(&msg, "msg", "ping", "Custom message to send in the ICMP request")
	pflag.StringVar(&spinnerStyle, "spinner", "default", "Choose spinner style: default, or clock")
	pflag.Parse()

	return ips, interval, msg, spinnerStyle
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

func printStatus(isAlive bool, spinner []string, i int, ip string, index int) {
	if isAlive {
		fmt.Printf("\033[%d;0H%s %sOK%s %s%s    \n", index+1, spinner[i], Green, Reset, ip, Reset)
	} else {
		fmt.Printf("\033[%d;0H%s %s%sDOWN%s %s%s \n", index+1, spinner[i], RedBG, White, Reset, ip, Reset)
	}
}

func main() {
	ips, interval, msg, spinnerStyle := parseFlags()
	spinner := getSpinnerStyle(spinnerStyle)

	fmt.Print("\033[2J\033[?25l")
	defer fmt.Print("\033[?25h")

	status := make([]bool, len(ips))
	spinnerIndex := 0

	go func() {
		for {
			spinnerIndex = (spinnerIndex + 1) % len(spinner)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i, ip := range ips {
		go func(i int, ip string) {
			for {
				status[i] = checkConnection(ip, msg)
				time.Sleep(time.Duration(interval) * time.Millisecond)
			}
		}(i, ip)
	}

	for {
		for i, ip := range ips {
			printStatus(status[i], spinner, spinnerIndex, ip, i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
