// +build !linux

package window

import (
	lua "github.com/yuin/gopher-lua"
	"github.com/go-vgo/robotgo"
)

func GetBounds(L *lua.LState) int {
	windowId := int32(L.ToInt(1))

	x, y, w, h := robotgo.GetBounds(windowId, int(windowId))

	L.Push(lua.LNumber(x))
	L.Push(lua.LNumber(y))
	L.Push(lua.LNumber(w))
	L.Push(lua.LNumber(h))
	return 4
}