package controller

import (
	"net/http"

	// modelパッケージをimportしてmというエイリアスをつける
	m "api-server/model"

	"github.com/labstack/echo"
	"strconv"
)

// 以下、handler関数の定義

// アカウント情報表示用のhandler
func GetAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))
	user := m.GetUser(uid)

	return c.JSON(http.StatusOK, user)

}

// アカウント作成用のhandler
func SignUp(c echo.Context) error {

	name := c.FormValue("name")
	password := c.FormValue("password")

	m.SignUp(name, password)

	return c.String(http.StatusOK, "user created")

}

// アカウント削除用のhandler
func DeleteAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))

	m.DeleteUser(uid)

	return c.String(http.StatusOK, "user deleted")

}

// アカウント編集用のhandler
func UpdateAccount(c echo.Context) error {

	uid, _ := strconv.Atoi(c.Param("uid"))
	name := c.FormValue("name")
	password := c.FormValue("password")
	discription := c.FormValue("discription")

	m.UpdateUser(uid, name, password, discription)

	return c.String(http.StatusOK, "user updated")

}
