package i18n

import (
	"github.com/flosch/pongo2"
	"github.com/jackc/pgx/pgtype"

)

func Get(key string, lang string, data ...map[string]interface{}) string {

	var context string

	switch typeStorage {
	case "sql_pg":
		context = getSqlPg(key, lang)
		break
	default:
		panic("Not support type storage"+typeStorage)
	}
	if (len(data) > 0) {
		tpl, err := pongo2.FromString(context)
		if (err != nil) {
			panic(err)
		}
		context, err = tpl.Execute(pongo2.Context(data[0]))
		if (err != nil) {
			panic(err)
		}
	}
	return context
}

func getSqlPg(key string, lang string) string {
	row := GetPool().QueryRow("select txt_translate from " + sqlTable + " where id=$1 and lang=$2", key, lang)
	var txt_translate pgtype.Varchar
	err := row.Scan(&txt_translate)
	if (err!=nil) {
		return key
	}
	return txt_translate.String

}