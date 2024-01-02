package main

import (
	"concur/cmd"
	"concur/pgk"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pgk.ConnectDB()
	cmd.Execute()
	pgk.DB.Close()
}
