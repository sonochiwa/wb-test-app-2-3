package repository

import (
	"database/sql"
	"fmt"
)

// GetProductInfo - возвращает информацию о продукте и его категориях из базы данных
func GetProductInfo(productID int) *sql.Row {
	return DB.QueryRow(getProductInfoQuery, productID)
}

// GetProductsByMark - возвращает продукты с заданным рейтингом из базы данных
func GetProductsByMark(markType string) (*sql.Rows, error) {
	var query string

	switch markType {
	case "Max":
		query = getProductsByMaxMark
	case "Min":
		query = getProductsByMinMark
	default:
		return nil, fmt.Errorf("invalid mark type: %s", markType)
	}

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
