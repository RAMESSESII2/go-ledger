package main

import (
	"flag"

	"github.com/RAMESSESII2/go-ledger/server"
)

// const DNS = "root:World@5261@tcp(127.0.0.1:3306)/ledgerDB?charset=utf8mb4&parseTime=True&loc=Local"

var (
	addressFlag    = flag.String("port", "9000", "HTTP server address")
	dbUserFlag     = flag.String("username", "root", "mysql username")
	dbPasswordFlag = flag.String("password", "World@5261", "user's mysql password")
)

func main() {
	flag.Parse()
	dbPath := (*dbUserFlag) + ":" + (*dbPasswordFlag) + "@tcp(127.0.0.1:3306)/ledgerDB?charset=utf8mb4&parseTime=True&loc=Local"
	server.StartServer(dbPath, ":"+*addressFlag)
}
