Request Example

Create book

mutation {
createBook(name:"First", author: "Gogol") {
name
}
}

query {
bookByName(name: "First") {
id
name
}
}

query {
book(id: 1) {
id
name
}
}


mutation {
createAuthor(name:"Gogol", book: 1) {
name
}
}

query {
author_list(limit: 5) {
name
}
}


Table structure
SELECT
table_name,
column_name,
data_type
FROM
information_schema.columns
WHERE
table_name = 'books';
books      | id          | integer
books      | name        | text
books      | author      | text


SELECT
table_name,
column_name,
data_type
FROM
information_schema.columns
WHERE
table_name = 'authors';
authors    | id          | integer
authors    | name        | text
authors    | books       | integer

