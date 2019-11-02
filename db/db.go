package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	hostKey   = "MYSQL_HOST"
	portKey   = "MYSQL_PORT"
	dbKey     = "MYSQL_DB"
	passwdKey = "MYSQL_ROOT_PASSWORD"
	user      = "root"
	loc       = "Asia/Shanghai"

	defaultHost     = "39.106.120.138"
	defaultPort     = "3306"
	defaultDB       = "oj_database"
	defaultPassword = "maxinz"

	selectResultCountStr        = `SELECT result, count(*) as _count FROM submission WHERE user_id=? AND contest_id is NULL GROUP BY result;`
	selectProfileSubmitCountStr = "SELECT submission_number, accepted_number FROM user_profile WHERE user_id=?;"
)

var (
	dbConn *sql.DB

	judgeStatus = map[int]string{
		-2: "COMPILE_ERROR",
		-1: "WRONG_ANSWER",
		0:  "ACCEPTE",
		8:  "PARTIALLY_ACCEPTED",
	}
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func init() {

	host := getEnv(hostKey, defaultHost)
	port := getEnv(portKey, defaultPort)
	db := getEnv(dbKey, defaultDB)
	password := getEnv(passwdKey, defaultPassword)

	var err error
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true",
		user, password, host, port, db, url.QueryEscape(loc))

	if dbConn, err = sql.Open("mysql", uri); err != nil {
		panic(err)
	}
}

// GetProfileSubmitCount 获取提交总数和通过总数
func GetProfileSubmitCount(userID string, res map[string]int) (map[string]int, error) {
	row := dbConn.QueryRow(selectProfileSubmitCountStr, userID)
	submitCount, acceptCount := 0, 0
	if err := row.Scan(&submitCount, &acceptCount); err != nil {
		log.Println("error:", err)
		return res, err
	}
	res["SUBMIT_COUNT"] = submitCount
	res["ACCEPT_COUNT"] = acceptCount
	return res, nil
}

// GetUserSubmitResultCount 获取某用户的提交结果统计
func GetUserSubmitResultCount(userID string, res map[string]int) (map[string]int, error) {
	rows, err := dbConn.Query(selectResultCountStr, userID)
	// defer rows.Close()
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}

	resultCount, resultType := 0, 0
	for rows.Next() {
		if err = rows.Scan(&resultType, &resultCount); err != nil {
			log.Println("error:", err)
			return res, err
		}
		if status, ok := judgeStatus[resultType]; ok {
			res[status] = resultCount
		}
	}

	return res, nil
}
