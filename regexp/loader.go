//origin code: https://github.com/vadv/gopher-lua-libs/tree/master/regexp
package regexp

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("regexp", Loader)
}

func Loader(L *lua.LState) int {

	regexp_ud := L.NewTypeMetatable(`regexp_ud`)
	L.SetGlobal(`regexp_ud`, regexp_ud)
	L.SetField(regexp_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"match":                    Match,
		"find_all_string_submatch": FindAllStringSubmatch,
	}))

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"compile":                  Compile,
	"match":                    SimpleMatch,
	"find_all_string_submatch": SimpleFindAllStringSubmatch,
}
