package handler

import (
	"context"
	"fmt"
	"log"
)

func (h *Handler) ListBook() {
	query := `select book_title
	from books
	join author on author.id = books.author_id
	where author.name = 'Jane smith';`

	ctx := context.Background()
	rows, err := h.DB.QueryContext(ctx, query)

	fmt.Println("book by Jane smith :")

	for rows.Next() {
		var booktitle string
		err = rows.Scan(&booktitle)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", booktitle)
	}

}
