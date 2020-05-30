package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/engelsjk/arcgis2geojson"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "arcgis2geojson",
		Usage: "convert arcgis json to geojson",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "id",
				Usage: "id attribute key (default `OBJECTID`)",
			},
		},
		Action: func(c *cli.Context) error {
			idAttributeKey := strings.ToUpper(c.String("id"))
			return Run(c.Args(), idAttributeKey)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Run(args cli.Args, idAttributeKey string) error {
	var data []byte
	var err error
	switch args.Len() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
	case 1:
		data, err = ioutil.ReadFile(args.Get(0))
		if err != nil {
			return err
		}
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
	}

	b, err := arcgis2geojson.Convert(data, idAttributeKey)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
