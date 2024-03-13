package plugin

import (
	"github.com/cat3306/gnetrpc"
	"github.com/cat3306/gnetrpc/rpclog"
	"github.com/cat3306/gnetrpc/util"
	"github.com/panjf2000/gnet/v2"
)

type ClosePlugin struct {
}

func (c *ClosePlugin) Type() gnetrpc.PluginType {
	return gnetrpc.PluginTypeOnClose
}
func (c *ClosePlugin) Init(args ...interface{}) gnetrpc.Plugin {
	return c
}
func (c *ClosePlugin) OnDo(args ...interface{}) interface{} {
	if len(args) == 0 {
		return false
	}
	conn := args[0].(gnet.Conn)
	rpclog.Warnf("client close id:%s,cause:%v", util.GetConnId(conn), args[1])
	return true
}
