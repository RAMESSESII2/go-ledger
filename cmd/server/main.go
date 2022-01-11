package main

import (
	"flag"
	"fmt"

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
	fmt.Print("Inside server's main")
	// dbPathD := (*dbUserFlag) + ":" + (*dbPasswordFlag) + "@tcp(127.0.0.1:3306)/ledgerDB?charset=utf8mb4&parseTime=True&loc=Local"
	dbPathD := "root:secret@tcp(db:3306)/ledger?charset=utf8&parseTime=True&loc=Local"
	server.StartServer(dbPathD, ":"+*addressFlag)
}
