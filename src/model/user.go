package model

// usersテーブルの操作を定義

// nameとpasswordを引数にとって新規レコードを追加する関数
func SignUp(name, password string) {

	db := InitDB()
	user := User{
		Name:     name,
		Password: password,
	}

	db.Create(&user)

	defer db.Close()

}

// user_idがuidに一致するレコードを返す関数
func GetUser(uid int) User {

	db := InitDB()
	user := User{}

	db.First(&user, uid)

	defer db.Close()
	
	return user

}

// uidがuser_idに合致するレコードの情報を更新する関数
func UpdateUser(uid int, name, password, discription string) {

	db := InitDB()
	user := User{}

	db.Model(&user).Updates(map[string]interface{}{"user_name": name, "user_password": password, "user_discription": discription})

	defer db.Close()

}

// uidがuser_idに合致するレコードを削除する関数
func DeleteUser(uid int) {

	db := InitDB()
	user := User{}

	db.Where("id = ?", uid).Delete(&user)

	defer db.Close()

}

