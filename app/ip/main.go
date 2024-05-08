package ip

import (
	"github.com/skeletonkey/lib-core-go/logger"
)

func GetIp() string {
	ipConfig := getConfig()
	log := logger.Get()

	log.Trace().Str("url", ipConfig.Url).Str("regex", ipConfig.Pattern).Msg("checking for ip")

	ip := "1.1.1.1"

	log.Trace().Str("ip", ip).Msg("fount IP address")

	return ip
}
