package check

import (
	"os/exec"
	"strconv"
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
	ip := exec.Command("sh", "script.sh")
	out, err := ip.Output()
	if err != nil {
		println(err)
	}

	// turn out into string
	newOut := out[:len(out)-1]
	println("\ncurrent return \n" + string(newOut) + "\n")

	// parse output based on \n
	newlines := strings.Split(string(newOut), "\n")
	// println(newString[5])

	// parse output based on :
	newString := strings.Split(newlines[5], ":")
	newString[1] = strings.TrimSpace(newString[1])
	ipAddr := newString[1]
	// println("here:" + ipAddr)

	return ipAddr
}

func CheckIp() string {
	// get current ip
	currentIp := GetIp()
	if currentIp == "" || currentIp == "No answer"{
		outageCount, _ := strconv.Atoi(env.GetKey("OUTAGE_COUNT"))
		env.SetKey("OUTAGE_COUNT", strconv.Itoa(outageCount + 1))
		println("Outage count: " + env.GetKey("OUTAGE_COUNT"))
		if outageCount > 12 {
			return "outage"
		}
		return ""
	}
	env.SetKey("OUTAGE_COUNT", strconv.Itoa(0))
	// get old ip
	oldIp := env.GetKey("CURRENT_IP")

	print("old ip: " + oldIp + " current ip: " + currentIp + "\n")
	if currentIp != oldIp {
		// set new ip
		env.SetKey("CURRENT_IP", currentIp)
		println("current env: " + env.GetKey("CURRENT_IP"))
		println("New IP found: " + currentIp)

		return currentIp
	}

	return ""
}
