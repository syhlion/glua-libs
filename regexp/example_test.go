//origin code: https://github.com/vadv/gopher-lua-libs/tree/master/regexp
package regexp

import (
	"log"

	inspect "github.com/vadv/gopher-lua-libs/inspect"
	lua "github.com/yuin/gopher-lua"
)

func ExampleMatch() {
	state := lua.NewState()
	Preload(state)
	source := `
    local regexp = require("regexp")
    local reg, err = regexp.compile("hello")
    if err then error(err) end
    local result = reg:match("string: 'hello world'")
    print(result)
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
}

func ExampleFindAllStringSubmatch() {
	state := lua.NewState()
	Preload(state)
	inspect.Preload(state)
	source := `
    local regexp = require("regexp")
    local inspect = require("inspect")
    local reg, err = regexp.compile("string: '(.*)\\s+(.*)'$")
    if err then error(err) end
    local result = reg:find_all_string_submatch("string: 'hello world'")
    print(inspect(result, {newline="", indent=""}))
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
}
