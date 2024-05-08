package ip

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/skeletonkey/lib-core-go/logger"
)

func GetIp() (string, error) {
	ipConfig := getConfig()
	log := logger.Get()

	log.Trace().Str("url", ipConfig.Url).Str("regex", ipConfig.Pattern).Msg("checking for ip")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ipConfig.Url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	compiledRegex, err := regexp.Compile(ipConfig.Pattern)
	if err != nil {
		return "", err
	}

	match := compiledRegex.FindStringSubmatch(string(body))

	if match != nil {
		if len(match) > 2 {
			log.Error().Str("matches", strings.Join(match, "|")).Msg("multiple matches found - first one used and others ignored")
		}
		log.Trace().Str("ip", match[1]).Msg("found IP address")
		return match[1], nil
	} else {
		err = errors.New("did not find pattern in body")
		log.Error().Bytes("body", body).Str("pattern", ipConfig.Pattern).Msg(err.Error())
		return "", err
	}
}
