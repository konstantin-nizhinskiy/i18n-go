package i18n

import (
	"github.com/jackc/pgx"
	"os"
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
	if(typeStorage=="sql_pg"){
		var err error
		connPoolConfig := pgx.ConnPoolConfig{
			ConnConfig: pgx.ConnConfig{
				Host:     getOS("I18N_SQL_HOST","localhost"),
				User:     getOS("I18N_SQL_USER","postgres"),
				Password: getOS("I18N_SQL_PASSWORD","qwerty"),
				Database: getOS("I18N_SQL_DATABASE","cc_data"),
			},
			MaxConnections: 5,
		}
		dbPool, err = pgx.NewConnPool(connPoolConfig)
		if err != nil {
			panic(err)
			os.Exit(1)
		}
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
