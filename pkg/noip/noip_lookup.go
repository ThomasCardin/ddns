package noip

import (
	"net"

	log "github.com/sirupsen/logrus"
)

type NoIpData struct {
	PingResult bool
	IP         string
}

func Ping(currentIp string, hostname string, notifyChan chan NoIpData) {
	ip, err := net.LookupIP(hostname)
	if err != nil {
		log.Errorf("Error looking up NoIP hostname IP %s : %s", hostname, err.Error())
		return
	}

	newIP := ip[0].String()
	if currentIp != newIP {
		log.Warnf("IP address updated from %s to: %s", currentIp, newIP)
		currentIp = newIP
		notifyChan <- NoIpData{
			PingResult: true,
			IP:         newIP,
		}
	} else {
		log.Info("Unchanged IP address")
		notifyChan <- NoIpData{
			PingResult: false,
		}
	}
}
