//origin code: https://github.com/vadv/gopher-lua-libs/tree/master/strings
package strings

import (
	"strings"

	lua "github.com/yuin/gopher-lua"
)

// Split(): lua strings.split(string, sep): port of go string.Split() returns table
func Split(L *lua.LState) int {
	str := L.CheckString(1)
	delim := L.CheckString(2)
	strSlice := strings.Split(str, delim)
	result := L.CreateTable(len(strSlice), 0)
	for _, str := range strSlice {
		result.Append(lua.LString(str))
	}
	L.Push(result)
	return 1
}

// HasPrefix(): lua strings.has_prefix(string, suffix): port of go string.HasPrefix() return bool
func HasPrefix(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.HasPrefix(str1, str2)
	L.Push(lua.LBool(result))
	return 1
}

// HasSuffix(): lua strings.has_suffix(string, prefix): port of go string.HasSuffix() returns bool
func HasSuffix(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.HasSuffix(str1, str2)
	L.Push(lua.LBool(result))
	return 1
}

// Trim(): lua strings.trim(string, cutset) Port of go string.Trim() returns string
func Trim(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.Trim(str1, str2)
	L.Push(lua.LString(result))
	return 1
}

// TrimPrefix(): lua strings.trim_prefix(string, cutset) Port of go string.TrimPrefix() returns string
func TrimPrefix(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.TrimPrefix(str1, str2)
	L.Push(lua.LString(result))
	return 1
}

// TrimSuffix(): lua strings.trim_suffix(string, cutset) Port of go string.TrimSuffix() returns string
func TrimSuffix(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.TrimSuffix(str1, str2)
	L.Push(lua.LString(result))
	return 1
}

// Contains(): lua strings.contains(string, cutset) Port of go string.Contains() returns bool
func Contains(L *lua.LState) int {
	str1 := L.CheckString(1)
	str2 := L.CheckString(2)
	result := strings.Contains(str1, str2)
	L.Push(lua.LBool(result))
	return 1
}
