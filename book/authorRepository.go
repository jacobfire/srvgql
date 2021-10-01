package book

import (
	"database/sql"
	"fmt"
	"graphql/infrastructure"
	"log"
)

func FetchAuthorByName(name string) (result interface{}) {
	var author AuthorEntity

	err := infrastructure.Postgres.QueryRow(
		"SELECT id, name, books FROM authors WHERE name = $1",
		name,
	).Scan(&author.Id, &author.Name, &author.Book)
	if err != nil {
		fmt.Println(err)
	}

	return author
}

func FetchAuthorById(id int) (result interface{}) {
	var author AuthorEntity

	err := infrastructure.Postgres.QueryRow(
		"SELECT id, name, books FROM authors WHERE id = $1",
		id,
	).Scan(&author.Id, &author.Name, &author.Book)
	if err != nil {
		fmt.Println(err)
	}

	return author
}

func FetchAuthorList(limit int) (result interface{}) {
	rows, err := infrastructure.Postgres.Query(
		"SELECT id, name, books FROM authors LIMIT $1",
		limit,
	)

	defer rows.Close()
	if err == sql.ErrNoRows {
		fmt.Println(err)
	}

	var authors []AuthorEntity

	author := AuthorEntity{}
	for rows.Next() {
		if err := rows.Scan(&author.Id, &author.Name, &author.Book); err != nil {
			log.Fatal(err)
		}

		authors = append(authors, author)
		log.Printf("id %d has role %s\n", author.Id, author)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return authors
}

func InsertAuthorEntity(author AuthorEntity) error {
	if err := infrastructure.Postgres.QueryRow(
		"INSERT INTO authors (name, books) VALUES($1, $2) RETURNING id",
		author.Name,
		author.Book,
	).Scan(&author.Id); err != nil {
		return err
	}

	return nil
}

func UpdateAuthorEntity(author AuthorEntity) error {
	_, err := infrastructure.Postgres.Exec("UPDATE authors SET name = $1, books = $2 WHERE id = $3",
		author.Name,
		author.Book,
		author.Id,
	)

	if err != nil {
		fmt.Println("update caused error")
		return err
	}

	return nil
}

func DeleteAuthorEntityById(id int) error {
	_, err := infrastructure.Postgres.Exec("DELETE FROM authors WHERE id = $1", id)

	if err != nil {
		fmt.Println("deleting caused error")
		return err
	}

	return nil
}

