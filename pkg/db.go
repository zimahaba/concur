package pkg

import (
	"database/sql"
	"fmt"
	"strings"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./concur.db")
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	row := DB.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='conversion'")
	var table string
	row.Scan(&table)

	if table == "" {
		DB.Exec("CREATE TABLE conversion (id TEXT PRIMARY KEY, value DECIMAL(18,8))")
	}

	DB.Exec("CREATE TABLE config (api TEXT PRIMARY KEY, host TEXT, active INTEGER)")
}

func Upsert(conversion Conversion) {
	for k, v := range conversion.Data {
		id := conversion.BaseCurrency + "-" + k
		_, err := DB.Exec("INSERT INTO conversion (id, value) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET value = $2;", id, v.Value)
		if err != nil {
			fmt.Println("unable to insert conversion")
		}
	}
}

func SetCachedCurrencies(conversion *Conversion, baseCurrency string, currencies []string) {
	ids := []interface{}{}
	params := make([]string, len(currencies))
	for i := 0; i < len(currencies); i++ {
		ids = append(ids, baseCurrency+"-"+currencies[i])
		params[i] = "?"
	}

	query := "SELECT * FROM conversion WHERE id IN (" + strings.Join(params, ",") + ")"
	rows, err := DB.Query(query, ids...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var currency Currency
		if err := rows.Scan(&currency.Code, &currency.Value); err != nil {
			panic(err)
		}
		key, _ := strings.CutPrefix(currency.Code, baseCurrency+"-")
		conversion.Data[key] = currency
	}
}
