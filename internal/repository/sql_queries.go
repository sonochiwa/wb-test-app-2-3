package repository

const (
	getProductInfoQuery = `
		SELECT p.id, p.name, p.mark, array_agg(c.name) as categories
		FROM products p
				 JOIN products_categories pc ON p.id = pc.product_id
				 JOIN categories c ON pc.category_id = c.id
		WHERE p.id = $1
		GROUP BY p.id, p.name, p.mark;
	`

	getProductsByMaxMark = `
		SELECT p.id, p.name, p.mark, c.name AS category_name
		FROM products p
		JOIN products_categories pc ON p.id = pc.product_id
		JOIN categories c ON pc.category_id = c.id
		WHERE p.mark = (
			SELECT MAX(p2.mark)
			FROM products p2
			JOIN products_categories pc2 ON p2.id = pc2.product_id
			WHERE pc2.category_id = pc.category_id
		)
	`

	getProductsByMinMark = `
		SELECT p.id, p.name, p.mark, c.name AS category_name
		FROM products p
		JOIN products_categories pc ON p.id = pc.product_id
		JOIN categories c ON pc.category_id = c.id
		WHERE p.mark = (
			SELECT MIN(p2.mark)
			FROM products p2
			JOIN products_categories pc2 ON p2.id = pc2.product_id
			WHERE pc2.category_id = pc.category_id
		)
	`
)
