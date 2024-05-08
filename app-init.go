package main

//go:generate go run app-init.go

import instanceGen "github.com/skeletonkey/lib-instance-gen-go/app"

func main() {
	goVersion := "1.22"
	app := instanceGen.NewApp("watch-my-ip", "app")
	app.
		WithGoVersion(goVersion).
		WithConfig().
		WithPackages("ip").
		// WithDependencies(
		// "github.com/labstack/echo/v4",
		// "github.com/rabbitmq/amqp091-go",
		// ).
		WithGithubWorkflows("linter", "test").
		WithMakefile()
}
