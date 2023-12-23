package handler

import (
	"context"
	"fmt"
)

func (h *Handler) Whistlist(userEmail string) {
	user, err := h.GetUserByEmail(userEmail)
	if err != nil {
		panic(err)
	}

	query := `
		SELECT games.gameId, games.gameName, orders.statusPayment
		FROM games
		JOIN orders ON orders.gameId = games.gameId
		WHERE orders.UserId = ?
			AND orders.statusPayment = 'pending'
	`

	ctx := context.Background()

	rows, err := h.DB.QueryContext(ctx, query, user.UserID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var gameID, gameName, statusPayment string
		err := rows.Scan(&gameID, &gameName, &statusPayment)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s %s %s\n", gameID, gameName, statusPayment)

		GetStatusPayment()

		if GetStatusPayment() == "done" {

			deleteQuery := `
				DELETE FROM whishlist
				WHERE UserId = ? AND gameId = ?
			`
			_, err := h.DB.ExecContext(ctx, deleteQuery, user.UserID, gameID)
			if err != nil {
				panic(err)
			}

			updateQuery := `
				INSERT INTO chart (UserId, gameId)
				VALUES (?, ?)
			`
			_, err = h.DB.ExecContext(ctx, updateQuery, user.UserID, gameID)
			if err != nil {
				panic(err)
			}
		}
	}

}
