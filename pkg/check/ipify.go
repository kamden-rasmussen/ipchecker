package check

import (
	"encoding/json"
	"io"
	"net/http"
)

// example response
// $ curl 'https://api.ipify.org?format=json'
// {"ip":"111.111.111.111"} // your ip address

func GetIpify() string {
	var IpIfy IpIfyResp

	resp, err := http.Get("http://api.ipify.org?format=json")
	if err != nil {
		println(err)
		return ""
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err)
		return ""
	}

	err = json.Unmarshal(bytes, &IpIfy)
	if err != nil {
		println(err)
		return ""
	}

	return IpIfy.Ip
}
