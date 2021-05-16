package main

import (
	"log"
	"os"

	// application
	"hex/internal/application/api"
	"hex/internal/application/core/arithmetic"

	// adapters
	gRPC "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
)

func main() {
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	// core
	core := arithmetic.New()

	// NOTE: The application's right side port for driven
	// adapters, in this case, a db adapter.
	// Therefore the type for the dbAdapter parameter
	// that is to be injected into the NewApplication will
	// be of type DbPort
	applicationAPI := api.NewApplication(dbAdapter, core)

	// NOTE: We use dependency injection to give the grpc
	// adapter access to the application, therefore
	// the location of the port is inverted. That is
	// the grpc adapter accesses the hexagon's driving port at the
	// application boundary via dependency injection,
	// therefore the type for the applicaitonAPI parameter
	// that is to be injected into the gRPC adapter will
	// be of type APIPort which is our hexagons left side
	// port for driving adapters
	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()
}
