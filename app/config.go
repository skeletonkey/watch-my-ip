// lib-instance-gen-go: File auto generated -- DO NOT EDIT!!!
package main

import "github.com/skeletonkey/lib-core-go/config"

var cfg *app

func getConfig() *app {
	config.LoadConfig("app", &cfg)
	return cfg
}
