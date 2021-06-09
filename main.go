package main

import (
	"context"
	"flag"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	flag.Parse()

	config, err := google.ConfigFromJSON(GetCredentials(), drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(getClient(config)))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	switch *Mode {
	case "upload":
		// uoload file
		Up(srv)
	case "download":
		// download file
		Down(srv)
	default:
		log.Fatalf("Unsupported mode.")
	}
}
