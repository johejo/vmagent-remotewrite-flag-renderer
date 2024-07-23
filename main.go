// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
)

var (
	config       string
	configFormat string

	vmagentSchema bool
)

func init() {
	flag.StringVar(&config, "config", "config.yaml", "Config file path JSON and YAML are supported")
	flag.StringVar(&configFormat, "config-format", "", `Config file format: "prometheus" or "vmagent"`)
	flag.BoolVar(&vmagentSchema, "vmagent-schema", false, "print vmagent config file json schema")
}

func main() {
	flag.Parse()
	if vmagentSchema {
		r := jsonschema.Reflector{
			RequiredFromJSONSchemaTags: true,
		}
		s := r.Reflect(VMAgentConfig{})
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		os.Exit(0)
	}
	s, err := run(config, configFormat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(s)
}

func run(config, configFormat string) (string, error) {
	if configFormat == "" {
		return "", errors.New("config format is not specified")
	}

	var f *os.File
	if config == "-" {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(config)
		if err != nil {
			return "", err
		}
		defer f.Close()
	}
	switch configFormat {
	case "prometheus":
		var cfg PrometheusConfig
		if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
			return "", err
		}
		return cfg.String(), nil
	case "vmagent":
		var cfg VMAgentConfig
		switch filepath.Ext(config) {
		case ".yaml", ".yml":
			if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
				return "", err
			}
		case ".json":
			if err := json.NewDecoder(f).Decode(&cfg); err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("unsupported file extension %s", config)
		}
		return cfg.String(), nil
	default:
		return "", fmt.Errorf("unsupported config format %s", configFormat)
	}
}

type builder struct {
	b *strings.Builder
}

func newBuilder() *builder {
	return &builder{b: &strings.Builder{}}
}

func (b *builder) W(format string, args ...any) {
	if b.b.Len() != 0 {
		b.b.WriteString(" ")
	}
	fmt.Fprintf(b.b, format, args...)
}

func (b *builder) String() string {
	return b.b.String()
}
