package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func Recommendation(db *dgo.Dgraph) *api.Response {

	ctx := context.Background()

	//Instance of a transaction
	txn := db.NewTxn()
	defer txn.Discard(ctx)

	//Looking in the DB if the data for that date has been uploaded
	query :=
		`
		{
			first_stage(func:has(transactionId)) {
				  TR as uid
			  }
			
			second_stage(func: uid(TR)){
			  includes {
				conteo as contar: count(~includes)
			  }
			}
			
			third_stage(func: uid(TR)){
			  includes(orderdesc: val(conteo), first:1) {
				conteo: count(~includes)
			  }
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
