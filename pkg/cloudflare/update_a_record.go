package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
)

func UpdateARecord(email string, apiKey string, recordName string, zoneId string, newIp string) {
	api, err := cloudflare.New(apiKey, email)
	if err != nil {
		log.Errorf("Error loging in to cloudflare: %s", err.Error())
		return
	}

	records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneId), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		log.Errorf("Error listing records: %s", err.Error())
		return
	}

	for _, r := range records {
		if r.Name == recordName {
			updatedRecord, err := api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneId), cloudflare.UpdateDNSRecordParams{
				Type:     r.Type,
				Name:     r.Name,
				Content:  newIp,
				Data:     r.Data,
				ID:       r.ID,
				Priority: r.Priority,
				TTL:      r.TTL,
				Proxied:  r.Proxied,
				Comment:  &r.Comment,
				Tags:     r.Tags,
			})
			if err != nil {
				log.Errorf("Error updating DNS record %d", r)
				return
			}

			log.Infof("| %s %s %s | updated to | %s %s %s |", r.Type, r.Name, r.Content, updatedRecord.Type, updatedRecord.Name, updatedRecord.Content)
		}
	}
}
