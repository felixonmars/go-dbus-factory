// Code generated by "./generator ./system/com.deepin.api.localehelper"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package localehelper

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type LocaleHelper interface {
	localeHelper // interface com.deepin.api.LocaleHelper
	proxy.Object
}

type objectLocaleHelper struct {
	interfaceLocaleHelper // interface com.deepin.api.LocaleHelper
	proxy.ImplObject
}

func NewLocaleHelper(conn *dbus.Conn) LocaleHelper {
	obj := new(objectLocaleHelper)
	obj.ImplObject.Init_(conn, "com.deepin.api.LocaleHelper", "/com/deepin/api/LocaleHelper")
	return obj
}

type localeHelper interface {
	GoGenerateLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call
	GenerateLocale(flags dbus.Flags, locale string) error
	GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call
	SetLocale(flags dbus.Flags, locale string) error
	ConnectSuccess(cb func(ok bool, reason string)) (dbusutil.SignalHandlerId, error)
}

type interfaceLocaleHelper struct{}

func (v *interfaceLocaleHelper) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLocaleHelper) GetInterfaceName_() string {
	return "com.deepin.api.LocaleHelper"
}

// method GenerateLocale

func (v *interfaceLocaleHelper) GoGenerateLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GenerateLocale", flags, ch, locale)
}

func (v *interfaceLocaleHelper) GenerateLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoGenerateLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// method SetLocale

func (v *interfaceLocaleHelper) GoSetLocale(flags dbus.Flags, ch chan *dbus.Call, locale string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLocale", flags, ch, locale)
}

func (v *interfaceLocaleHelper) SetLocale(flags dbus.Flags, locale string) error {
	return (<-v.GoSetLocale(flags, make(chan *dbus.Call, 1), locale).Done).Err
}

// signal Success

func (v *interfaceLocaleHelper) ConnectSuccess(cb func(ok bool, reason string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Success", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Success",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ok bool
		var reason string
		err := dbus.Store(sig.Body, &ok, &reason)
		if err == nil {
			cb(ok, reason)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
