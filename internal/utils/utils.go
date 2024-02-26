package utils

import (
	"fmt"
	"strings"
	"wb-test-app-2-3/internal/models"
)

// PrintProductInfo - выводит информацию о продукте и его категориях
func PrintProductInfo(product models.Product, categories []models.Category) {
	var categoryNames []string
	for _, category := range categories {
		categoryNames = append(categoryNames, strings.ToLower(category.Name))
	}

	fmtCategories := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(categoryNames)), ", "), "[]")

	fmt.Println("Идентификатор:\t", product.ID)
	fmt.Println("Название:\t", product.Name)
	fmt.Printf("Рейтинг:\t %d/10\n", product.Mark)
	fmt.Println("Категории:\t", fmtCategories)
}

func PrintProductsByMark(products []models.Product, markType string) {
	fmt.Printf("%s - ", markType)
	for _, product := range products {
		fmt.Printf("%s: %s [рейтинг %d/10]; ", product.Category, product.Name, product.Mark)
	}
	fmt.Println()
}
