package main

import (
	"log"

	"hex-arch-go-grpc/internal/adapters/app/api"
	"hex-arch-go-grpc/internal/adapters/framework/out/db"
	"hex-arch-go-grpc/internal/ports"

	gRPC "hex-arch-go-grpc/internal/adapters/framework/in/grpc"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var arithAdapter ports.ArithemeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	// init db: Any database or mock database that implements DbPort can be used here.
	// dbaseAdapter is a secondary actor and it is driven by the application.
	dbaseAdapter, err = db.NewAdapter("mysql", "admin:Admin123@tcp(staticdoc.ctmspwuawvbg.ap-northeast-1.rds.amazonaws.com:3306)/staticdoc")
	if err != nil {
		log.Fatalf("Failed to initiate database connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	// plug:
	appAdapter = api.NewAdapter(dbaseAdapter, arithAdapter)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
