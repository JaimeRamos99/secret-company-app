package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func Same_ips(db *dgo.Dgraph, user_id string) *api.Response {
	ctx := context.Background()

	//Instance of a transaction
	txn := db.NewTxn()
	defer txn.Discard(ctx)

	//Looking in the DB if the data for that date has been uploaded
	query :=
		`
			query SameIps($user_id: string){
				first_stage(func:eq(userId, $user_id)) {
					~madeBy{
						ips as ip
				    }
				}
			  
			    second_stage(func:eq(ip, val(ips))){
					transactionId
					ip
					madeBy {
						userId
						userName
					}
			    }
			}
   		`

	req := &api.Request{
		Query: query,
		Vars:  map[string]string{"$user_id": user_id},
	}

	//Commiting the transaction and returning the result
	res, err := txn.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}
