package check

import (
	"log"
	"os/exec"
	"strings"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
)

// Server:		resolver1.opendns.com
// Address:	208.67.222.222#53

// Non-authoritative answer:
// Name:	myip.opendns.com
// Address: 11.111.111.111

type Check struct {
	Name string
	Address string
}

func GetIp() string {
	
	// run script to get ip
	ip := exec.Command("bash", "script.sh")
	out, err := ip.Output()
	if err != nil {
		log.Println(err)
	}

	// turn out into string
	newOut := out[:len(out)-1]

	// parse output based on \n
	newlines := strings.Split(string(newOut), "\n")
	// log.Println(newString[5])

	// parse output based on :
	newString := strings.Split(newlines[5], ":")
	newString[1] = strings.TrimSpace(newString[1])
	ipAddr := newString[1]
	// log.Println("here:" + ipAddr)

	return ipAddr
}

func CheckIp() string {
	// get current ip
	currentIp := GetIp()

	// get old ip
	oldIp := env.GetKey("CURRENT_IP")

	log.Print("old ip: " + oldIp + " current ip: " + currentIp)
	if currentIp != oldIp {
		// set new ip
		env.SetKey("CURRENT_IP", currentIp)
		log.Println("current env: " + env.GetKey("CURRENT_IP"))
		log.Println("New IP found: " + currentIp)

		return currentIp
	}

	return ""
}
