/**
Copyright (c) 2015 The PipeScript Contributors (see AUTHORS)
Licensed under the MIT license.
**/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/bytestreams"
	"github.com/connectordb/pipescript/transforms"

	"github.com/codegangsta/cli"
)

func getReader(s string) (*os.File, error) {
	if s == "STDIN" || s == "" {
		return os.Stdin, nil
	}
	return os.Open(s)
}

func getWriter(s string) (*os.File, error) {
	if s == "STDOUT" || s == "" {
		return os.Stdout, nil
	}
	return os.Create(s)
}

func runner(c *cli.Context, str string) {

	r, err := getReader(c.String("input"))
	if err != nil {
		log.Fatal(err)
	}
	w, err := getWriter(c.String("output"))
	if err != nil {
		log.Fatal(err)
	}
	// Now get the pipescript
	s, err := pipescript.Parse(str)
	if err != nil {
		log.Fatal(fmt.Errorf("%s\nIf using bash, make sure that '$' has spaces around it so that it is not mistaken for a bash variable.\n", err.Error()))
	}

	// Now set up the datapoint reader
	var dpr pipescript.DatapointIterator
	switch c.String("ifmt") {
	case "dp":

		dpr, err = bytestreams.NewDatapointReader(r)
		if err != nil {
			log.Fatal(err)
		}
	case "json":
		dpr, err = bytestreams.NewJSONDatapointReader(r, c.String("timestamp"), c.Bool("notimestamp"))
		if err != nil {
			log.Fatal(err)
		}
	case "csv":
		dpr, err = bytestreams.NewCSVDatapointReader(r, c.String("timestamp"), c.Bool("notimestamp"))
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unrecognized input format")
	}

	s.SetInput(dpr)

	// Now set the output json stream writer
	var jr io.Reader
	if !c.Bool("minify") {
		jr, err = bytestreams.NewJsonReader(s, "[\n", ",\n", "\n]", "", "  ")
	} else {
		jr, err = bytestreams.NewJsonArrayReader(s)
	}
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(w, jr); err != nil {
		log.Fatal(err)
	}

	// End with newline on stdout
	if w == os.Stdout {
		w.WriteString("\n")
	}

}

func main() {
	transforms.Register()

	app := cli.NewApp()
	app.Name = "pipes"
	app.Usage = "Run the PipeScript data analysis engine on your given datasets."
	app.Version = pipescript.Version

	app.Copyright = "This software is available under the MIT license"

	app.Usage = "pipes <transform or command> [command options]"

	transformArray := make([]cli.Command, 0, len(pipescript.TransformRegistry)+1)

	// We add the dump markdown option
	transformArray = append(transformArray, cli.Command{
		Name:  "dumpdocs",
		Usage: "dumpdocs dumps the markdown documentation files for all transforms to the given directory.",
		Action: func(c *cli.Context) error {
			p, err := filepath.Abs(c.Args().First())
			if err != nil {
				log.Fatal(err)
			}
			os.MkdirAll(p, 0777)
			for key := range pipescript.TransformRegistry {
				t := pipescript.TransformRegistry[key]
				docs := t.Documentation
				if docs == "" {
					docs = "*This transform is currently undocumented. You can help out by [adding documentation](https://github.com/connectordb/pipescript/tree/master/resources/docs/transforms)*"
				}
				title := "# " + key + "\n\n"
				err = ioutil.WriteFile(filepath.Join(p, key+".md"), []byte(title+docs), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}
			return nil
		},
	})

	for key := range pipescript.TransformRegistry {
		t := pipescript.TransformRegistry[key]
		desc := t.Description
		if len(desc) > 60 {
			desc = desc[0:60] + "..."
		}
		transformArray = append(transformArray, cli.Command{
			Name:  t.Name,
			Usage: desc,
			Action: func(c *cli.Context) error {
				b, err := json.MarshalIndent(t, "", "  ")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s\n", string(b))
				return nil
			},
		})
	}

	app.Commands = []cli.Command{
		{
			Name:    "transforms",
			Aliases: []string{"t"},
			Usage:   "Documentation for specific transforms",
			Action: func(c *cli.Context) error {
				s, _ := json.Marshal(pipescript.TransformRegistry)
				fmt.Printf("%s", string(s))
				return nil
			},
			Subcommands: transformArray,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "Run a transform.",
			Action: func(c *cli.Context) error {
				runner(c, c.Args().First())
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Value: "STDIN",
					Usage: "The input file to perform analysis on.",
				},
				cli.StringFlag{
					Name:  "output,o",
					Value: "STDOUT",
					Usage: "File to write as output",
				},
				cli.BoolFlag{
					Name:  "minify,m",
					Usage: "Whether to indent the json for easy reading",
				},
				cli.BoolFlag{
					Name:  "notimestamp",
					Usage: "If set, does not look for a timestamp field in input",
				},
				cli.StringFlag{
					Name:  "ifmt",
					Value: "dp",
					Usage: "The data format to use for input data (dp,csv,json)",
				},
				cli.StringFlag{
					Name:  "timestamp",
					Value: "",
					Usage: "Allows to explicitly set the field name to use for timestamp values",
				},
			},
		},
	}

	app.Run(os.Args)

}
