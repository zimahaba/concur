package main

import (
	"concur/cmd"
	"concur/pkg"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	pkg.ConnectDB()
	cmd.Execute()
	pkg.DB.Close()
}
