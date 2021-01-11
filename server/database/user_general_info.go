package database

import (
	context "context"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

func User_general_info(db *dgo.Dgraph, user_id string) *api.Response {

	ctx := context.Background()

	variables := map[string]string{"$user_id": user_id}
	//Looking in the DB if the data for that date has been uploaded
	query :=
		`
			query Info($user_id: string){
				info(func:eq(userId, $user_id)) {
					userId
					userAge
					userName
					transactions : ~madeBy{
						ip
						transactionId
						device
						includes{
							productPrice
							productName
							productId
						}
					}
				}
			}
   		`

	//Commiting the transaction and returning the result
	res, err := db.NewTxn().QueryWithVars(ctx, query, variables)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}
