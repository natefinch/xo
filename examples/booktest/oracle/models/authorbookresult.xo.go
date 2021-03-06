// Package models contains the types for schema 'booktest'.
package models

// GENERATED BY XOXO. DO NOT EDIT.

// AuthorBookResult is the result of a search.
type AuthorBookResult struct {
	AuthorID   float64 // author_id
	AuthorName string  // author_name
	BookID     float64 // book_id
	BookIsbn   string  // book_isbn
	BookTitle  string  // book_title
	BookTags   string  // book_tags
}

// AuthorBookResultsByTags runs a custom query, returning results as AuthorBookResult.
func AuthorBookResultsByTags(db XODB, tags string) ([]*AuthorBookResult, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`a.author_id AS author_id, ` +
		`a.name AS author_name, ` +
		`b.book_id AS book_id, ` +
		`b.isbn AS book_isbn, ` +
		`b.title AS book_title, ` +
		`b.tags AS book_tags ` +
		`FROM books b ` +
		`JOIN authors a ON a.author_id = b.author_id ` +
		`WHERE b.tags LIKE '%' || :1 || '%'`

	// run query
	XOLog(sqlstr, tags)
	q, err := db.Query(sqlstr, tags)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthorBookResult{}
	for q.Next() {
		abr := AuthorBookResult{}

		// scan
		err = q.Scan(&abr.AuthorID, &abr.AuthorName, &abr.BookID, &abr.BookIsbn, &abr.BookTitle, &abr.BookTags)
		if err != nil {
			return nil, err
		}

		res = append(res, &abr)
	}

	return res, nil
}
