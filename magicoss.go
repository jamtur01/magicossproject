// Copyright 2018 James Turnbull. All Rights Reserved.
// This file is available under the Apache license.

package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type seqStringFlag []string

func (f *seqStringFlag) String() string {
	return fmt.Sprint(*f)
}

func (f *seqStringFlag) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		*f = append(*f, v)
	}
	return nil
}

type seqIntFlag []int

func (f *seqIntFlag) String() string {
	return fmt.Sprint(*f)
}

func (f *seqIntFlag) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		val, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*f = append(*f, val)
	}
	return nil
}

var (
	port    = flag.String("port", "8080", "HTTP port to listen on.")
	address = flag.String("interface", "", "Host or IP address on which to bind the HTTP listener")
	version = flag.Bool("version", false, "Print magicoss version information.")
)

var (
	Version = "v1"
	Revision = "749e81eac3aef0c6015fffb035f55d981ade1261"
	GoVersion = runtime.Version()
)

func buildInfo() string {
	return fmt.Sprintf("magicoss version %s git revision %s go version %s", Version, Revision, GoVersion)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", buildInfo())
		fmt.Fprintf(os.Stderr, "\nUsage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Println(buildInfo())
		os.Exit(1)
	}
}
