package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

//Function that upserts the schemas
func CreateEntitiesSchemas(db *dgo.Dgraph) {
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
				 make: [uid] .

				 transactionId: string @index(exact) .
				 ip: string @index(exact) .
				 os: string .
				 includes: [uid] .



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
					make
				 }
				 
				 type Transaction {
					 transactionId
					 ip
					 os
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
