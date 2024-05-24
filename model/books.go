package model

import (
	"books-store/utils"
	"encoding/base64"
	"fmt"

	"strings"

	"github.com/lib/pq"
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
		var picData []byte
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, &picData)
		base64Data := base64.StdEncoding.EncodeToString(picData)
		if base64Data == "" {
			book.ImgPath = ""
		} else {
			book.ImgPath = book.Id
		}
		books = append(books, book)
	}
	return books, nil
}

func AddBook(book *Book) (*Book, error) {
	fileBytes, err := base64.StdEncoding.DecodeString(book.ImgPath)
	if err != nil {
		return nil, err
	}

	sqlStr := "insert into books(title,author,pyear,price,sales,stock,imgpath) values($1,$2,$3,$4,$5,$6,$7)"
	_, err = utils.Db.Exec(sqlStr, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, fileBytes)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func EditBook(book *Book) (*Book, error) {
	// fmt.Println(book)
	fileBytes, err := base64.StdEncoding.DecodeString(book.ImgPath)
	if err != nil {
		return nil, err
	}

	sqlStr := "update books set title=$2,author=$3,pyear=$4,price=$5,sales=$6,stock=$7,imgpath=$8 where id = $1"
	_, err = utils.Db.Exec(sqlStr, &book.Id, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, fileBytes)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id string) error {
	fmt.Println(id)
	idSlice := strings.Split(id, ",")
	sqlStr := "delete from books where id = any($1)"
	_, err := utils.Db.Exec(sqlStr, pq.Array(idSlice))
	if err != nil {
		return err
	}
	return nil
}

func GetBookById(bookId string) *Book {
	sqlStr := "select id,title,author,pyear,price,sales,stock,imgpath from books where id = $1"
	row := utils.Db.QueryRow(sqlStr, bookId)
	book := &Book{}
	var picData []byte
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Pyear, &book.Price, &book.Sales, &book.Stock, &picData)
	if err != nil {
		return nil
	}
	base64Data := base64.StdEncoding.EncodeToString(picData)
	if base64Data == "" {
		book.ImgPath = ""
	} else {
		book.ImgPath = book.Id
	}
	return book
}

func GetBookImgById(bookId string) []byte {
	sqlStr := "select imgpath from books where id = $1"
	row := utils.Db.QueryRow(sqlStr, bookId)
	var picData []byte
	row.Scan(&picData)
	if picData == nil {
		return nil
	}
	return picData
}
