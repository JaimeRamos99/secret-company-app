package database

import (
	context "context"
	json "encoding/json"
	log "log"

	dgo "github.com/dgraph-io/dgo/v200"
	api "github.com/dgraph-io/dgo/v200/protos/api"
)

type Date struct {
	Uid   string   `json:"uid"`
	Date  string   `json:"date,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

func CheckDate(db *dgo.Dgraph, date string) *api.Response {

	ctx := context.Background()

	//Instance of a transaction
	txn := db.NewTxn()
	defer txn.Discard(ctx)

	//Looking in the DB if the data for that date has been uploaded
	query :=
		`query all($input: string) {
      query(func: eq(date, $input)){
        date
      }
   }`

	req := &api.Request{
		Query: query,
		Vars:  map[string]string{"$input": date},
	}

	//Commiting the transaction and returning the result
	res, err := txn.Do(ctx, req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return res
}

func AddUploadDate(db *dgo.Dgraph, date string) {

	ctx := context.Background()

	//Instance of a transaction
	txn := db.NewTxn()
	defer txn.Discard(ctx)

	//Struct instance
	date_struct := Date{
		Date:  date,
		DType: []string{"Date"},
	}

	//parse struct to json format
	date_json, err := json.Marshal(date_struct)
	if err != nil {
		log.Fatal(err)
	}

	//Set the transaction json mode
	mu := &api.Mutation{
		SetJson: date_json,
	}

	//Commit the transaction
	req := &api.Request{CommitNow: true, Mutations: []*api.Mutation{mu}}
	_, er := txn.Do(ctx, req)
	if er != nil {
		log.Fatal(err)
	}
}
