module github.com/skeletonkey/watch-my-ip

go 1.22

require (
	github.com/skeletonkey/lib-core-go v0.2.1
	github.com/skeletonkey/lib-instance-gen-go v0.4.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace (
	github.com/skeletonkey/lib-core-go => ../lib-core-go
	github.com/skeletonkey/lib-instance-gen-go => ../lib-instance-gen-go
)
