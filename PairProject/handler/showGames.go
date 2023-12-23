package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"pairproject/entity"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) ShowGames() {
	var games entity.Games

	query := `select * from Games;`

	ctx := context.Background()
	rows, err := h.DB.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&games.GameId, &games.GameName, &games.AgeCategory, &games.Genre, &games.Details, &games.Price)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("No: %d, Name: %s, Age category: %d, Genre: %s, Details: %s, Price: $%d\n", games.GameId, games.GameName, games.AgeCategory, games.Genre, games.Details, games.Price)

	}

}
