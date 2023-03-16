package gomysql

import (
	"log"
	"db"
)

func TableCount(tableName string) (uint64, error) {
	var count uint64
	if db.QueryRow("SELECT COUNT(*) FROM "+tableName).Scan(&count) != nil {
		//switch {
		//case err != nil:
		//		log.Printf(err)
		return 0,err
	}
	return count, nil
}
