package check

import (
	"strconv"
	"time"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
)

type Check struct {
	Name    string
	Address string
}

type IPIfyResp struct {
	IP string `json:"ip,omitempty"`
}

func CheckIP() string {
	println("\n\n" + time.Now().Format("2006-01-02 15:04:05"))

	// get current ip
	currentIP := GetIpify()

	if currentIP == "" || currentIP == "No answer" {
		outageCount, _ := strconv.Atoi(env.GetKey("OUTAGE_COUNT"))
		env.SetKey("OUTAGE_COUNT", strconv.Itoa(outageCount+1))
		println("Outage count: " + env.GetKey("OUTAGE_COUNT"))
		if outageCount > 12 {
			return "outage"
		}
		return ""
	}
	env.SetKey("OUTAGE_COUNT", strconv.Itoa(0))
	// get old ip
	oldIp := env.GetKey("CURRENT_IP")

	print("old ip: " + oldIp + " current ip: " + currentIP + "\n")
	if currentIP != oldIp {
		// set new ip
		env.SetKey("CURRENT_IP", currentIP)
		println("current env: " + env.GetKey("CURRENT_IP"))
		println("New IP found: " + currentIP)

		return currentIP
	}

	return ""
}
