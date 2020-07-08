package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	idtoken "github.com/yuzuy/firebase-idtoken"
	"google.golang.org/api/option"
)

var (
	apiKey    = flag.String("apikey", "", "specify api key of your firebase project")
	projectID = flag.String("p", "", "specify your firebase project id. by default read from application default credentials")
	credFile  = flag.String("credfile", "", "specify your firebase service account credential filepath. by default read from GOOGLE_APPLICATION_CREDENTIALS or FIREBASE_CONFIG")
)

var existCode int

func main() {
	run()
	os.Exit(existCode)
}

func run() {
	flag.Parse()
	if flag.NArg() < 1 {
		report(errors.New("please specify uid"))
		return
	}

	ctx := context.Background()

	var firebaseOpt option.ClientOption
	var itkOpt option.ClientOption
	var config *firebase.Config
	if credFile != nil {
		firebaseOpt = option.WithCredentialsFile(*credFile)
	}
	if apiKey != nil {
		itkOpt = option.WithAPIKey(*apiKey)
	}
	if projectID != nil {
		config = &firebase.Config{ProjectID: *projectID}
	}

	app, err := firebase.NewApp(ctx, config, firebaseOpt)
	if err != nil {
		report(err)
		return
	}
	client, err := app.Auth(ctx)
	if err != nil {
		report(err)
		return
	}

	uid := flag.Arg(0)
	token, err := idtoken.Generate(ctx, client, uid, itkOpt)
	if err != nil {
		report(err)
		return
	}

	fmt.Println(token)
}

func report(err error) {
	fmt.Fprintf(os.Stderr, "firebase-idtoken-gen: %s\n", err)
	existCode = 1
}
