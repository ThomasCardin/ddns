# ddns

Managed Cloudflare A record based on No-IP hostname IP.

## Environment variables

- `NOIP_HOSTNAME`: The hostname or domain name associated with the NOIP service
- `CLOUDFLARE_API_KEY`: An authentication key used to access the Cloudflare API for performing various operations
- `CLOUDFLARE_A_RECORD`: A specific type of DNS record (Address Record) in the Cloudflare DNS settings, typically used to map domain names to IPv4 addresses
- `CLOUDFLARE_ZONE_ID`: The unique identifier associated with a domain's zone in Cloudflare's system. It's used to specify which zone the DNS record belongs to

## Test

Building the image with your arguments:

```bash
docker build --build-arg ARG_NOIP_HOSTNAME=<ARG_NOIP_HOSTNAME> \
             --build-arg ARG_CLOUDFLARE_EMAIL=<ARG_CLOUDFLARE_EMAIL> \
             --build-arg ARG_CLOUDFLARE_API_KEY=<ARG_CLOUDFLARE_API_KEY> \
             --build-arg ARG_CLOUDFLARE_A_RECORD_NAME=<ARG_CLOUDFLARE_A_RECORD_NAME> \
             --build-arg ARG_CLOUDFLARE_ZONE_ID=<ARG_CLOUDFLARE_ZONE_ID> \
             -t ddns --no-cache .
```

```bash
docker run ddns
```

Pass your environment variables after the build

```bash
docker build -t ddns .
```

```bash
docker run ddns \
        -e NOIP_HOSTNAME=<NOIP_HOSTNAME> \
        -e CLOUDFLARE_EMAIL=<CLOUDFLARE_EMAIL> \
        -e CLOUDFLARE_API_KEY=<CLOUDFLARE_API_KEY> \
        -e CLOUDFLARE_A_RECORD_NAME=<CLOUDFLARE_A_RECORD_NAME> \
        -e CLOUDFLARE_ZONE_ID=<CLOUDFLARE_ZONE_ID> \

```

## Dependencies

 - [Logrus](https://github.com/sirupsen/logrus)
 - [cloudflare-go](https://github.com/cloudflare/cloudflare-go)
