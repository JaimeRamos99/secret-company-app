package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
	grpc "google.golang.org/grpc"
)

type CancelFunc func()

func NewClient() (*dgo.Dgraph, CancelFunc) {
	// dial a gRPC connection.
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Established connection with Dgraph")
	}
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	// return a function to close the connection
	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		} else {
			log.Println("Connection closed")
		}
	}
}

// function that Drop all the data
func DeleteAll(db *dgo.Dgraph) {
	op := api.Operation{DropOp: api.Operation_ALL}
	ctx := context.Background()
	if err := db.Alter(ctx, &op); err != nil {
		log.Fatal(err)
	}
}
