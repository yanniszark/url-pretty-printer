package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

type URLInfo struct {
	Scheme   string            `json:"scheme"`
	Host     string            `json:"host"`
	Path     string            `json:"path"`
	Query    map[string]string `json:"query"`
	Fragment string            `json:"fragment"`
}

func NewURLInfoFromURL(url *url.URL) *URLInfo {

	joinAllKeyValues := func(kvs map[string][]string) map[string]string {
		res := map[string]string{}
		for k, v := range kvs {
			res[k] = strings.Join(v, ",")
		}
		return res
	}

	return &URLInfo{
		Scheme:   url.Scheme,
		Host:     url.Host,
		Path:     url.Path,
		Query:    joinAllKeyValues(url.Query()),
		Fragment: url.Fragment,
	}
}

func pprint(v interface{}) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Fatalf("Pretty-print failed: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}

func main() {
	app := &cli.App{
		Name:  "url-pretty-printer",
		Usage: "Pretty-print a URL's parts",
		Action: func(c *cli.Context) error {
			urlString := c.Args().Get(0)
			if urlString == "" {
				return errors.New("Must specify a URL to print.")
			}
			url, err := url.Parse(urlString)
			if err != nil {
				return err
			}
			pprint(NewURLInfoFromURL(url))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
