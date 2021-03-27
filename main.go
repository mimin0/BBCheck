package main

import (
	"fmt"
	"log"

	helper "github.com/BBCheck/helper"
)

func main() {
	// reading YAML config and validation congig path
	cfg, err := helper.NewConfig("./config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	result := helper.GetRepos(cfg.Host, cfg.Repo, cfg.Project)

	for _, item := range result["values"].([]interface{}) {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Printf("%v\n", item.(map[string]interface{})["title"])
	}
	// fmt.Println("response Body:", string(body))
}
