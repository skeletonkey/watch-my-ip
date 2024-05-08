// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package ip

import "github.com/skeletonkey/lib-core-go/config"

var cfg *ip

func getConfig() *ip {
	config.LoadConfig("ip", &cfg)
	return cfg
}
