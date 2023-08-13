package main

import (
	"encoding/json"
	"flag"
	"log"
	"runtime/debug"
	"strings"
)

var (
	version = "(devel)"
	jsonPtr = flag.Bool("json", false, "print build info as json")
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	flag.Parse()

	buildInfo, ok := debug.ReadBuildInfo()
	if ok && buildInfo != nil && buildInfo.Main.Version != "" {
		version = buildInfo.Main.Version
	}
	if *jsonPtr {
		by, err := json.MarshalIndent(buildInfo, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(by))
	} else {
		log.Printf("printver: %s", strings.TrimSpace(version))
	}
}
