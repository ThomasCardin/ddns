package noip

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type NoIpData struct {
	PingResult bool
	IP         string
}

var currentIP string

func Ping(hostname string, notifyChan chan NoIpData) {
	ip, err := net.LookupIP(hostname)
	if err != nil {
		log.Errorf("error looking up NoIP hostname IP %s : %s", hostname, err.Error())
		return
	}

	newIP := ip[0].String()
	if currentIP != newIP {
		log.Warnf("IP address updated from %s to: %s", currentIP, newIP)
		currentIP = newIP
		notifyChan <- NoIpData{
			PingResult: true,
			IP:         newIP,
		}
	} else {
		log.Info("unchanged IP address")
		notifyChan <- NoIpData{
			PingResult: false,
		}
	}
}
