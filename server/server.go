package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/LukianLeopoldo013/micro-service-email/jsonapi"
	"github.com/LukianLeopoldo013/micro-service-email/mdb"
	"github.com/alexflint/go-arg"
)

var args struct {
	DbPath   string `args:"env:MAILINGLIST_DB"`
	BindJson string `args:"env:MAILINGLIST_BIND_JSON"`
}

func main() {
	arg.MustParse(&args)

	if args.DbPath == "" {
		args.DbPath = "list.db"
	}
	if args.BindJson == "" {
		args.BindJson = ":8080"
	}

	log.Printf("using database %v\n", args.DbPath)

	db, err := sql.Open("sqlite3", args.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mdb.TryCreate(db)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		log.Print("starting JSON API server...\n")
		jsonapi.Server(db, args.BindJson)
		wg.Done()
	}()

	wg.Wait()

}
