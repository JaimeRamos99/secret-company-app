package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func GetAllProducts(db *dgo.Dgraph) *api.Response {
	ctx := context.Background()

	//Instance of a transaction
	txn := db.NewTxn()
	defer txn.Discard(ctx)

	//Looking in the DB if the data for that date has been uploaded
	query :=
		`{
		 	products(func:has(productId)){
				 uid
				productId
			}
		 }
	    `

	req := &api.Request{
		Query: query,
	}

	//Commiting the transaction and returning the result
	res, err := txn.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}
