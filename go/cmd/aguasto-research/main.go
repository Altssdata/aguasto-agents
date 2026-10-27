package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	aguasto "github.com/Altssdata/aguasto-agents/go"
	"os"
	"strings"
)

func run() error {
	flags := flag.NewFlagSet("aguasto-research", flag.ContinueOnError)
	base := flags.String("base-url", "https://www.aguasto.com", "Public HTTPS origin; HTTP loopback for local tests")
	flags.Usage = func() {
		fmt.Fprintln(flags.Output(), "Usage: aguasto-research [--base-url ORIGIN] search QUERY | catalog | calculate NAME PARAMETERS_JSON | fixtures\nPublic research and educational estimates only. No GPU provisioning or inference.")
	}
	if err := flags.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return err
	}
	args := flags.Args()
	if len(args) == 0 {
		flags.Usage()
		return nil
	}
	if args[0] == "fixtures" && len(args) == 1 {
		_, err := os.Stdout.Write(aguasto.OfflineFixtures())
		return err
	}
	client, err := aguasto.NewClient(*base)
	if err != nil {
		return err
	}
	var result map[string]any
	switch args[0] {
	case "catalog":
		if len(args) != 1 {
			return fmt.Errorf("catalog takes no arguments")
		}
		result, err = client.Catalog(context.Background())
	case "search":
		if len(args) < 2 {
			return fmt.Errorf("search requires a query")
		}
		result, err = client.Search(context.Background(), strings.Join(args[1:], " "))
	case "calculate":
		if len(args) != 3 {
			return fmt.Errorf("calculate requires NAME and PARAMETERS_JSON")
		}
		var parameters map[string]float64
		if err := json.Unmarshal([]byte(args[2]), &parameters); err != nil {
			return err
		}
		result, err = client.Calculate(context.Background(), args[1], parameters)
	default:
		return fmt.Errorf("unknown command; use --help")
	}
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
func main() {
	if err := run(); err != nil {
		_ = json.NewEncoder(os.Stderr).Encode(map[string]string{"error": err.Error()})
		os.Exit(1)
	}
}
