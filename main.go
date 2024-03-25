package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/provider"
)

var (
	// TODO set correct version
	//  milestone: 1

	// These are intended to be set in the build command
	version string
	address string
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if version == "" {
		version = "local-dev"
	}

	if address == "" {
		address = " hostname/namespace/type"
	}

	opts := providerserver.ServeOpts{
		Address: address,
		Debug:   debug,
	}
	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

// TODO Create a CONTRIBUTION.md doc
//  milestone: 2
//  include meta info such as
//  - [ ] how schema is used
//  - [ ] how the code generation works
//  - [ ] why so many extra files are committed to the repo
//  - [ ] diagram of the makefile workflow

// TODO Autogenerate CHANGELOG.md
//  milestone: 2
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#versioning-specification
//  Ref: https://developer.hashicorp.com/terraform/plugin/best-practices/versioning#changelog-specification
//  - [ ] Investigate how to add "chglog" friendly part to commit message
//  - [ ] Add git commit messages
//  - [ ] Autogenerate resource changes from provider-schema files

// TODO Investigate official code gen method
//  Ref: https://developer.hashicorp.com/terraform/plugin/code-generation/design
//  milestone: 6

// TODO Create node override system
//  To be able to add/change/remove parameters/attributes of the autogenerated
//  resource/datasource a system must be in place. This will allow fixing bugs
//  and problems that is inherently fundamental due to the conversion between
//  systems, coding languages, and schema formats.
//  milestone: 3

// TODO tools: build terraform resources: split up template into smaller files
//  milestone: 6
//  Consolidate named & global templates where practical

// TODO implement import state interface in resource generation template
//  milestone: 2
//  If we need to create a custom import function add validation for the interface

// TODO improve Readme and index.md
//  milestone: 1
