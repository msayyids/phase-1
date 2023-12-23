package handler

import (
	"context"
	"fmt"
	"log"
)

func (h *Handler) TotalSales() {
	query := `select type,sum(price) as total
	from books
	join orders on orders.book_id = books.id
	GROUP BY type;`

	ctx := context.Background()
	rows, err := h.DB.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%-20s %-20s\n", "type", "total")

	for rows.Next() {
		var types string
		var total float32
		err := rows.Scan(&types, &total)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%-20s %-20v\n", types, total)

	}
}
