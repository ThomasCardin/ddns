package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
)

func UpdateARecord(email string, apiKey string, recordName string, zoneName string) {
	api, err := cloudflare.New(apiKey, email)
	if err != nil {
		log.Errorf("error loging in to cloudflare: %s", err.Error())
		return
	}

	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Errorf("error getting the cloudflare zone id: %s", err.Error())
		return
	}

	records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		log.Errorf("error listing records: %s", err.Error())
		return
	}

	for _, r := range records {
		if r.Name == recordName {
			_, err := api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{
				Type:     r.Type,
				Name:     r.Name,
				Content:  r.Content,
				Data:     r.Data,
				ID:       r.ID,
				Priority: r.Priority,
				TTL:      r.TTL,
				Proxied:  r.Proxied,
				Comment:  &r.Comment,
				Tags:     r.Tags,
			})
			if err != nil {
				log.Errorf("error updating DNS record %d", r)
				return
			}
		}
	}
}
