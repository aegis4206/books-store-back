package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("測試userdao中的函數")
	t.Run("驗證用戶名與密碼:", testLogin)
	t.Run("驗證用戶名:", testRegist)
	t.Run("保存用戶:", testSave)

}

func testLogin(t *testing.T) {
	user, _ := CheckEmailAndPassword("white@white.white", "123456")
	fmt.Println("用戶訊息是:", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckEmail("admin")
	fmt.Println("用戶訊息是:", user)
}
func testSave(t *testing.T) {
	SaveUser("white@white.white", "123456")
}
