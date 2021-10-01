package book

import (
	"database/sql"
	"fmt"
	"graphql/infrastructure"
	"log"
)

func FetchBookByName(name string) (result interface{}) {
	var book BookEntity

	err := infrastructure.Postgres.QueryRow(
		"SELECT id, name, author FROM books WHERE name = $1",
		name,
	).Scan(&book.Id, &book.Name, &book.Author)
	if err != nil {
		fmt.Println(err)
	}

	return book
}

func FetchBookById(id int) (result interface{}) {
	var book BookEntity

	err := infrastructure.Postgres.QueryRow(
		"SELECT id, name, author FROM books WHERE id = $1",
		id,
	).Scan(&book.Id, &book.Name, &book.Author)
	if err != nil {
		fmt.Println(err)
	}

	return book
}

func FetchBookList(limit int) (result interface{}) {
	rows, err := infrastructure.Postgres.Query(
		"SELECT id, name, author FROM books LIMIT $1",
		limit,
	)

	defer rows.Close()
	if err == sql.ErrNoRows {
		fmt.Println(err)
	}

	var books []BookEntity

	book := BookEntity{}
	for rows.Next() {
		if err := rows.Scan(&book.Id, &book.Name, &book.Author); err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
		log.Printf("id %d has role %s\n", book.Id, book)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return books
}

func InsertBookEntity(book BookEntity) error {
	if err := infrastructure.Postgres.QueryRow(
		"INSERT INTO books (name, author) VALUES($1, $2) RETURNING id",
		book.Name,
		book.Author,
	).Scan(&book.Id); err != nil {
		return err
	}

	return nil
}

func UpdateBookEntity(book BookEntity) error {
	_, err := infrastructure.Postgres.Exec("UPDATE books SET name = $1, author = $2 WHERE id = $3",
		book.Name,
		book.Author,
		book.Id,
		)

	if err != nil {
		fmt.Println("update caused error")
		return err
	}

	return nil
}

func DeleteBookEntityById(id int) error {
	_, err := infrastructure.Postgres.Exec("DELETE FROM books WHERE id = $1", id)

	if err != nil {
		fmt.Println("deleting caused error")
		return err
	}

	return nil
}
