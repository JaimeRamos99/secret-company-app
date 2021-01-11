package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

//Function that upserts the schemas
func CreateEntitiesSchemas(db *dgo.Dgraph) {
	//first delete other schemas and data
	DeleteAll(db)

	//Schemas to create
	op := api.Operation{
		Schema: `
				 date: string @index(exact) .

				 productId: string @index(exact) .
				 productName: string .
				 productPrice: string .

				 userId: string @index(exact) .
				 userName: string .
				 userAge: int .
				 madeBy: [uid] @reverse .

				 transactionId: string @index(exact) .
				 ip: string @index(exact) .
				 device: string .
				 includes: [uid] @reverse .



				 type Date {
				 	date
				 }

				 type Product {
					productId
					productName
					productPrice
				 }

				 type User {
					userId
					userName
					userAge
				 }
				 
				 type Transaction {
					 madeBy
					 transactionId
					 ip
					 device
					 includes
				 }
				`,
	}

	ctx := context.Background()
	if err := db.Alter(ctx, &op); err != nil {
		log.Fatal(err)
		return
	}
	log.Print("Schemas created succesfully")
}
