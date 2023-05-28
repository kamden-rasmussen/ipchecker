package cloudflare

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/kamden-rasmussen/ipchecker/pkg/env"
)

// example curl request
// curl -X PUT "https://api.cloudflare.com/client/v4/zones/yourzoneidhere/dns_records/yourdnsidhere" \
//      -H "X-Auth-Email: user@example.com" \
//      -H "Authorization": yourauthkeyhere" \
//      -H "Content-Type: application/json" \
//      --data '{"type":"A","name":"example.com","content":"yournewiphere","ttl":{},"proxied":false}'

func PutNewIP(ip string) (int, error) {
	zoneID := env.GetKey("CLOUDFLARE_ZONE_ID")
	dnsID := env.GetKey("CLOUDFLARE_DNS_ID")
	email := env.GetKey("CLOUDFLARE_EMAIL")
	apiKey := "Bearer " + env.GetKey("CLOUDFLARE_API_KEY")
	domainName := env.GetKey("DOMAIN_NAME")

	// add ip to the body
	body := fmt.Sprintf(`{"type":"A","name":"%s","content":"%s","ttl":1,"proxied":false}`, domainName, ip)

	// create the request
	req, err := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/zones/"+zoneID+"/dns_records/"+dnsID, bytes.NewReader([]byte(body)))
	if err != nil {
		return -1, err
	}

	// add the headers
	req.Header.Add("X-Auth-Email", email)
	req.Header.Add("Authorization", apiKey)
	req.Header.Add("Content-Type", "application/json")

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// read the body
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		println("error updating ip address with cloudflare")
		print(string(respBody))
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil

}
