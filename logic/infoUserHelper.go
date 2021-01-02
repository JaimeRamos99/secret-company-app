package logic

import (
	structs "github.com/JaimeRamos99/prueba-truora-2/utils/structs"
)

func ThreeBestSellers(trans *structs.Result) []structs.TopProduct {
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
	var podium []structs.TopProduct
	first := *structs.NewTopProduct(first_name, first_score)
	second := *structs.NewTopProduct(second_name, second_score)
	third := *structs.NewTopProduct(third_name, third_score)
	podium = append(podium, first)
	podium = append(podium, second)
	podium = append(podium, third)
	return podium
}
