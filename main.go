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
	CLOUDFLARE_ZONE_NAME     = "CLOUDFLARE_ZONE_NAME"
)

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

	cZoneName, found := os.LookupEnv("CLOUDFLARE_ZONE_NAME")
	if !found {
		log.Fatalf("%s env var not found!", CLOUDFLARE_ZONE_NAME)
	}

	// Start
	ticker := time.NewTicker(30 * time.Second)
	notifyChan := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				noip.Ping(noIpHostname, notifyChan)
			case <-notifyChan:
				cloudflare.UpdateARecord(cEmail, cApiKey, cRecordName, cZoneName)
			}
		}
	}()

	log.Info("Program closed gracefully")
}
