package logic

import (
	"fmt"

	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
)

func ThreeBestSellers(trans *structs.Result) {
	first_score, second_score, third_score := 0, 0, 0
	first_name, second_name, third_name := "", "", ""
	unique_products := make(map[string]int)

	for _, group := range trans.Result {
		for _, prod := range group.Includes {
			unique_products[prod.ProductName] = prod.Score
		}
	}

	for key, element := range unique_products {
		if element > third_score {
			if element > second_score {
				if element > first_score {
					first_score = element
					first_name = key
				} else {
					second_score = element
					second_name = key
				}
			} else {
				third_score = element
				third_name = key
			}
		}
	}
	fmt.Println("1. ", first_name, first_score)
	fmt.Println("2. ", second_name, second_score)
	fmt.Println("3. ", third_name, third_score)
}
