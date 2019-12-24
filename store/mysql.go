package store

import (
	"fmt"
	"robot_chat/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DBconn *sqlx.DB
)

func NewMysql(user, pass, addr, database string) {
	var err error
	fmt.Println("dbcon :", user, pass, addr, database, " end")

	DBconn, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&multiStatements=true", user, pass, addr, database))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	DBconn.SetMaxIdleConns(100)
	DBconn.SetMaxOpenConns(100)
	fmt.Println("[InitMysql] MySQL connected.")
}

func CloseMysql() {
	DBconn.Close()
}

//查询是否是使用管家
func SelectRobotStatusMysql(uid string) (UserInfoData, error) {
	sqlStmt := fmt.Sprintf("Select robot_status,greet_status from chat_user where uid = '%v'", uid)
	userInfoArr := make([]UserInfoData, 0)
	userInfo := UserInfoData{}
	err := DBconn.Select(&userInfoArr, sqlStmt)

	if err != nil {
		logger.Logger.Error("[SelectRobotStatusMysql] sqlStmt:%v err:%v", sqlStmt, err)
		return userInfo, err
	}
	fmt.Printf("sqlStmt = %v\n", sqlStmt)
	fmt.Printf("userInfoArr = %v\n", userInfoArr)
	if len(userInfoArr) > 0 {
		userInfo = userInfoArr[0]
	}
	return userInfo, err

}

//更新管家状态
func UpdateRobotStatusMysql(uid string, status int) error {
	sqlStmt := fmt.Sprintf("Update chat_user set robot_status = %v where uid = '%v'", status, uid)
	_, err := DBconn.Exec(sqlStmt)
	if err != nil {
		logger.Logger.Error("[UpdateRobotStatusMysql] sqlStmt:%v err:%v", sqlStmt, err)
	}
	return err
}

//查询关键词(按字数长度排序)
func SelectAllKeysMysql() ([]*string, error) {
	sqlStmt := fmt.Sprintf("Select keyid from chat_keyworld")
	keyArr := make([]*string, 0)
	err := DBconn.Select(&keyArr, sqlStmt)

	if err != nil {
		logger.Logger.Error("[SelectAllKeysMysql] sqlStmt:%v err:%v", sqlStmt, err)
		return nil, err
	}
	return keyArr, err
}

func SelectAllKeyWorldsMysql() ([]*KeyWordData, error) {
	sqlStmt := fmt.Sprintf("Select keyid,back_world from chat_keyworld")
	keyInfoArr := make([]*KeyWordData, 0)
	err := DBconn.Select(&keyInfoArr, sqlStmt)

	if err != nil {
		logger.Logger.Error("[SelectAllKeysMysql] sqlStmt:%v err:%v", sqlStmt, err)
		return nil, err
	}
	return keyInfoArr, err
}

// 查询关键词对应的话术(用上面的)

//查询机器人回复 时 用来判断是否无效 的 错误的关键字
func SelectAllBadWorldMysql() ([]*string, error) {
	sqlStmt := fmt.Sprintf("Select bad_world from chat_badworld")
	keyArr := make([]*string, 0)
	err := DBconn.Select(&keyArr, sqlStmt)

	if err != nil {
		logger.Logger.Error("[SelectAllBadWorldMysql] sqlStmt:%v err:%v", sqlStmt, err)
		return nil, err
	}
	return keyArr, err
}

func SelectAllExchangeWorldMysql() ([]*ExchangeWordData, error) {
	sqlStmt := fmt.Sprintf("Select keyid,exchange_world from chat_exchangeworld")
	keyInfoArr := make([]*ExchangeWordData, 0)
	err := DBconn.Select(&keyInfoArr, sqlStmt)

	if err != nil {
		logger.Logger.Error("[SelectAllExchangeWorldMysql] sqlStmt:%v err:%v", sqlStmt, err)
		return nil, err
	}
	return keyInfoArr, err
}
