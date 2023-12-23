package handler

import (
	"context"
	"fmt"
	"pairproject/entity"
	"time"
)

func GetStatusPayment() string {

	fmt.Println("Bayar sekarang Y/N?")
	var statusPayment string
	fmt.Scanln(&statusPayment)
	if statusPayment == "y" || statusPayment == "Y" {

		return "done"
	}
	return "pending"
}

func (h *Handler) GetGameByID(gameID int) (entity.Games, error) {
	game := entity.Games{}
	ctx := context.Background()
	query := `
		SELECT GameId, GameName, AgeCategory, Genre, Details, Price
		FROM Games WHERE GameId = ?
	`

	rows, err := h.DB.QueryContext(ctx, query, gameID)
	if err != nil {
		return game, err
	}

	if rows.Next() {
		err := rows.Scan(&game.GameId, &game.GameName, &game.AgeCategory, &game.Genre, &game.Details, &game.Price)
		if err != nil {
			return game, err
		}
	}

	return game, nil
}

func (h *Handler) GetUserByEmail(email string) (entity.Users, error) {
	user := entity.Users{}
	ctx := context.Background()
	query := `
		SELECT userId, name, age, email, password
		FROM Users WHERE Email = ?
	`

	rows, err := h.DB.QueryContext(ctx, query, email)
	if err != nil {
		return user, err
	}

	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Email, &user.Password)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (h *Handler) BuyGames(userEmail string) {
	totalPrice := 0
	time := time.Now()
	h.ShowGames()
	user, err := h.GetUserByEmail(userEmail)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mau beli game yang mana? (Masukkan nomor game)")

	var inputGames int
	fmt.Scanln(&inputGames)

	game, err := h.GetGameByID(inputGames)
	if err != nil {
		panic(err)
	}

	if user.Age < game.AgeCategory {
		fmt.Println("Anda belum cukup umur.")
		h.BuyGames(userEmail)
	}

	totalPrice += game.Price
	fmt.Printf("total pembayaran : $%d", totalPrice)

	if GetStatusPayment() == "done" {
		query := `
			INSERT INTO Orders (UserId, GameId, orderDate, StatusPayment)
			VALUES (?, ?, ?, ?)
		`

		ctx := context.Background()
		_, err = h.DB.ExecContext(ctx, query, user.UserID, inputGames, time, GetStatusPayment())
		if err != nil {
			panic(err)
		}

		fmt.Println("Pembelian berhasil dilakukan!")
	} else if GetStatusPayment() == "pending" {
		query := `
			INSERT INTO Wishtlist (UserId, GameId)
			VALUES (?, ?)
		`

		ctx := context.Background()
		_, err = h.DB.ExecContext(ctx, query, user.UserID, inputGames)
		if err != nil {
			panic(err)
		}

		fmt.Println("Game telah ditambahkan ke dalam Wishlist!")
	} else {
		fmt.Println("Pembelian dibatalkan karena status pembayaran tidak valid.")
	}
	query := `
	INSERT INTO Orders (UserId, GameId, orderDate, StatusPayment)
	VALUES (?, ?, ?, ?)
	`

	ctx := context.Background()
	_, err = h.DB.ExecContext(ctx, query, user.UserID, inputGames, time, GetStatusPayment())
	if err != nil {
		panic(err)
	}

}
