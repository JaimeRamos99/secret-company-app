package database

import (
  "log"
  "context"
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
  res, err := txn.Do(ctx, req)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return res
}
