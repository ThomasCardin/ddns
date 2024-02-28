package main

import (
	"os"
	"time"

	"github.com/ThomasCardin/ddns/pkg/cloudflare"
	"github.com/ThomasCardin/ddns/pkg/noip"
	log "github.com/sirupsen/logrus"
)

const (
	NOIP_HOSTNAME = "NOIP_HOSTNAME"

	CLOUDFLARE_EMAIL         = "CLOUDFLARE_EMAIL"
	CLOUDFLARE_API_KEY       = "CLOUDFLARE_API_KEY"
	CLOUDFLARE_A_RECORD_NAME = "CLOUDFLARE_A_RECORD_NAME"
	CLOUDFLARE_ZONE_ID       = "CLOUDFLARE_ZONE_ID"
)

var IP string

func main() {
	// Lookup env variables
	noIpHostname, found := os.LookupEnv("NOIP_HOSTNAME")
	if !found {
		log.Fatalf("%s env var not found!", NOIP_HOSTNAME)
	}

	cEmail, found := os.LookupEnv("CLOUDFLARE_EMAIL")
	if !found {
		log.Fatalf("%s env var not found!", CLOUDFLARE_EMAIL)
	}

	cApiKey, found := os.LookupEnv("CLOUDFLARE_API_KEY")
	if !found {
		log.Fatalf("%s env var not found!", CLOUDFLARE_API_KEY)
	}

	cRecordName, found := os.LookupEnv("CLOUDFLARE_A_RECORD_NAME")
	if !found {
		log.Fatalf("%s env var not found!", CLOUDFLARE_A_RECORD_NAME)
	}

	cZoneName, found := os.LookupEnv("CLOUDFLARE_ZONE_ID")
	if !found {
		log.Fatalf("%s env var not found!", CLOUDFLARE_ZONE_ID)
	}

	// Start
	ticker := time.NewTicker(30 * time.Second)
	notifyChan := make(chan noip.NoIpData, 1)

	for {
		select {
		case <-ticker.C:
			noip.Ping(noIpHostname, notifyChan)
		case noIpData := <-notifyChan:
			if noIpData.PingResult {
				cloudflare.UpdateARecord(cEmail, cApiKey, cRecordName, cZoneName, noIpData.IP)
			}
		}
	}
}
