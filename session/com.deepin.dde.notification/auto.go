// Code generated by "./generator ./session/com.deepin.dde.notification"; DO NOT EDIT.

package notification

import "errors"
import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type Notification interface {
	notification // interface com.deepin.dde.Notification
	proxy.Object
}

type objectNotification struct {
	interfaceNotification // interface com.deepin.dde.Notification
	proxy.ImplObject
}

func NewNotification(conn *dbus.Conn) Notification {
	obj := new(objectNotification)
	obj.ImplObject.Init_(conn, "com.deepin.dde.osd", "/com/deepin/dde/Notification")
	return obj
}

type notification interface {
	GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call
	CloseNotification(flags dbus.Flags, arg0 uint32) error
	GoGetCapbilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetCapbilities(flags dbus.Flags) ([]string, error)
	GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetServerInformation(flags dbus.Flags) (string, string, string, string, error)
	GoNotify(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) *dbus.Call
	Notify(flags dbus.Flags, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) (uint32, error)
	GoGetAllRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAllRecords(flags dbus.Flags) (string, error)
	GoGetRecordById(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetRecordById(flags dbus.Flags, arg0 string) (string, error)
	GoGetRecordsFromId(flags dbus.Flags, ch chan *dbus.Call, arg0 int32, arg1 string) *dbus.Call
	GetRecordsFromId(flags dbus.Flags, arg0 int32, arg1 string) (string, error)
	GoRemoveRecord(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	RemoveRecord(flags dbus.Flags, arg0 string) error
	GoClearRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearRecords(flags dbus.Flags) error
	GoGetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	GetAppSetting(flags dbus.Flags, arg0 string) (string, error)
	GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Toggle(flags dbus.Flags) error
	GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Show(flags dbus.Flags) error
	GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Hide(flags dbus.Flags) error
	GoRecordCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	RecordCount(flags dbus.Flags) (uint32, error)
	GoGetAppList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAppList(flags dbus.Flags) ([]string, error)
	GoGetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call
	GetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32) (dbus.Variant, error)
	GoGetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call
	GetSystemInfo(flags dbus.Flags, arg0 uint32) (dbus.Variant, error)
	GoSetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 dbus.Variant) *dbus.Call
	SetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32, arg2 dbus.Variant) error
	GoSetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32, arg1 dbus.Variant) *dbus.Call
	SetSystemInfo(flags dbus.Flags, arg0 uint32, arg1 dbus.Variant) error
	GoSetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	SetAppSetting(flags dbus.Flags, arg0 string) error
	ConnectNotificationClosed(cb func(arg0 uint32, arg1 uint32)) (dbusutil.SignalHandlerId, error)
	ConnectActionInvoked(cb func(arg0 uint32, arg1 string)) (dbusutil.SignalHandlerId, error)
	ConnectRecordAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectAppInfoChanged(cb func(arg0 string, arg1 uint32, arg2 dbus.Variant)) (dbusutil.SignalHandlerId, error)
	ConnectSystemInfoChanged(cb func(arg0 uint32, arg1 dbus.Variant)) (dbusutil.SignalHandlerId, error)
	ConnectAppAddedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectAppRemovedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectAppRemoved(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectAppAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectAppSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	ConnectSystemSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error)
	AllSetting() proxy.PropString
	SystemSetting() proxy.PropString
}

type interfaceNotification struct{}

func (v *interfaceNotification) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceNotification) GetInterfaceName_() string {
	return "com.deepin.dde.Notification"
}

// method CloseNotification

func (v *interfaceNotification) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CloseNotification", flags, ch, arg0)
}

func (v *interfaceNotification) CloseNotification(flags dbus.Flags, arg0 uint32) error {
	return (<-v.GoCloseNotification(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetCapbilities

func (v *interfaceNotification) GoGetCapbilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapbilities", flags, ch)
}

func (*interfaceNotification) StoreGetCapbilities(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceNotification) GetCapbilities(flags dbus.Flags) ([]string, error) {
	return v.StoreGetCapbilities(
		<-v.GoGetCapbilities(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetServerInformation

func (v *interfaceNotification) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetServerInformation", flags, ch)
}

func (*interfaceNotification) StoreGetServerInformation(call *dbus.Call) (arg0 string, arg1 string, arg2 string, arg3 string, err error) {
	err = call.Store(&arg0, &arg1, &arg2, &arg3)
	return
}

func (v *interfaceNotification) GetServerInformation(flags dbus.Flags) (string, string, string, string, error) {
	return v.StoreGetServerInformation(
		<-v.GoGetServerInformation(flags, make(chan *dbus.Call, 1)).Done)
}

// method Notify

func (v *interfaceNotification) GoNotify(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Notify", flags, ch, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (*interfaceNotification) StoreNotify(call *dbus.Call) (arg8 uint32, err error) {
	err = call.Store(&arg8)
	return
}

func (v *interfaceNotification) Notify(flags dbus.Flags, arg0 string, arg1 uint32, arg2 string, arg3 string, arg4 string, arg5 []string, arg6 map[string]dbus.Variant, arg7 int32) (uint32, error) {
	return v.StoreNotify(
		<-v.GoNotify(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7).Done)
}

// method GetAllRecords

func (v *interfaceNotification) GoGetAllRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllRecords", flags, ch)
}

func (*interfaceNotification) StoreGetAllRecords(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceNotification) GetAllRecords(flags dbus.Flags) (string, error) {
	return v.StoreGetAllRecords(
		<-v.GoGetAllRecords(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetRecordById

func (v *interfaceNotification) GoGetRecordById(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetRecordById", flags, ch, arg0)
}

func (*interfaceNotification) StoreGetRecordById(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceNotification) GetRecordById(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetRecordById(
		<-v.GoGetRecordById(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method GetRecordsFromId

func (v *interfaceNotification) GoGetRecordsFromId(flags dbus.Flags, ch chan *dbus.Call, arg0 int32, arg1 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetRecordsFromId", flags, ch, arg0, arg1)
}

func (*interfaceNotification) StoreGetRecordsFromId(call *dbus.Call) (arg2 string, err error) {
	err = call.Store(&arg2)
	return
}

func (v *interfaceNotification) GetRecordsFromId(flags dbus.Flags, arg0 int32, arg1 string) (string, error) {
	return v.StoreGetRecordsFromId(
		<-v.GoGetRecordsFromId(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method RemoveRecord

func (v *interfaceNotification) GoRemoveRecord(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveRecord", flags, ch, arg0)
}

func (v *interfaceNotification) RemoveRecord(flags dbus.Flags, arg0 string) error {
	return (<-v.GoRemoveRecord(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method ClearRecords

func (v *interfaceNotification) GoClearRecords(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearRecords", flags, ch)
}

func (v *interfaceNotification) ClearRecords(flags dbus.Flags) error {
	return (<-v.GoClearRecords(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method getAppSetting

func (v *interfaceNotification) GoGetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".getAppSetting", flags, ch, arg0)
}

func (*interfaceNotification) StoreGetAppSetting(call *dbus.Call) (arg1 string, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceNotification) GetAppSetting(flags dbus.Flags, arg0 string) (string, error) {
	return v.StoreGetAppSetting(
		<-v.GoGetAppSetting(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method Toggle

func (v *interfaceNotification) GoToggle(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Toggle", flags, ch)
}

func (v *interfaceNotification) Toggle(flags dbus.Flags) error {
	return (<-v.GoToggle(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *interfaceNotification) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *interfaceNotification) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Hide

func (v *interfaceNotification) GoHide(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hide", flags, ch)
}

func (v *interfaceNotification) Hide(flags dbus.Flags) error {
	return (<-v.GoHide(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method recordCount

func (v *interfaceNotification) GoRecordCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".recordCount", flags, ch)
}

func (*interfaceNotification) StoreRecordCount(call *dbus.Call) (arg0 uint32, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceNotification) RecordCount(flags dbus.Flags) (uint32, error) {
	return v.StoreRecordCount(
		<-v.GoRecordCount(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAppList

func (v *interfaceNotification) GoGetAppList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAppList", flags, ch)
}

func (*interfaceNotification) StoreGetAppList(call *dbus.Call) (arg0 []string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *interfaceNotification) GetAppList(flags dbus.Flags) ([]string, error) {
	return v.StoreGetAppList(
		<-v.GoGetAppList(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAppInfo

func (v *interfaceNotification) GoGetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAppInfo", flags, ch, arg0, arg1)
}

func (*interfaceNotification) StoreGetAppInfo(call *dbus.Call) (arg2 dbus.Variant, err error) {
	err = call.Store(&arg2)
	return
}

func (v *interfaceNotification) GetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32) (dbus.Variant, error) {
	return v.StoreGetAppInfo(
		<-v.GoGetAppInfo(flags, make(chan *dbus.Call, 1), arg0, arg1).Done)
}

// method GetSystemInfo

func (v *interfaceNotification) GoGetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSystemInfo", flags, ch, arg0)
}

func (*interfaceNotification) StoreGetSystemInfo(call *dbus.Call) (arg1 dbus.Variant, err error) {
	err = call.Store(&arg1)
	return
}

func (v *interfaceNotification) GetSystemInfo(flags dbus.Flags, arg0 uint32) (dbus.Variant, error) {
	return v.StoreGetSystemInfo(
		<-v.GoGetSystemInfo(flags, make(chan *dbus.Call, 1), arg0).Done)
}

// method SetAppInfo

func (v *interfaceNotification) GoSetAppInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 string, arg1 uint32, arg2 dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAppInfo", flags, ch, arg0, arg1, arg2)
}

func (v *interfaceNotification) SetAppInfo(flags dbus.Flags, arg0 string, arg1 uint32, arg2 dbus.Variant) error {
	return (<-v.GoSetAppInfo(flags, make(chan *dbus.Call, 1), arg0, arg1, arg2).Done).Err
}

// method SetSystemInfo

func (v *interfaceNotification) GoSetSystemInfo(flags dbus.Flags, ch chan *dbus.Call, arg0 uint32, arg1 dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSystemInfo", flags, ch, arg0, arg1)
}

func (v *interfaceNotification) SetSystemInfo(flags dbus.Flags, arg0 uint32, arg1 dbus.Variant) error {
	return (<-v.GoSetSystemInfo(flags, make(chan *dbus.Call, 1), arg0, arg1).Done).Err
}

// method setAppSetting

func (v *interfaceNotification) GoSetAppSetting(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".setAppSetting", flags, ch, arg0)
}

func (v *interfaceNotification) SetAppSetting(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetAppSetting(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// signal NotificationClosed

func (v *interfaceNotification) ConnectNotificationClosed(cb func(arg0 uint32, arg1 uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "NotificationClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".NotificationClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 uint32
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ActionInvoked

func (v *interfaceNotification) ConnectActionInvoked(cb func(arg0 uint32, arg1 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ActionInvoked", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ActionInvoked",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 string
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal RecordAdded

func (v *interfaceNotification) ConnectRecordAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "RecordAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".RecordAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AppInfoChanged

func (v *interfaceNotification) ConnectAppInfoChanged(cb func(arg0 string, arg1 uint32, arg2 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppInfoChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppInfoChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		var arg1 uint32
		var arg2 dbus.Variant
		err := dbus.Store(sig.Body, &arg0, &arg1, &arg2)
		if err == nil {
			cb(arg0, arg1, arg2)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal SystemInfoChanged

func (v *interfaceNotification) ConnectSystemInfoChanged(cb func(arg0 uint32, arg1 dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "SystemInfoChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".SystemInfoChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 uint32
		var arg1 dbus.Variant
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AppAddedSignal

func (v *interfaceNotification) ConnectAppAddedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppAddedSignal", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppAddedSignal",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal AppRemovedSignal

func (v *interfaceNotification) ConnectAppRemovedSignal(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AppRemovedSignal", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AppRemovedSignal",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal appRemoved

func (v *interfaceNotification) ConnectAppRemoved(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal appAdded

func (v *interfaceNotification) ConnectAppAdded(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal appSettingChanged

func (v *interfaceNotification) ConnectAppSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "appSettingChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".appSettingChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal systemSettingChanged

func (v *interfaceNotification) ConnectSystemSettingChanged(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "systemSettingChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".systemSettingChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property allSetting s

func (v *interfaceNotification) AllSetting() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "allSetting",
	}
}

// property systemSetting s

func (v *interfaceNotification) SystemSetting() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "systemSetting",
	}
}