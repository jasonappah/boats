package main

import (
	"log"
	"boats/lib"
)

func main() {
    app := lib.InitPocketbase(true)

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}