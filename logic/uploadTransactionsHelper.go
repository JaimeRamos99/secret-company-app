package logic

import (
	json "encoding/json"
	log "log"
	strings "strings"
	unicode "unicode"

	utils "github.com/JaimeRamos99/prueba-truora-2/utils"
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
)

//Given a transaction string, this func converts it
//into a valid Transaction struct
func splitTransactions(tr string) structs.Transaction {
	allowed_runes := utils.LoopDigits()
	acum := ""

	//creating a mirror string, except for the especial rune
	//instead a _ is added, in order to split later
	for _, rune := range tr {
		_, is_digit_or_special := allowed_runes[string(rune)]
		if unicode.IsLetter(rune) || is_digit_or_special {
			acum = acum + string(rune)
		} else {
			acum = acum + "_"
		}
	}

	//splitting transactions attrs
	tran_splitted := strings.Split(acum, "_")
	len_prods_str := len(tran_splitted[4])
	products := strings.Split(tran_splitted[4][1:len_prods_str-1], ",")

	//Creating the Products array struct
	var products_Array []structs.ProductId
	for _, product := range products {
		prod := *structs.NewProductId(product)
		products_Array = append(products_Array, prod)
	}

	//instance of Transaction
	tran := *structs.NewTransaction(tran_splitted[0], tran_splitted[1], tran_splitted[2], tran_splitted[3], products_Array)
	return tran
}

//creates an array of structs valid to create an edge between a transaction and products
func createProductsStruct(products []structs.ProductId, prods_map map[string]string) []structs.ProductUid {
	var productUids []structs.ProductUid
	for _, prod := range products {
		new_prod := *structs.NewProductUId(prods_map[prod.ProductId])
		productUids = append(productUids, new_prod)
	}
	return productUids
}

//creates the complete json provided to dgrpah to create the transaction nodes
//and the edges with the users and products
func CreateTransactionJson(transactions []structs.Transaction, usrs_map map[string]string, prods_map map[string]string) []byte {
	var transaction_mutation []structs.TransactionMutation
	for _, tr := range transactions {
		products := createProductsStruct(tr.Products, prods_map)
		made_by := *structs.NewMadeBy(usrs_map[tr.BuyerId])
		new_transaction := *structs.NewTransactionMutation(
			made_by,
			tr.IdTransaction,
			tr.Ip,
			tr.Device,
			products,
		)
		transaction_mutation = append(transaction_mutation, new_transaction)
	}
	//fmt.Printf("%+v/n", transaction_mutation)
	transactions_json, errorr := json.Marshal(transaction_mutation)
	if errorr != nil {
		log.Fatal(errorr)
	}
	return transactions_json
}
