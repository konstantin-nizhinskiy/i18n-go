i18n-go
=======

i18n-go is a Go package that helps you translate Go programs into multiple languages.

### Download
```sh
   $ go get github.com/konstantin-nizhinskiy/i18n-go
```

### Config environment

```bash

export I18N_SQL_TABLE="error_translate";
export I18N_TYPE_STORAGE="sql_pg";
export I18N_COOKIE_NAME="cc_lang";
export I18N_DEFAULT="en";
export I18N_SQL_HOST="localhost";
export I18N_SQL_USER="postgres";
export I18N_SQL_PASSWORD="qwerty";
export I18N_SQL_DATABASE="cc_data";

```
 * I18N_SQL_TABLE - Name table
 * I18N_TYPE_STORAGE - Type storage translate(now support only sql_pg)
   * sql_pg - PostgresSql
 * I18N_COOKIE_NAME - Cookie name if you use function GetGin
 * I18N_SQL_HOST - Host name sql server
 * I18N_SQL_USER - User sql server
 * I18N_SQL_PASSWORD - Password sql server
 * I18N_SQL_DATABASE - Database sql server
### Example

```go

package emailBack

import (
	"github.com/konstantin-nizhinskiy/i18n-go"
)

func Controller(c *gin.Context) {
    //...

    i18n.GetGin("Not.found._key_", c, map[string]interface{}{"key":"myKey"}) // en: "Not.found._key_"  : "Not found {{ key }}"
    // return Not found myKey
    
    i18n.GetGin("Not.found._key_", c) // en: "Not.found._key_"  : "Not found {{ key }}"
    // return Not found 
    i18n.Get("Not.found._key_", "en",map[string]interface{}{"key":"myKey2"}) // en: "Not.found._key_"  : "Not found {{ key }}"
    // return Not found myKey2
    i18n.Get("Not.found._key_", "en") // en: "Not.found._key_"  : "Not found {{ key }}"
    // return Not found 

    //...
}

```
### TODO
Add support XML,JSON,YML storage
Add function native net/http
Add MySql support