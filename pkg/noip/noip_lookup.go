package noip

import (
	"net"

	log "github.com/sirupsen/logrus"
)

var currentIP string

func Ping(hostname string, notifyChan chan bool) {
	ip, err := net.LookupIP(hostname)
	if err != nil {
		log.Errorf("error looking up NoIP hostname IP %s : %s", hostname, err.Error())
		return
	}

	newIP := ip[0].String()
	if currentIP != newIP {
		log.Warnf("IP %s changed to %s ", currentIP, newIP)
		notifyChan <- true
	} else {
		log.Info("IP didnt change")
	}
}
