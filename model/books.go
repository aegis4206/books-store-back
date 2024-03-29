package model

import (
	"books-store/utils"
	"fmt"
)

type Book struct {
	Id      string
	Title   string
	Author  string
	Pyear   string
	Price   string
	Sales   string
	Stock   string
	ImgPath string
}

// 獲取所有書籍
func GetBooks() ([]*Book, error) {
	sqlStr := "select id,title,author,pyear,price,sales,stock,imgpath from books order by id desc"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*Book
	for rows.Next() {
		var book *Book = &Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func AddBook(book *Book) (*Book, error) {
	fmt.Println(book)
	sqlStr := "insert into books(title,author,pyear,price,sales,stock,imgpath) values($1,$2,$3,$4,$5,$6,$7)"
	_, err := utils.Db.Exec(sqlStr, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, nil
}
