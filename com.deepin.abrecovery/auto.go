// Code generated by "./generator ./com.deepin.abrecovery"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package abrecovery

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type ABRecovery interface {
	abRecovery // interface com.deepin.ABRecovery
	proxy.Object
}

type objectABRecovery struct {
	interfaceAbRecovery // interface com.deepin.ABRecovery
	proxy.ImplObject
}

func NewABRecovery(conn *dbus.Conn) ABRecovery {
	obj := new(objectABRecovery)
	obj.ImplObject.Init_(conn, "com.deepin.ABRecovery", "/com/deepin/ABRecovery")
	return obj
}

type abRecovery interface {
	GoCanBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanBackup(flags dbus.Flags) (bool, error)
	GoCanRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CanRestore(flags dbus.Flags) (bool, error)
	GoStartBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StartBackup(flags dbus.Flags) error
	GoStartRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	StartRestore(flags dbus.Flags) error
	ConnectJobEnd(cb func(kind string, success bool, errMsg string)) (dbusutil.SignalHandlerId, error)
	BackingUp() proxy.PropBool
	Restoring() proxy.PropBool
	ConfigValid() proxy.PropBool
	HasBackedUp() proxy.PropBool
}

type interfaceAbRecovery struct{}

func (v *interfaceAbRecovery) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceAbRecovery) GetInterfaceName_() string {
	return "com.deepin.ABRecovery"
}

// method CanBackup

func (v *interfaceAbRecovery) GoCanBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanBackup", flags, ch)
}

func (*interfaceAbRecovery) StoreCanBackup(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfaceAbRecovery) CanBackup(flags dbus.Flags) (bool, error) {
	return v.StoreCanBackup(
		<-v.GoCanBackup(flags, make(chan *dbus.Call, 1)).Done)
}

// method CanRestore

func (v *interfaceAbRecovery) GoCanRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CanRestore", flags, ch)
}

func (*interfaceAbRecovery) StoreCanRestore(call *dbus.Call) (can bool, err error) {
	err = call.Store(&can)
	return
}

func (v *interfaceAbRecovery) CanRestore(flags dbus.Flags) (bool, error) {
	return v.StoreCanRestore(
		<-v.GoCanRestore(flags, make(chan *dbus.Call, 1)).Done)
}

// method StartBackup

func (v *interfaceAbRecovery) GoStartBackup(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartBackup", flags, ch)
}

func (v *interfaceAbRecovery) StartBackup(flags dbus.Flags) error {
	return (<-v.GoStartBackup(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method StartRestore

func (v *interfaceAbRecovery) GoStartRestore(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartRestore", flags, ch)
}

func (v *interfaceAbRecovery) StartRestore(flags dbus.Flags) error {
	return (<-v.GoStartRestore(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal JobEnd

func (v *interfaceAbRecovery) ConnectJobEnd(cb func(kind string, success bool, errMsg string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "JobEnd", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".JobEnd",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var kind string
		var success bool
		var errMsg string
		err := dbus.Store(sig.Body, &kind, &success, &errMsg)
		if err == nil {
			cb(kind, success, errMsg)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property BackingUp b

func (v *interfaceAbRecovery) BackingUp() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "BackingUp",
	}
}

// property Restoring b

func (v *interfaceAbRecovery) Restoring() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Restoring",
	}
}

// property ConfigValid b

func (v *interfaceAbRecovery) ConfigValid() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "ConfigValid",
	}
}

// property HasBackedUp b

func (v *interfaceAbRecovery) HasBackedUp() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "HasBackedUp",
	}
}
