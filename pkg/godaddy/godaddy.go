package godaddy

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Godaddy struct {
	Domain string
	Type   string
	Name   string
	Key    string
	Secret string
}

/* example curl request
curl -X PUT "https://api.godaddy.com/v1/domains/$domain/records/$type/$name" \
-H "accept: application/json" \
-H "Content-Type: application/json" \
-H "Authorization: sso-key $key:$secret" \
-d "[{\"data\": \"$currentIp\"}]"
*/

const BASE_URL = "https://api.godaddy.com/v1/domains"

func (g Godaddy) PutNewIP(ip string) (int, error) {

	// add ip to the body
	body := fmt.Sprintf(`[{"data":"%s"}]`, ip)

	// create the request
	req, err := http.NewRequest("PUT",
		fmt.Sprintf("%s/%s/records/%s/%s", BASE_URL, g.Domain, g.Type, g.Name),
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
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// read the body
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		println("Error updating ip address!")
		print(string(respBody))
		return resp.StatusCode, err
	}

	return resp.StatusCode, nil

}
