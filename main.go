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

	info, ok := debug.ReadBuildInfo()
	if ok && info != nil && info.Main.Version != "" {
		version = info.Main.Version
	}
	if *jsonPtr {
		by, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(by))
	} else {
		log.Printf("printver: %s", strings.TrimSpace(version))
	}
}
