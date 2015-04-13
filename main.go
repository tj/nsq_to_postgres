package main

import "github.com/tj/nsq_to_postgres/handler"
import "github.com/tj/nsq_to_postgres/client"
import "github.com/segmentio/go-queue"
import "github.com/tj/go-gracefully"
import "github.com/tj/docopt"
import "gopkg.in/yaml.v2"
import "io/ioutil"
import "log"

var Version = "0.0.1"

const Usage = `
  Usage:
    nsq_to_postgres --config file
    nsq_to_postgres -h | --help
    nsq_to_postgres --version

  Options:
    -c, --config file   configuration file path
    -h, --help          output help information
    -v, --version       output version

`

type Config struct {
	Postgres *client.Config
	Nsq      map[string]interface{}
}

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Printf("starting nsq_to_postgres version %s", Version)

	// Read config
	file := args["--config"].(string)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	// Unmarshal config
	config := new(Config)
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	// Validate config
	err = config.Postgres.Validate()
	if err != nil {
		log.Fatalf("configuration error: %s", err)
	}

	// Apply nsq config
	c := queue.NewConsumer("", "")

	for k, v := range config.Nsq {
		c.Set(k, v)
	}

	// Connect
	log.Printf("connecting to postgres")
	db, err := client.New(config.Postgres)
	if err != nil {
		log.Fatalf("error connecting: %s", err)
	}

	// Bootstrap with table
	err = db.Bootstrap()
	if err != nil {
		log.Printf("error bootstrapping: %s", err)
	}

	// Start consumer
	log.Printf("starting consumer")
	c.Start(handler.New(db))
	gracefully.Shutdown()
	log.Printf("stopping consumer")
	c.Stop()

	log.Printf("bye :)")
}
