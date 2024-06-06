package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	app := instanceGen.NewApp("watch-my-ip", "app")
	app.SetupApp(
		app.WithCodeOwners("* @skeletonkey"),
		app.WithConfig(),
		app.WithGithubWorkflows("linter", "test"),
		app.WithGoVersion("1.22"),
		app.WithMakefile(),
		app.WithPackages("ip"),
	).Generate()
}
