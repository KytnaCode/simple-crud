package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kytnacode/simple-crud/internal/db"
)

const usage = `Usage:

cli <command>

Commands:

database-init
`

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load `.env` file: %v", err.Error())
		return
	}

	if len(os.Args) <= 1 {
		fmt.Fprint(os.Stderr, usage)
		return
	}

	command := os.Args[1]
	if command == "database-init" {
		value, ok := os.LookupEnv("APP_DATABASE_FILE")
		if !ok {
			fmt.Fprintln(os.Stderr, "APP_DATABASE_FILE needs to be set")
			fmt.Fprint(os.Stderr, err)

			return
		}

		err = db.Init(value)
		if err != nil {
			fmt.Println(err)
		}

		database := db.Get()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err = db.PopulateDb(ctx, database)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
	}
}
