package main

import (
	"fmt"
	"wb-test-app-2-3/internal/repository"
	"wb-test-app-2-3/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	// Закрываем подключение к БД в конце выполнения программы
	defer repository.DB.Close()

	// Получаем информацию о продукте с id=8
	service.GetProductInfo(8)

	fmt.Println()

	// Получаем продукты по минимальному и максимальному рейтингу
	service.GetProductsByMark("Min")
	service.GetProductsByMark("Max")
}
