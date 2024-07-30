package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/knakazawa99/gam-sdk-go/gamwsdl"
)

func main() {
	managementFilePath := "gamwsdl/generate_management.json"
	file, err := os.ReadFile(managementFilePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var generateManagement gamwsdl.GenerateManagement
	if err := json.Unmarshal(file, &generateManagement); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Generate Go code from each WSDL
	for i := range generateManagement.Services {
		service := &generateManagement.Services[i]
		_, err := os.Stat(service.GenerateDestination)
		if !os.IsNotExist(err) {
			if generateManagement.Version == service.Version {
				continue
			}
		}
		if err := gamwsdl.GenerateGoCodeFromGAMWSDL(service.URL, service.GenerateDestination); err != nil {
			fmt.Println(err.Error())
		}
		log.Printf("Generated %s\n", service.GenerateDestination)
		service.Version = generateManagement.Version
		service.LastUpdatedAt = time.Now().UTC()
	}

	file, err = json.MarshalIndent(generateManagement, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := os.WriteFile(managementFilePath, file, 0644); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
