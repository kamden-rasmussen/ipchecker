package godaddy

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Godaddy struct {
	Domain  string
	Domains []string
	Type    string
	Name    string
	Key     string
	Secret  string
}

/* example curl request
curl -X PUT "https://api.godaddy.com/v1/domains/$domain/records/$type/$name" \
-H "accept: application/json" \
-H "Content-Type: application/json" \
-H "Authorization: sso-key $key:$secret" \
-d "[{\"data\": \"$currentIp\"}]"
*/

func (g Godaddy) PutNewIP(ip string) (int, error) {
	filteredDomains := make([]string, 0, len(g.Domains))

	for _, d := range g.Domains {
		if strings.TrimSpace(d) == "" {
			continue
		}

		filteredDomains = append(filteredDomains, d)
	}

	if (g.Domain == "" && len(filteredDomains) == 0) || g.Key == "" || g.Secret == "" || g.Type == "" || g.Name == "" {
		return -1, fmt.Errorf("GoDaddy config invalid. Please ensure all envs for GoDaddy are properly defined")
	}

	// add ip to the body
	body := fmt.Sprintf(`[{"data":"%s"}]`, ip)

	if len(filteredDomains) == 0 && g.Domain != "" {
		filteredDomains = []string{g.Domain}
	}

	for _, d := range filteredDomains {
		// create the request
		req, err := http.NewRequest("PUT",
			fmt.Sprintf("https://api.godaddy.com/v1/domains/%s/records/%s/%s", d, g.Type, g.Name),
			strings.NewReader(body),
		)
		if err != nil {
			return -1, err
		}

		// add the headers
		req.Header.Add("accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", g.Key, g.Secret))

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

		if len(g.Domains) > 1 {
			fmt.Printf("Successfully updated domain: %s\n", d)
		}
	}

	return http.StatusOK, nil

}
