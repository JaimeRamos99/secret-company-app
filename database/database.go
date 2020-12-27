package database

import (
  "log"
  "context"
  "github.com/dgraph-io/dgo"
  "github.com/dgraph-io/dgo/protos/api"
  "google.golang.org/grpc"
)

type CancelFunc func()

func NewClient() (*dgo.Dgraph, CancelFunc) {
	// Dial a gRPC connection.
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	} else {
    log.Print("Established connection")
  }
  dc := api.NewDgraphClient(conn)
  dg := dgo.NewDgraphClient(dc)

  //Return a function to close the connection
	return dg, func() {
      if err := conn.Close(); err != nil {
        log.Printf("Error while closing connection:%v", err)
      } else {
        log.Println("conexión cerrada")
      }
    }
}

//Function that Drop all the schemas
func DeleteAll(db *dgo.Dgraph){
  op := api.Operation{DropAll: true}
  ctx := context.Background()
  if err := db.Alter(ctx, &op); err != nil {
    log.Fatal(err)
  }
}

//Function that upserts the schema that tracks the dates that have been updated
func CreateSchemaUploadDates(){

}
