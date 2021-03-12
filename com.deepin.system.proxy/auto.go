package proxy

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type App struct {
	app // interface com.deepin.system.proxy.App
	proxy.Object
}

func NewApp(conn *dbus.Conn) *App {
	obj := new(App)
	obj.Object.Init_(conn, "com.deepin.system.proxy", "/com/deepin/system/proxy/App")
	return obj
}

type app struct{}

func (v *app) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*app) GetInterfaceName_() string {
	return "com.deepin.system.proxy.App"
}

// method AddProxy

func (v *app) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxy", flags, ch, proto, name, proxy)
}

func (v *app) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	return (<-v.GoAddProxy(flags, make(chan *dbus.Call, 1), proto, name, proxy).Done).Err
}

// method AddProxyApps

func (v *app) GoAddProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxyApps", flags, ch, app)
}

func (v *app) AddProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoAddProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method ClearProxy

func (v *app) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearProxy", flags, ch)
}

func (v *app) ClearProxy(flags dbus.Flags) error {
	return (<-v.GoClearProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method DelProxyApps

func (v *app) GoDelProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DelProxyApps", flags, ch, app)
}

func (v *app) DelProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoDelProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method GetProxy

func (v *app) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch)
}

func (*app) StoreGetProxy(call *dbus.Call) (proxy string, err error) {
	err = call.Store(&proxy)
	return
}

func (v *app) GetProxy(flags dbus.Flags) (proxy string, err error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetProxies

func (v *app) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxies", flags, ch, proxies)
}

func (v *app) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	return (<-v.GoSetProxies(flags, make(chan *dbus.Call, 1), proxies).Done).Err
}

// method StartProxy

func (v *app) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartProxy", flags, ch, proto, name, udp)
}

func (v *app) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	return (<-v.GoStartProxy(flags, make(chan *dbus.Call, 1), proto, name, udp).Done).Err
}

// method StopProxy

func (v *app) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopProxy", flags, ch)
}

func (v *app) StopProxy(flags dbus.Flags) error {
	return (<-v.GoStopProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Proxy

func (v *app) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Proxy", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Proxy",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var proxy []interface{}
		err := dbus.Store(sig.Body, &proxy)
		if err == nil {
			cb(proxy)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

type Global struct {
	global // interface com.deepin.system.proxy.Global
	proxy.Object
}

func NewGlobal(conn *dbus.Conn) *Global {
	obj := new(Global)
	obj.Object.Init_(conn, "com.deepin.system.proxy", "/com/deepin/system/proxy/Global")
	return obj
}

type global struct{}

func (v *global) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*global) GetInterfaceName_() string {
	return "com.deepin.system.proxy.Global"
}

// method AddProxy

func (v *global) GoAddProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, proxy []uint8) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddProxy", flags, ch, proto, name, proxy)
}

func (v *global) AddProxy(flags dbus.Flags, proto string, name string, proxy []uint8) error {
	return (<-v.GoAddProxy(flags, make(chan *dbus.Call, 1), proto, name, proxy).Done).Err
}

// method ClearProxy

func (v *global) GoClearProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearProxy", flags, ch)
}

func (v *global) ClearProxy(flags dbus.Flags) error {
	return (<-v.GoClearProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetProxy

func (v *global) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetProxy", flags, ch)
}

func (*global) StoreGetProxy(call *dbus.Call) (proxy string, err error) {
	err = call.Store(&proxy)
	return
}

func (v *global) GetProxy(flags dbus.Flags) (proxy string, err error) {
	return v.StoreGetProxy(
		<-v.GoGetProxy(flags, make(chan *dbus.Call, 1)).Done)
}

// method IgnoreProxyApps

func (v *global) GoIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".IgnoreProxyApps", flags, ch, app)
}

func (v *global) IgnoreProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoIgnoreProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// method SetProxies

func (v *global) GoSetProxies(flags dbus.Flags, ch chan *dbus.Call, proxies []interface{}) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetProxies", flags, ch, proxies)
}

func (v *global) SetProxies(flags dbus.Flags, proxies []interface{}) error {
	return (<-v.GoSetProxies(flags, make(chan *dbus.Call, 1), proxies).Done).Err
}

// method StartProxy

func (v *global) GoStartProxy(flags dbus.Flags, ch chan *dbus.Call, proto string, name string, udp bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartProxy", flags, ch, proto, name, udp)
}

func (v *global) StartProxy(flags dbus.Flags, proto string, name string, udp bool) error {
	return (<-v.GoStartProxy(flags, make(chan *dbus.Call, 1), proto, name, udp).Done).Err
}

// method StopProxy

func (v *global) GoStopProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StopProxy", flags, ch)
}

func (v *global) StopProxy(flags dbus.Flags) error {
	return (<-v.GoStopProxy(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method UnIgnoreProxyApps

func (v *global) GoUnIgnoreProxyApps(flags dbus.Flags, ch chan *dbus.Call, app []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnIgnoreProxyApps", flags, ch, app)
}

func (v *global) UnIgnoreProxyApps(flags dbus.Flags, app []string) error {
	return (<-v.GoUnIgnoreProxyApps(flags, make(chan *dbus.Call, 1), app).Done).Err
}

// signal Proxy

func (v *global) ConnectProxy(cb func(proxy []interface{})) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Proxy", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Proxy",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var proxy []interface{}
		err := dbus.Store(sig.Body, &proxy)
		if err == nil {
			cb(proxy)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property IgnoreApp as

func (v *global) IgnoreApp() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "IgnoreApp",
	}
}
