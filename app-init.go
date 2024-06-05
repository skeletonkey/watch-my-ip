package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	goVersion := "1.22"
	app := instanceGen.NewApp("watch-my-ip", "app")
	app.SetupApp(
		app.WithConfig(),
		app.WithGithubWorkflows("linter", "test"),
		app.WithGoVersion(goVersion),
		app.WithMakefile(),
		app.WithPackages("ip"),
	).Generate()
}
