package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wolfeidau/go-buildkite/buildkite"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	apiToken = kingpin.Flag("token", "API token").Required().String()
	org      = kingpin.Flag("org", "Orginization slug").Required().String()
	debug    = kingpin.Flag("debug", "Enable debugging").Bool()
)

func main() {
	kingpin.Parse()

	config, err := buildkite.NewTokenConfig(*apiToken, *debug)

	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	client := buildkite.NewClient(config.Client())

	projects, _, err := client.Projects.List(*org, nil)

	if err != nil {
		log.Fatalf("list projects failed: %s", err)
	}

	data, err := json.MarshalIndent(projects, "", "\t")

	if err != nil {
		log.Fatalf("json encode failed: %s", err)
	}

	fmt.Fprintf(os.Stdout, "%s", string(data))
}
