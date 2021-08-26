package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	idtoken "github.com/yuzuy/firebase-idtoken-gen"
)

var (
	apiKey    = flag.String("apikey", "", "specify api key of your firebase project")
	projectID = flag.String("p", "", "specify your firebase project id. by default read from application default credentials")
	credFile  = flag.String("credfile", "", "specify your firebase service account credential filepath. by default read from GOOGLE_APPLICATION_CREDENTIALS or FIREBASE_CONFIG")
)

var exitCode int

func main() {
	run()
	os.Exit(exitCode)
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
	if *credFile != "" {
		firebaseOpt = option.WithCredentialsFile(*credFile)
	}
	if *apiKey != "" {
		itkOpt = option.WithAPIKey(*apiKey)
	}
	if *projectID != "" {
		config = &firebase.Config{ProjectID: *projectID}
	}

	var app *firebase.App
	var err error
	if firebaseOpt == nil {
		app, err = firebase.NewApp(ctx, config)
	} else {
		app, err = firebase.NewApp(ctx, config, firebaseOpt)
	}

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
	exitCode = 1
}
