package window

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xwindow"
	lua "github.com/yuin/gopher-lua"
	"gomint.io/e2e/test/logger"
)

func GetBounds(L *lua.LState) int {
	windowId := uint32(L.ToInt(1))

	EnsureConnection()

	geom, err := xwindow.RawGeometry(xgbUtil, xproto.Drawable(windowId))
	if err != nil {
		logger.Error(err)
	}
	cookie := xproto.TranslateCoordinates(xgbUtil.Conn(), xproto.Window(windowId), xgbUtil.RootWin(), 0, 0)
	reply, err := cookie.Reply()
	if err != nil {
		logger.Error(err)
	}

	L.Push(lua.LNumber(reply.DstX))
	L.Push(lua.LNumber(reply.DstY))
	L.Push(lua.LNumber(geom.Width()))
	L.Push(lua.LNumber(geom.Height()))
	return 4
}
