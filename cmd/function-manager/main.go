///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/config"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/entity-store"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager/gen/restapi"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager/gen/restapi/operations"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/openfaas"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/openwhisk"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/runner"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/secretinjector"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/functions/validator"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/trace"
)

var drivers = map[string]func() functions.FaaSDriver{
	"openfaas": func() functions.FaaSDriver {
		faas, err := openfaas.New(&openfaas.Config{
			Gateway:       config.Global.OpenFaas.Gateway,
			ImageRegistry: config.Global.OpenFaas.ImageRegistry,
			RegistryAuth:  config.Global.OpenFaas.RegistryAuth,
		})
		if err != nil {
			log.Fatalf("Error starting OpenFaaS driver: %+v", err)
		}
		return faas
	},
	"openwhisk": func() functions.FaaSDriver {
		faas, err := openwhisk.New(&openwhisk.Config{
			AuthToken: config.Global.Openwhisk.AuthToken,
			Host:      config.Global.Openwhisk.Host,
			Insecure:  true,
		})
		if err != nil {
			log.Fatalf("Error getting OpenWhisk client: %+v", err)
		}
		return faas
	},
}

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
	boltdb.Register()
}

var debugFlags = struct {
	DebugEnabled   bool `long:"debug" description:"Enable debugging messages"`
	TracingEnabled bool `long:"trace" description:"Enable tracing messages (enables debugging)"`
}{}

func configureFlags() []swag.CommandLineOptionsGroup {
	return []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Function manager Flags",
			LongDescription:  "",
			Options:          &functionmanager.FunctionManagerFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "Debug options",
			LongDescription:  "",
			Options:          &debugFlags,
		},
	}
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "2.0")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewFunctionManagerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Function Manager"
	parser.LongDescription = "This is the API server for the serverless function manager service.\n"

	optsGroups := configureFlags()
	for _, optsGroup := range optsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if debugFlags.DebugEnabled {
		log.SetLevel(log.DebugLevel)
	}
	if debugFlags.TracingEnabled {
		log.SetLevel(log.DebugLevel)
		trace.Enable()
	}

	config.Global = config.LoadConfiguration(functionmanager.FunctionManagerFlags.Config)

	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{functionmanager.FunctionManagerFlags.DbFile},
		&store.Config{
			Bucket:            "function",
			ConnectionTimeout: 1 * time.Second,
			PersistConnection: true,
		},
	)
	if err != nil {
		log.Fatalf("Error creating/opening the entity store: %v", err)
	}
	faas := drivers[functionmanager.FunctionManagerFlags.Faas]()
	handlers := &functionmanager.Handlers{
		FaaS: faas,
		Runner: runner.New(&runner.Config{
			Faas:           faas,
			Validator:      validator.New(),
			SecretInjector: secretinjector.New(functionmanager.SecretStoreClient()),
		}),
		Store:     entitystore.New(kv),
		ImgClient: functionmanager.ImageManagerClient(),
	}

	handlers.ConfigureHandlers(api)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
