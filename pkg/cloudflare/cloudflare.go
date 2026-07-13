package cloudflare

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Cloudflare struct {
	ZoneID      string
	DNSID       string
	Email       string
	APIKey      string
	DomainName  string
	DomainNames []string
}

// example curl request
// curl -X PUT "https://api.cloudflare.com/client/v4/zones/yourzoneidhere/dns_records/yourdnsidhere" \
//      -H "X-Auth-Email: user@example.com" \
//      -H "Authorization": yourauthkeyhere" \
//      -H "Content-Type: application/json" \
//      --data '{"type":"A","name":"example.com","content":"yournewiphere","ttl":1,"proxied":false}'

func (c Cloudflare) PutNewIP(ip string) (int, error) {
	filteredDomains := make([]string, 0, len(c.DomainNames))

	for _, d := range c.DomainNames {
		if strings.TrimSpace(d) == "" {
			continue
		}

		filteredDomains = append(filteredDomains, d)
	}

	if c.APIKey == "" || c.ZoneID == "" || c.DNSID == "" || c.Email == "" || (c.DomainName == "" && len(filteredDomains) == 0) {
		return -1, fmt.Errorf("Cloudflare config invalid. Please ensure all envs for Cloudflare are properly defined")
	}

	if len(filteredDomains) == 0 && c.DomainName != "" {
		filteredDomains = []string{c.DomainName}
	}

	for _, d := range filteredDomains {
		// add ip to the body
		body := fmt.Sprintf(`{"type":"A","name":"%s","content":"%s","ttl":1,"proxied":false}`, d, ip)

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

		if resp.StatusCode != http.StatusOK {
			// read the body
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Error updating ip address for domain: %s!\n", d)

			fmt.Println(string(respBody))

			return resp.StatusCode, err
		}

		fmt.Printf("Successfully updated domain: %s\n", d)
	}

	return http.StatusOK, nil

}
