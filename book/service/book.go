package service

import (
	"book/model"
)

// BookService 構造体宣言
type BookService struct{}

/*** DB Insert ***/
func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

/*** DB Select ***/
func (BookService) GetBookList() []model.Book {
	// 取得したDB名称を変数に格納 Find()でDB指定するため
	tests := make([]model.Book, 0)
	// Distinctでカラムを指定 Limitでレスポンスの件数を指定 FindでDB指定で探す
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}

/*** DB UPDATE ***/
func (BookService) UpdateBook(newBook *model.Book) error {
	// Id()でmodel.Book構造体からIDだけを取得 Update()でmodel.Book構造体に入った値に更新
	_, err := DbEngine.Id(newBook.Id).Update(newBook)
	if err != nil {
		return err
	}
	return nil
}

/*** DB DELETE ***/
func (BookService) DeleteBook(id int) error {
	// model.Book構造体をDeleteメソッド引数のために宣言
	book := new(model.Book)
	// Id()でidを指定し、指定した番号を削除  Delete()はDBを指定するため
	_, err := DbEngine.Id(id).Delete(book)
	if err != nil {
		return err
	}
	return nil
}
