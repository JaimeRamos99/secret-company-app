package database

import (
  "log"
  "context"
  "encoding/json"
  "github.com/dgraph-io/dgo/v200"
  "github.com/dgraph-io/dgo/v200/protos/api"
)

type Date struct {
  Uid   string   `json:"uid"`
	Date  string   `json:"name,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

func CheckDate(db *dgo.Dgraph, date string) *api.Response  {

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
    Vars: map[string]string{"$input": date},
  }

  //Commiting the transaction and returning the result
  res, err := txn.Do(ctx, req)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return res
}

func AddUploadDate(db *dgo.Dgraph, date string){
    date_struct := Person{
    	Date:  date,
    	DType: []string{"Date"},
    }
    date_json, err := json.Marshal(date_struct)
    if err != nil {
    	log.Fatal(err)
    }

    mu := &api.Mutation{
    	SetJson: date_json,
    }
    res, err := txn.Mutate(ctx, mu)
    if err != nil {
	     log.Fatal(err)
    }
}
