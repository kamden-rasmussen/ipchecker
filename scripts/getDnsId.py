import os
import requests

env_zone_id = os.environ.get('CLOUDFLARE_ZONE_ID')
env_bearer_token = os.environ.get('CLOUDFLARE_API_KEY')

env_email = os.environ.get('CLOUDFLARE_EMAIL')


cloudflare_api = "https://api.cloudflare.com/client/v4/"
zone_id = env_zone_id
auth_key = "Bearer " + env_bearer_token
headers = {'Authorization':auth_key, 'X-Auth-Email':env_email, 'Content-Type':'application/json'}

cloudflare_dns = cloudflare_api + "zones/" + zone_id + "/dns_records"  
cloudflare_dns_respon = requests.get(cloudflare_dns, headers=headers)

if cloudflare_dns_respon.status_code == 200:
    print("Ok")
    print(cloudflare_dns_respon.content)
    
else:
    print(cloudflare_dns_respon
          .content)
