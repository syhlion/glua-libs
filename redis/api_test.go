package redis

import (
	"fmt"
	"testing"

	"time"

	lua "github.com/yuin/gopher-lua"
)

func TestApi(t *testing.T) {
	state := lua.NewState()

	Preload(state)

	fmt.Println("prepare")
	time.Sleep(5 * time.Second)
	for i := 0; i <= 100000; i++ {
		fmt.Println(i)
		if err := state.DoFile("./test/test_api.lua"); err != nil {
			t.Fatalf("execute test: %s\n", err.Error())
		}
	}
	select {}

}
