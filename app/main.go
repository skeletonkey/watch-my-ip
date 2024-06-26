package main

import (
	"fmt"
	"time"

	"github.com/skeletonkey/lib-core-go/logger"
	"github.com/skeletonkey/lib-core-go/notify/pushover"
	"github.com/skeletonkey/watch-my-ip/app/ip"
)

func main() {
	config := getConfig()
	log := logger.Get()

	log.Info().Msg("starting watch-my-ip service")

	var myIP string
	for {
		if curIP, err := ip.GetIp(); err != nil {
			log.Error().Err(err).Msg("error while attempting to get my IP address")
		} else {
			if myIP == "" { // first run
				myIP = curIP
			} else if myIP != curIP {
				err := pushover.Notify(fmt.Sprintf("IP has been changed from %s to %s", myIP, curIP))
				if err != nil {
					log.Error().Err(err).Msg("issue with pushover service")
				}
				log.Trace().Str("oldIP", myIP).Str("newIP", curIP).Msg("updating IP")
				myIP = curIP
			}
		}

		log.Trace().Int("interval", config.Interval).Msg("sleeping for a while")
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}
