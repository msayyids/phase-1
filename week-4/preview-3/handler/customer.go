package handler

import (
	"context"
	"fmt"
	"log"
)

func (h *Handler) Coustomer() {
	query := `SELECT customers.name, COUNT(orders.customer_id) AS total_order
	FROM customers
	JOIN orders ON orders.customer_id = customers.id
	GROUP BY customers.name
	HAVING total_order > 1
	ORDER BY total_order DESC
	LIMIT 1;;`

	ctx := context.Background()
	rows, err := h.DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s %-20s\n", "customer name", "total order")

	for rows.Next() {
		var customerName string
		var totalOrder int
		err := rows.Scan(&customerName, &totalOrder)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-20s %-20v\n", customerName, totalOrder)

	}
}
