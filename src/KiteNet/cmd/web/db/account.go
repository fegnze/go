package db

import (
	"KiteNet/log"
	"database/sql"
)

type AccountData struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	OpenID   string `json:"openid"`
	Mail     string `json:"mail"`
}

//GetUserWithOpenID 根据OpenID查询user数据
func GetUserWithOpenID(openid string) *sql.Row {
	query := "SELECT * FROM user_account WHERE openid = ?"
	return DB.QueryRow(query, openid)
}

//插入一条account记录 InsertUserAccount
func InsertUserAccount(account string,password string,mail string,openid string) error {
	stmt,err := DB.Prepare("insert into user_account(account, password, mail, openid) values (?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	result,err := stmt.Exec(account,password,mail,openid)
	if err != nil{
		return err
	}

	//此处可以获取插入的id
	id,err := result.LastInsertId()
	if err != nil{
		return err
	}
	//可以获取影响的行数
	affect,err := result.RowsAffected()
	if err != nil{
		return err
	}
	glog.Info("InsertUserAccount:",id,affect)

	return nil
}

//UpdateUserAccount 更新account数据
func UpdateUserAccount(account string,password string,mail string,openid string) error{
	stmt,err := DB.Prepare("update user_account set account=?,password=?,mail=? where openid=?")
	defer stmt.Close()
	if err != nil {
		return err
	}

	result,err := stmt.Exec(account,password,mail,openid)
	if err != nil {
		return err
	}

	//更新条目
	affect,err := result.RowsAffected()
	if err != nil {
		return err
	}
	glog.Info("UpdataUserAccount:" , affect)

	return nil
}