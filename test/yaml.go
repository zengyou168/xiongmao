/**
 *
 * @Author: ZengYou
 * @Date: 2024/7/23
 */
package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gopkg.in/yaml.v3"
	"os"
	"panda/application"
)

func main() {
	data, err := os.ReadFile("application/application.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	var config application.Application
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("Parsed YAML data: %#v\n", config)
}
