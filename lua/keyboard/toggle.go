package keyboard

import (
	"github.com/go-vgo/robotgo"
	lua "github.com/yuin/gopher-lua"
)

func Toggle(L *lua.LState) int {
	key := L.ToString(1)
	arg1 := L.ToString(2)
	robotgo.KeyToggle(key, arg1)
	return 0
}
