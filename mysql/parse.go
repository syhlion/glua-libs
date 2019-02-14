//origin code: https://github.com/vadv/gopher-lua-libs/tree/master/db
package db

import (
	"database/sql"
	"log"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func parseRows(sqlRows *sql.Rows, L *lua.LState) (*lua.LTable, *lua.LTable, error) {

	cols, err := sqlRows.Columns()
	if err != nil {
		return nil, nil, err
	}
	columns := L.CreateTable(len(cols), 1)
	for _, col := range cols {
		columns.Append(lua.LString(col))
	}

	luaRows := L.CreateTable(0, len(cols))
	rowCount := 1
	for sqlRows.Next() {
		columns := make([]interface{}, len(cols))
		pointers := make([]interface{}, len(cols))
		for i := range columns {
			pointers[i] = &columns[i]
		}
		err := sqlRows.Scan(pointers...)
		if err != nil {
			return nil, nil, err
		}
		luaRow := L.CreateTable(0, len(cols))
		for i := range cols {
			valueP := pointers[i].(*interface{})
			value := *valueP
			switch converted := value.(type) {
			case bool:
				luaRow.RawSetInt(i+1, lua.LBool(converted))
			case float64:
				luaRow.RawSetInt(i+1, lua.LNumber(converted))
			case int64:
				luaRow.RawSetInt(i+1, lua.LNumber(converted))
			case string:
				luaRow.RawSetInt(i+1, lua.LString(converted))
			case time.Time:
				tt := float64(converted.UTC().UnixNano()) / float64(time.Second)
				luaRow.RawSetInt(i+1, lua.LNumber(tt))
			case []byte:
				luaRow.RawSetInt(i+1, lua.LString(string(converted)))
			case nil:
				luaRow.RawSetInt(i+1, lua.LNil)
			default:
				log.Printf("[ERROR] unknown type (value: `%#v`, converted: `%#v`)\n", value, converted)
				luaRow.RawSetInt(i+1, lua.LNil)
			}
		}
		luaRows.RawSet(lua.LNumber(rowCount), luaRow)
		rowCount++
	}
	return luaRows, columns, nil
}
