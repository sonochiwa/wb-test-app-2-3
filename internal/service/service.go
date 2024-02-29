package service

import (
	"log"
	"strings"
	"wb-test-app-2-3/internal/models"
	"wb-test-app-2-3/internal/repository"
	"wb-test-app-2-3/internal/utils"
)

func GetProductInfo(productID int) {
	// Запрашиваем информацию о продукте из базы данных
	row := repository.GetProductInfo(productID)

	// Обрабатываем результат запроса
	var product models.Product
	var categoriesString string

	if err := row.Scan(&product.ID, &product.Name, &product.Mark, &categoriesString); err != nil {
		log.Println(err)
	}

	// Преобразование строки категорий в слайс строк
	categories := strings.Split(categoriesString[1:len(categoriesString)-1], ",")
	for i := range categories {
		categories[i] = strings.TrimSpace(categories[i])
	}

	// Выводим информации о продукте и его категориях
	utils.PrintProductInfo(product, categories)
}

func GetProductsByMark(markType string) {
	// Запрашиваем информацию о продукте из базы данных
	rows, err := repository.GetProductsByMark(markType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Mark, &product.Category); err != nil {
			log.Println(err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	// Выводим информации о минимальных и максимальных рейтингах продуктов в категории
	utils.PrintProductsByMark(products, markType)
}
