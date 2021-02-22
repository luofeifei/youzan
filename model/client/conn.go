package client

import (
	"base/model/imp/serverAdmin"
	"base/model/imp/serverBase"
	"base/model/imp/serverCo"
	"base/model/imp/serverPlugin"
	"base/model/imp/serverShop"
	"base/model/imp/serverUser"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

var Conn *tars.Communicator
var Obj string

func init() {
	Conn = tars.NewCommunicator()
}

func ConnServer(obj string, ip, port string, comm *tars.Communicator) string {
	obj = fmt.Sprintf("%v@tcp -h %v -p %v -t 60000", obj, ip, port)
	return obj
}

func ServerBaseSys() *serverBase.Sys {
	Obj = ConnServer("mall.serverBase.sysObj", "192.168.2.21", "10005", Conn)
	App := new(serverBase.Sys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerPlugin() *serverPlugin.Sys {
	Obj = ConnServer("mall.serverPlugin.sysObj", "192.168.3.110", "10006", Conn)
	App := new(serverPlugin.Sys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerAdminSys() *serverAdmin.AdminSys {
	Obj = ConnServer("mall.serverAdmin.sysObj", "192.168.3.110", "10000", Conn)
	App := new(serverAdmin.AdminSys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerUserSys() *serverUser.UserSys {
	Obj = ConnServer("mall.serverUser.sysObj", "192.168.3.110", "10002", Conn)
	App := new(serverUser.UserSys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerCoSys() *serverCo.Sys {
	Obj = ConnServer("mall.serverCo.sysObj", "192.168.3.110", "10003", Conn)
	App := new(serverCo.Sys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerCoUser() *serverCo.User {
	Obj = ConnServer("mall.serverCo.userObj", "192.168.3.110", "10004", Conn)
	App := new(serverCo.User)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}

func ServerShopSys() *serverShop.ShopSys {
	// dev
	//Obj = ConnServer("mall.serverShop.sysObj", "192.168.3.87", "10003", Conn)
	//Obj = ConnServer("mall.serverShop.sysObj", "192.168.3.20", "10003", Conn)
	Obj = ConnServer("mall.serverShop.sysObj", "192.168.31.114", "10004", Conn)
	//Obj = ConnServer("mall.serverShop.sysObj", "192.168.2.21","10003", Conn)
	App := new(serverShop.ShopSys)
	Conn.StringToProxy(Obj, App)
	App.TarsSetTimeout(3000)
	return App
}
