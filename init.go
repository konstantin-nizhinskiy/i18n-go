package i18n

import (
	"github.com/jackc/pgx"
	"os"
	"strconv"
)

var typeStorage string
var cookieName string
var langDefault string
var sqlTable string
var dbPool *pgx.ConnPool
func init(){
	sqlTable=getOS("I18N_SQL_TABLE","error_translate")
	typeStorage=getOS("I18N_TYPE_STORAGE","sql_pg")
	cookieName=getOS("I18N_COOKIE_NAME","cc_lang")
	langDefault=getOS("I18N_DEFAULT","en")
}
func Connect(host string,port string, user string, password string,database string) {
	var err error
	port_int, err := strconv.ParseUint(port, 10, 16)
	if(err!=nil){
		panic(err)
	}
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     host,
			Port:     uint16(port_int),
			User:     user,
			Password: password,
			Database: database,
		},
		MaxConnections: 5,
	}
	dbPool, err = pgx.NewConnPool(connPoolConfig)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
func getOS(keyos string,def string) string {
	value,err:=os.LookupEnv(keyos)
	if(err){
		return value
	}else{
		return def
	}
}
func GetPool() *pgx.ConnPool {
	return dbPool
}
