# IPChecker

IPChecker is a simple Go application that periodically checks your public IP address and sends you an email when it changes. This can be useful if you need to remotely access your home network and your ISP assigns you a dynamic IP address.

## Prerequisites

Before running IPChecker, you need to have the following:

1. A working email account that can send and receive emails
2. A SendGrid account and API key
3. A stationary system that can run Go applications
4. Docker installed on the system

## Installation

1. Clone this repository to your local machine
2. Install the dependencies using the following command:

    ```bash
    go mod download
    ```

3. Set the necessary environment variables by creating a .env file at the root of the project directory. The following variables are required:

    ```bash
    SENDER_EMAIL #SET THIS TO YOUR EMAIL ADDRESS
    RECEIVER_EMAIL #SET THIS TO THE EMAIL ADDRESS YOU WANT TO RECEIVE THE ALERTS
    CURRENT_IP #SET THIS TO 11.111.111.111
    SENDGRID_API_KEY
    ```

4. Build the application using the following command:

    ```bash
    make build
    make run
    ```

## Cloudflare integration

I have added the ability to update a DNS record on cloudflare.

You need to add a few things to your env file
```CLOUDFLARE_ZONE_ID=``` (Zone ID)
```CLOUDFLARE_DNS_ID=``` (DNS record ID)
```CLOUDFLARE_API_KEY=``` (API key from cloudflare)
```CLOUDFLARE_EMAIL=``` (email used to login to cloudflare)
```DOMAIN_NAME=``` (example.com)
```CLOUDFLARE=``` (set to true)

Find your ZoneID [here](https://developers.cloudflare.com/fundamentals/get-started/basic-tasks/find-account-and-zone-ids/).

Find your DNSID by ensuring your envs are up to date then running ```make get-dns-id```.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Example

![Example](example.png)
