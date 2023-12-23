package main

import (
	"connectdb/cli"
	"connectdb/config"
	"connectdb/handler"
)

// func handlePanic() {
// 	if r := recover(); r != nil {
// 		fmt.Println("recovered from panic", r)
// 	}
// }

// func readSQLCommand(filename string) (string, error) {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), nil
// }

// func createTableCommand(db *sql.DB, command string) error {
// 	_, err := db.Exec(command)
// 	return err
// }

func main() {

	db := config.ConnectToDatabase()
	defer db.Close()
	userHandler := handler.NewUserHandler(db)
	cli := cli.NewCli(userHandler)
	cli.Show()

	// defer handlePanic()

	// args := os.Args[1:]

	// if len(args) < 1 {
	// 	panic("no sql file provided")
	// }

	// filename := args[0]

	// sqlCommand, err := readSQLCommand(filename)
	// if err != nil {
	// 	panic(err)
	// }

	// db, err := connectToDatabase()
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// err = createTableCommand(db, sqlCommand)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("sukses migrasi database")
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nyobadb")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("failed connecting")
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("sukses konek")

	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INT AUTO_INCREMENT PRIMARY KEY,name VARCHAR(50))")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("sukses bikin table")
	// }

	// _, err = db.Exec("INSERT INTO test(name) VALUES(?)", "John")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("sukses insert")
	// }

	// rows, err := db.Query("SELECT * FROM test")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	fmt.Println(id, name)
	// }

}
