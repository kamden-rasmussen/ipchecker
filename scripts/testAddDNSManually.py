from dotenv import load_dotenv
import os
import requests
loaded = load_dotenv(".env")
print(loaded)

# // example curl request
# // curl -X PUT "https://api.cloudflare.com/client/v4/zones/yourzoneidhere/dns_records/yourdnsidhere" \
# //      -H "X-Auth-Email: user@example.com" \
# //      -H "X-Auth-Key: yourauthkeyhere" \
# //      -H "Content-Type: application/json" \
# //      --data '{"type":"A","name":"example.com","content":"yournewiphere","ttl":{},"proxied":false}'

# set env path

env_zone_id = str(os.environ['CLOUDFLARE_ZONE_ID'])
env_bearer_token = "Bearer " + str(os.environ['CLOUDFLARE_API_KEY'])
env_email = str(os.environ['CLOUDFLARE_EMAIL'])
env_domain = str(os.environ['DOMAIN_NAME'])
env_dns_id = str(os.environ['CLOUDFLARE_DNS_ID'])

cloudflare_api = "https://api.cloudflare.com/client/v4/zones/" + env_zone_id + "/dns_records/" + env_dns_id
headers = {'Authorization':env_bearer_token, 'X-Auth-Email':env_email, 'Content-Type':'application/json'}

cloudflare_dns = cloudflare_api
cloudflare_dns_response = requests.put(cloudflare_dns, headers=headers, data='{"type":"A","name":"' + env_domain + '","content":"11.111.111.111","ttl":{},"proxied":false}')
print(cloudflare_dns_response.status_code)

if cloudflare_dns_response.status_code == 200:
    print("Ok")
    print(cloudflare_dns_response.content)

else:
    print(cloudflare_dns_response.content)


