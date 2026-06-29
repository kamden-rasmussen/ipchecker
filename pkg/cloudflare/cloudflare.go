package cloudflare

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Cloudflare struct {
	ZoneID     string
	DNSID      string
	Email      string
	APIKey     string
	DomainName string
}

// example curl request
// curl -X PUT "https://api.cloudflare.com/client/v4/zones/yourzoneidhere/dns_records/yourdnsidhere" \
//      -H "X-Auth-Email: user@example.com" \
//      -H "Authorization": yourauthkeyhere" \
//      -H "Content-Type: application/json" \
//      --data '{"type":"A","name":"example.com","content":"yournewiphere","ttl":1,"proxied":false}'

func (c Cloudflare) PutNewIP(ip string) (int, error) {
	if c.APIKey == "" || c.ZoneID == "" || c.DNSID == "" {
		return -1, fmt.Errorf("Cloudflare config invalid. Please ensure all envs for Cloudflare are properly defined")
	}

	// add ip to the body
	body := fmt.Sprintf(`{"type":"A","name":"%s","content":"%s","ttl":1,"proxied":false}`, c.DomainName, ip)

	// create the request
	req, err := http.NewRequest("PUT", "https://api.cloudflare.com/client/v4/zones/"+c.ZoneID+"/dns_records/"+c.DNSID, strings.NewReader(body))
	if err != nil {
		return -1, err
	}

	// add the headers
	req.Header.Add("X-Auth-Email", c.Email)
	req.Header.Add("Authorization", c.APIKey)
	req.Header.Add("Content-Type", "application/json")

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		// read the body
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		println("Error updating ip address with cloudflare")
		print(string(respBody))
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil

}
