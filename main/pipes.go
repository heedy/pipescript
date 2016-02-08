/**
Copyright (c) 2015 The PipeScript Contributors (see AUTHORS)
Licensed under the MIT license.
**/

package main

import (
	"io"
	"log"
	"os"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/bytestreams"
	"github.com/connectordb/pipescript/transforms"

	"github.com/codegangsta/cli"
)

func getReader(s string) (io.Reader, error) {
	if s == "STDIN" {
		return os.Stdin, nil
	}
	return os.Open(s)
}

func getWriter(s string) (io.Writer, error) {
	if s == "STDOUT" {
		return os.Stdout, nil
	}
	return os.Create(s)
}

func main() {
	transforms.Register()

	app := cli.NewApp()
	app.Name = "pipes"
	app.Usage = "Run the PipeScript data analysis engine on your given datasets."
	app.Version = pipescript.Version

	app.Copyright = "This software is available under the MIT license"

	app.Flags = []cli.Flag{
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
			Name:  "pretty,p",
			Usage: "Whether to indent the json for easy reading",
		},
		cli.BoolFlag{
			Name:  "list,l",
			Usage: "List the available transforms",
		},
	}

	app.CommandNotFound = func(c *cli.Context, str string) {

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
			log.Fatal(err)
		}

		// Now set up the datapoint reader
		dpr, err := bytestreams.NewDatapointReader(r)
		if err != nil {
			log.Fatal(err)
		}
		s.SetInput(dpr)

		// Now set the output json stream writer
		var jr io.Reader
		if c.Bool("pretty") {
			jr, err = bytestreams.NewJsonReader(s, "[\n", ",\n", "\n]", "", "\t")
		} else {
			jr, err = bytestreams.NewJsonArrayReader(s)
		}
		if err != nil {
			log.Fatal(err)
		}
		if _, err = io.Copy(w, jr); err != nil {
			log.Fatal(err)
		}

	}

	app.Run(os.Args)

}
