package i18n

import (
	"github.com/jackc/pgx"
	consulapi "github.com/hashicorp/consul/api"
	"os"
	"strconv"
	"encoding/json"
	"errors"
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
type dbConfig struct{
	Id string `json:"id"`
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	DataBase string `json:"database"`
	Shem string `json:"shem"`
	Password string `json:"password"`
}
func isPanic(err error){
	if(err!=nil){
		panic(err)
	}
}
func ConnectConsul(Consul *consulapi.Client, path string) {
	var config dbConfig
	kv := Consul.KV()
	pair, _, err := kv.Get(path, nil)
	isPanic(err)
	err = json.Unmarshal(pair.Value, &config)
	if (pair == nil) {
		panic(errors.New("Not foud db app"))
	}
	Connect(config.Host,config.Port,config.User,config.Password,config.DataBase)

}
func Connect(host string,port string, user string, password string,database string) {
	var err error
	port_int, err := strconv.ParseUint(port, 10, 16)
	isPanic(err)
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
	isPanic(err)
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
