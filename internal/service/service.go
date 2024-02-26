package service

import (
	"log"
	"wb-test-app-2-3/internal/models"
	"wb-test-app-2-3/internal/repository"
	"wb-test-app-2-3/internal/utils"
)

func GetProductInfo(productID int) {
	// Запрашиваем информацию о продукте из базы данных
	rows, err := repository.GetProductInfo(productID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Обрабатываем результат запроса
	var product models.Product
	var categories []models.Category

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&product.ID, &product.Name, &product.Mark, &category.ID, &category.Name); err != nil {
			log.Println(err)
		}
		categories = append(categories, category)
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
