// Code generated by "./generator ./com.deepin.daemon.fprintd"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package fprintd

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Fprintd interface {
	fprintd // interface com.deepin.daemon.Fprintd
	proxy.Object
}

type objectFprintd struct {
	interfaceFprintd // interface com.deepin.daemon.Fprintd
	proxy.ImplObject
}

func NewFprintd(conn *dbus.Conn) Fprintd {
	obj := new(objectFprintd)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Fprintd", "/com/deepin/daemon/Fprintd")
	return obj
}

type fprintd interface {
	GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error)
	GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error)
	GoTriggerUDevEvent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	TriggerUDevEvent(flags dbus.Flags) error
	Devices() proxy.PropObjectPathArray
}

type interfaceFprintd struct{}

func (v *interfaceFprintd) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceFprintd) GetInterfaceName_() string {
	return "com.deepin.daemon.Fprintd"
}

// method GetDefaultDevice

func (v *interfaceFprintd) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (*interfaceFprintd) StoreGetDefaultDevice(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *interfaceFprintd) GetDefaultDevice(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreGetDefaultDevice(
		<-v.GoGetDefaultDevice(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetDevices

func (v *interfaceFprintd) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*interfaceFprintd) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *interfaceFprintd) GetDevices(flags dbus.Flags) ([]dbus.ObjectPath, error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

// method TriggerUDevEvent

func (v *interfaceFprintd) GoTriggerUDevEvent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TriggerUDevEvent", flags, ch)
}

func (v *interfaceFprintd) TriggerUDevEvent(flags dbus.Flags) error {
	return (<-v.GoTriggerUDevEvent(flags, make(chan *dbus.Call, 1)).Done).Err
}

// property Devices ao

func (v *interfaceFprintd) Devices() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

type Device interface {
	device // interface com.deepin.daemon.Fprintd.Device
	proxy.Object
}

type objectDevice struct {
	interfaceDevice // interface com.deepin.daemon.Fprintd.Device
	proxy.ImplObject
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectDevice)
	obj.ImplObject.Init_(conn, "com.deepin.daemon.Fprintd", path)
	return obj, nil
}

type device interface {
	GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	Claim(flags dbus.Flags, username string) error
	GoClaimForce(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	ClaimForce(flags dbus.Flags, username string) error
	GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call
	DeleteEnrolledFinger(flags dbus.Flags, username string, finger string) error
	GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	DeleteEnrolledFingers(flags dbus.Flags, username string) error
	GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call
	EnrollStart(flags dbus.Flags, finger string) error
	GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	EnrollStop(flags dbus.Flags) error
	GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetCapabilities(flags dbus.Flags) ([]string, error)
	GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call
	ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error)
	GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Release(flags dbus.Flags) error
	GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call
	VerifyStart(flags dbus.Flags, finger string) error
	GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	VerifyStop(flags dbus.Flags) error
	ConnectEnrollStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error)
	ConnectVerifyFingerSelected(cb func(finger string)) (dbusutil.SignalHandlerId, error)
	ScanType() proxy.PropString
}

type interfaceDevice struct{}

func (v *interfaceDevice) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceDevice) GetInterfaceName_() string {
	return "com.deepin.daemon.Fprintd.Device"
}

// method Claim

func (v *interfaceDevice) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, username)
}

func (v *interfaceDevice) Claim(flags dbus.Flags, username string) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method ClaimForce

func (v *interfaceDevice) GoClaimForce(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimForce", flags, ch, username)
}

func (v *interfaceDevice) ClaimForce(flags dbus.Flags, username string) error {
	return (<-v.GoClaimForce(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method DeleteEnrolledFinger

func (v *interfaceDevice) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger)
}

func (v *interfaceDevice) DeleteEnrolledFinger(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoDeleteEnrolledFinger(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

// method DeleteEnrolledFingers

func (v *interfaceDevice) GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFingers", flags, ch, username)
}

func (v *interfaceDevice) DeleteEnrolledFingers(flags dbus.Flags, username string) error {
	return (<-v.GoDeleteEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done).Err
}

// method EnrollStart

func (v *interfaceDevice) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStart", flags, ch, finger)
}

func (v *interfaceDevice) EnrollStart(flags dbus.Flags, finger string) error {
	return (<-v.GoEnrollStart(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// method EnrollStop

func (v *interfaceDevice) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStop", flags, ch)
}

func (v *interfaceDevice) EnrollStop(flags dbus.Flags) error {
	return (<-v.GoEnrollStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetCapabilities

func (v *interfaceDevice) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapabilities", flags, ch)
}

func (*interfaceDevice) StoreGetCapabilities(call *dbus.Call) (caps []string, err error) {
	err = call.Store(&caps)
	return
}

func (v *interfaceDevice) GetCapabilities(flags dbus.Flags) ([]string, error) {
	return v.StoreGetCapabilities(
		<-v.GoGetCapabilities(flags, make(chan *dbus.Call, 1)).Done)
}

// method ListEnrolledFingers

func (v *interfaceDevice) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (*interfaceDevice) StoreListEnrolledFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *interfaceDevice) ListEnrolledFingers(flags dbus.Flags, username string) ([]string, error) {
	return v.StoreListEnrolledFingers(
		<-v.GoListEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

// method Release

func (v *interfaceDevice) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Release", flags, ch)
}

func (v *interfaceDevice) Release(flags dbus.Flags) error {
	return (<-v.GoRelease(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method VerifyStart

func (v *interfaceDevice) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStart", flags, ch, finger)
}

func (v *interfaceDevice) VerifyStart(flags dbus.Flags, finger string) error {
	return (<-v.GoVerifyStart(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

// method VerifyStop

func (v *interfaceDevice) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStop", flags, ch)
}

func (v *interfaceDevice) VerifyStop(flags dbus.Flags) error {
	return (<-v.GoVerifyStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal EnrollStatus

func (v *interfaceDevice) ConnectEnrollStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var done bool
		err := dbus.Store(sig.Body, &status, &done)
		if err == nil {
			cb(status, done)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *interfaceDevice) ConnectVerifyStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var done bool
		err := dbus.Store(sig.Body, &status, &done)
		if err == nil {
			cb(status, done)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyFingerSelected

func (v *interfaceDevice) ConnectVerifyFingerSelected(cb func(finger string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyFingerSelected", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyFingerSelected",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var finger string
		err := dbus.Store(sig.Body, &finger)
		if err == nil {
			cb(finger)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ScanType s

func (v *interfaceDevice) ScanType() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ScanType",
	}
}
