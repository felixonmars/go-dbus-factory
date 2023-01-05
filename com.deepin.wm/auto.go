// Code generated by "./generator ./com.deepin.wm"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package wm

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
)

type Wm interface {
	wm // interface com.deepin.wm
	proxy.Object
}

type objectWm struct {
	interfaceWm // interface com.deepin.wm
	proxy.ImplObject
}

func NewWm(conn *dbus.Conn) Wm {
	obj := new(objectWm)
	obj.ImplObject.Init_(conn, "com.deepin.wm", "/com/deepin/wm")
	return obj
}

type wm interface {
	GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call
	SwitchApplication(flags dbus.Flags, backward bool) error
	GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call
	TileActiveWindow(flags dbus.Flags, side uint32) error
	GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	BeginToMoveActiveWindow(flags dbus.Flags) error
	GoToggleActiveWindowMaximize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ToggleActiveWindowMaximize(flags dbus.Flags) error
	GoMinimizeActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	MinimizeActiveWindow(flags dbus.Flags) error
	GoShowWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ShowWorkspace(flags dbus.Flags) error
	GoShowWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ShowWindow(flags dbus.Flags) error
	GoShowAllWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ShowAllWindow(flags dbus.Flags) error
	GoClearMoveStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	ClearMoveStatus(flags dbus.Flags) error
	GoTouchToMove(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32) *dbus.Call
	TouchToMove(flags dbus.Flags, x int32, y int32) error
	GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call
	PerformAction(flags dbus.Flags, type0 int32) error
	GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call
	PreviewWindow(flags dbus.Flags, xid uint32) error
	GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	CancelPreviewWindow(flags dbus.Flags) error
	GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetCurrentWorkspaceBackground(flags dbus.Flags) (string, error)
	GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call
	SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error
	GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call
	GetWorkspaceBackground(flags dbus.Flags, index int32) (string, error)
	GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call
	SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error
	GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call
	SetTransientBackground(flags dbus.Flags, arg0 string) error
	GoGetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call
	GetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, strMonitorName string) (string, error)
	GoSetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call
	SetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error
	GoGetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call
	GetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string) (string, error)
	GoSetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call
	SetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string, uri string) error
	GoSetTransientBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call
	SetTransientBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error
	GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetCurrentWorkspace(flags dbus.Flags) (int32, error)
	GoWorkspaceCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	WorkspaceCount(flags dbus.Flags) (int32, error)
	GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call
	SetCurrentWorkspace(flags dbus.Flags, index int32) error
	GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	PreviousWorkspace(flags dbus.Flags) error
	GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	NextWorkspace(flags dbus.Flags) error
	GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetAllAccels(flags dbus.Flags) (string, error)
	GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	GetAccel(flags dbus.Flags, id string) ([]string, error)
	GoGetDefaultAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	GetDefaultAccel(flags dbus.Flags, id string) ([]string, error)
	GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call
	SetAccel(flags dbus.Flags, data string) (bool, error)
	GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call
	RemoveAccel(flags dbus.Flags, id string) error
	GoSetDecorationTheme(flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call
	SetDecorationTheme(flags dbus.Flags, themeType string, themeName string) error
	GoSetDecorationDeepinTheme(flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call
	SetDecorationDeepinTheme(flags dbus.Flags, deepinThemeName string) error
	GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call
	ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error
	GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call
	SwitchToWorkspace(flags dbus.Flags, backward bool) error
	GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call
	PresentWindows(flags dbus.Flags, xids []uint32) error
	GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	EnableZoneDetected(flags dbus.Flags, enabled bool) error
	GoGetIsShowDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetIsShowDesktop(flags dbus.Flags) (bool, error)
	GoGetMultiTaskingStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	GetMultiTaskingStatus(flags dbus.Flags) (bool, error)
	ConnectWorkspaceBackgroundChanged(cb func(index int32, newUri string)) (dbusutil.SignalHandlerId, error)
	ConnectWorkspaceBackgroundChangedForMonitor(cb func(index int32, monitor string, newUri string)) (dbusutil.SignalHandlerId, error)
	ConnectCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error)
	ConnectWmCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error)
	ConnectWorkspaceCountChanged(cb func(count int32)) (dbusutil.SignalHandlerId, error)
	ConnectWorkspaceSwitched(cb func(from int32, to int32)) (dbusutil.SignalHandlerId, error)
	CompositingEnabled() proxy.PropBool
	CompositingPossible() proxy.PropBool
	CompositingAllowSwitch() proxy.PropBool
	ZoneEnabled() proxy.PropBool
	CursorTheme() proxy.PropString
	CursorSize() proxy.PropInt32
}

type interfaceWm struct{}

func (v *interfaceWm) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceWm) GetInterfaceName_() string {
	return "com.deepin.wm"
}

// method SwitchApplication

func (v *interfaceWm) GoSwitchApplication(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchApplication", flags, ch, backward)
}

func (v *interfaceWm) SwitchApplication(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchApplication(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

// method TileActiveWindow

func (v *interfaceWm) GoTileActiveWindow(flags dbus.Flags, ch chan *dbus.Call, side uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TileActiveWindow", flags, ch, side)
}

func (v *interfaceWm) TileActiveWindow(flags dbus.Flags, side uint32) error {
	return (<-v.GoTileActiveWindow(flags, make(chan *dbus.Call, 1), side).Done).Err
}

// method BeginToMoveActiveWindow

func (v *interfaceWm) GoBeginToMoveActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".BeginToMoveActiveWindow", flags, ch)
}

func (v *interfaceWm) BeginToMoveActiveWindow(flags dbus.Flags) error {
	return (<-v.GoBeginToMoveActiveWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ToggleActiveWindowMaximize

func (v *interfaceWm) GoToggleActiveWindowMaximize(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ToggleActiveWindowMaximize", flags, ch)
}

func (v *interfaceWm) ToggleActiveWindowMaximize(flags dbus.Flags) error {
	return (<-v.GoToggleActiveWindowMaximize(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method MinimizeActiveWindow

func (v *interfaceWm) GoMinimizeActiveWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MinimizeActiveWindow", flags, ch)
}

func (v *interfaceWm) MinimizeActiveWindow(flags dbus.Flags) error {
	return (<-v.GoMinimizeActiveWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowWorkspace

func (v *interfaceWm) GoShowWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowWorkspace", flags, ch)
}

func (v *interfaceWm) ShowWorkspace(flags dbus.Flags) error {
	return (<-v.GoShowWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowWindow

func (v *interfaceWm) GoShowWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowWindow", flags, ch)
}

func (v *interfaceWm) ShowWindow(flags dbus.Flags) error {
	return (<-v.GoShowWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ShowAllWindow

func (v *interfaceWm) GoShowAllWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ShowAllWindow", flags, ch)
}

func (v *interfaceWm) ShowAllWindow(flags dbus.Flags) error {
	return (<-v.GoShowAllWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ClearMoveStatus

func (v *interfaceWm) GoClearMoveStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClearMoveStatus", flags, ch)
}

func (v *interfaceWm) ClearMoveStatus(flags dbus.Flags) error {
	return (<-v.GoClearMoveStatus(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method TouchToMove

func (v *interfaceWm) GoTouchToMove(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TouchToMove", flags, ch, x, y)
}

func (v *interfaceWm) TouchToMove(flags dbus.Flags, x int32, y int32) error {
	return (<-v.GoTouchToMove(flags, make(chan *dbus.Call, 1), x, y).Done).Err
}

// method PerformAction

func (v *interfaceWm) GoPerformAction(flags dbus.Flags, ch chan *dbus.Call, type0 int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PerformAction", flags, ch, type0)
}

func (v *interfaceWm) PerformAction(flags dbus.Flags, type0 int32) error {
	return (<-v.GoPerformAction(flags, make(chan *dbus.Call, 1), type0).Done).Err
}

// method PreviewWindow

func (v *interfaceWm) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, xid uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviewWindow", flags, ch, xid)
}

func (v *interfaceWm) PreviewWindow(flags dbus.Flags, xid uint32) error {
	return (<-v.GoPreviewWindow(flags, make(chan *dbus.Call, 1), xid).Done).Err
}

// method CancelPreviewWindow

func (v *interfaceWm) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CancelPreviewWindow", flags, ch)
}

func (v *interfaceWm) CancelPreviewWindow(flags dbus.Flags) error {
	return (<-v.GoCancelPreviewWindow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetCurrentWorkspaceBackground

func (v *interfaceWm) GoGetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspaceBackground", flags, ch)
}

func (*interfaceWm) StoreGetCurrentWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceWm) GetCurrentWorkspaceBackground(flags dbus.Flags) (string, error) {
	return v.StoreGetCurrentWorkspaceBackground(
		<-v.GoGetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetCurrentWorkspaceBackground

func (v *interfaceWm) GoSetCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *interfaceWm) SetCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoSetCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method GetWorkspaceBackground

func (v *interfaceWm) GoGetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWorkspaceBackground", flags, ch, index)
}

func (*interfaceWm) StoreGetWorkspaceBackground(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceWm) GetWorkspaceBackground(flags dbus.Flags, index int32) (string, error) {
	return v.StoreGetWorkspaceBackground(
		<-v.GoGetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index).Done)
}

// method SetWorkspaceBackground

func (v *interfaceWm) GoSetWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, index int32, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWorkspaceBackground", flags, ch, index, uri)
}

func (v *interfaceWm) SetWorkspaceBackground(flags dbus.Flags, index int32, uri string) error {
	return (<-v.GoSetWorkspaceBackground(flags, make(chan *dbus.Call, 1), index, uri).Done).Err
}

// method SetTransientBackground

func (v *interfaceWm) GoSetTransientBackground(flags dbus.Flags, ch chan *dbus.Call, arg0 string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTransientBackground", flags, ch, arg0)
}

func (v *interfaceWm) SetTransientBackground(flags dbus.Flags, arg0 string) error {
	return (<-v.GoSetTransientBackground(flags, make(chan *dbus.Call, 1), arg0).Done).Err
}

// method GetCurrentWorkspaceBackgroundForMonitor

func (v *interfaceWm) GoGetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspaceBackgroundForMonitor", flags, ch, strMonitorName)
}

func (*interfaceWm) StoreGetCurrentWorkspaceBackgroundForMonitor(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceWm) GetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, strMonitorName string) (string, error) {
	return v.StoreGetCurrentWorkspaceBackgroundForMonitor(
		<-v.GoGetCurrentWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), strMonitorName).Done)
}

// method SetCurrentWorkspaceBackgroundForMonitor

func (v *interfaceWm) GoSetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspaceBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *interfaceWm) SetCurrentWorkspaceBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	return (<-v.GoSetCurrentWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done).Err
}

// method GetWorkspaceBackgroundForMonitor

func (v *interfaceWm) GoGetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName)
}

func (*interfaceWm) StoreGetWorkspaceBackgroundForMonitor(call *dbus.Call) (result string, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceWm) GetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string) (string, error) {
	return v.StoreGetWorkspaceBackgroundForMonitor(
		<-v.GoGetWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), index, strMonitorName).Done)
}

// method SetWorkspaceBackgroundForMonitor

func (v *interfaceWm) GoSetWorkspaceBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, index int32, strMonitorName string, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetWorkspaceBackgroundForMonitor", flags, ch, index, strMonitorName, uri)
}

func (v *interfaceWm) SetWorkspaceBackgroundForMonitor(flags dbus.Flags, index int32, strMonitorName string, uri string) error {
	return (<-v.GoSetWorkspaceBackgroundForMonitor(flags, make(chan *dbus.Call, 1), index, strMonitorName, uri).Done).Err
}

// method SetTransientBackgroundForMonitor

func (v *interfaceWm) GoSetTransientBackgroundForMonitor(flags dbus.Flags, ch chan *dbus.Call, uri string, strMonitorName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTransientBackgroundForMonitor", flags, ch, uri, strMonitorName)
}

func (v *interfaceWm) SetTransientBackgroundForMonitor(flags dbus.Flags, uri string, strMonitorName string) error {
	return (<-v.GoSetTransientBackgroundForMonitor(flags, make(chan *dbus.Call, 1), uri, strMonitorName).Done).Err
}

// method GetCurrentWorkspace

func (v *interfaceWm) GoGetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCurrentWorkspace", flags, ch)
}

func (*interfaceWm) StoreGetCurrentWorkspace(call *dbus.Call) (index int32, err error) {
	err = call.Store(&index)
	return
}

func (v *interfaceWm) GetCurrentWorkspace(flags dbus.Flags) (int32, error) {
	return v.StoreGetCurrentWorkspace(
		<-v.GoGetCurrentWorkspace(flags, make(chan *dbus.Call, 1)).Done)
}

// method WorkspaceCount

func (v *interfaceWm) GoWorkspaceCount(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".WorkspaceCount", flags, ch)
}

func (*interfaceWm) StoreWorkspaceCount(call *dbus.Call) (count int32, err error) {
	err = call.Store(&count)
	return
}

func (v *interfaceWm) WorkspaceCount(flags dbus.Flags) (int32, error) {
	return v.StoreWorkspaceCount(
		<-v.GoWorkspaceCount(flags, make(chan *dbus.Call, 1)).Done)
}

// method SetCurrentWorkspace

func (v *interfaceWm) GoSetCurrentWorkspace(flags dbus.Flags, ch chan *dbus.Call, index int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCurrentWorkspace", flags, ch, index)
}

func (v *interfaceWm) SetCurrentWorkspace(flags dbus.Flags, index int32) error {
	return (<-v.GoSetCurrentWorkspace(flags, make(chan *dbus.Call, 1), index).Done).Err
}

// method PreviousWorkspace

func (v *interfaceWm) GoPreviousWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PreviousWorkspace", flags, ch)
}

func (v *interfaceWm) PreviousWorkspace(flags dbus.Flags) error {
	return (<-v.GoPreviousWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method NextWorkspace

func (v *interfaceWm) GoNextWorkspace(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".NextWorkspace", flags, ch)
}

func (v *interfaceWm) NextWorkspace(flags dbus.Flags) error {
	return (<-v.GoNextWorkspace(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method GetAllAccels

func (v *interfaceWm) GoGetAllAccels(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAllAccels", flags, ch)
}

func (*interfaceWm) StoreGetAllAccels(call *dbus.Call) (data string, err error) {
	err = call.Store(&data)
	return
}

func (v *interfaceWm) GetAllAccels(flags dbus.Flags) (string, error) {
	return v.StoreGetAllAccels(
		<-v.GoGetAllAccels(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetAccel

func (v *interfaceWm) GoGetAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetAccel", flags, ch, id)
}

func (*interfaceWm) StoreGetAccel(call *dbus.Call) (data []string, err error) {
	err = call.Store(&data)
	return
}

func (v *interfaceWm) GetAccel(flags dbus.Flags, id string) ([]string, error) {
	return v.StoreGetAccel(
		<-v.GoGetAccel(flags, make(chan *dbus.Call, 1), id).Done)
}

// method GetDefaultAccel

func (v *interfaceWm) GoGetDefaultAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultAccel", flags, ch, id)
}

func (*interfaceWm) StoreGetDefaultAccel(call *dbus.Call) (data []string, err error) {
	err = call.Store(&data)
	return
}

func (v *interfaceWm) GetDefaultAccel(flags dbus.Flags, id string) ([]string, error) {
	return v.StoreGetDefaultAccel(
		<-v.GoGetDefaultAccel(flags, make(chan *dbus.Call, 1), id).Done)
}

// method SetAccel

func (v *interfaceWm) GoSetAccel(flags dbus.Flags, ch chan *dbus.Call, data string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAccel", flags, ch, data)
}

func (*interfaceWm) StoreSetAccel(call *dbus.Call) (result bool, err error) {
	err = call.Store(&result)
	return
}

func (v *interfaceWm) SetAccel(flags dbus.Flags, data string) (bool, error) {
	return v.StoreSetAccel(
		<-v.GoSetAccel(flags, make(chan *dbus.Call, 1), data).Done)
}

// method RemoveAccel

func (v *interfaceWm) GoRemoveAccel(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveAccel", flags, ch, id)
}

func (v *interfaceWm) RemoveAccel(flags dbus.Flags, id string) error {
	return (<-v.GoRemoveAccel(flags, make(chan *dbus.Call, 1), id).Done).Err
}

// method SetDecorationTheme

func (v *interfaceWm) GoSetDecorationTheme(flags dbus.Flags, ch chan *dbus.Call, themeType string, themeName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDecorationTheme", flags, ch, themeType, themeName)
}

func (v *interfaceWm) SetDecorationTheme(flags dbus.Flags, themeType string, themeName string) error {
	return (<-v.GoSetDecorationTheme(flags, make(chan *dbus.Call, 1), themeType, themeName).Done).Err
}

// method SetDecorationDeepinTheme

func (v *interfaceWm) GoSetDecorationDeepinTheme(flags dbus.Flags, ch chan *dbus.Call, deepinThemeName string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetDecorationDeepinTheme", flags, ch, deepinThemeName)
}

func (v *interfaceWm) SetDecorationDeepinTheme(flags dbus.Flags, deepinThemeName string) error {
	return (<-v.GoSetDecorationDeepinTheme(flags, make(chan *dbus.Call, 1), deepinThemeName).Done).Err
}

// method ChangeCurrentWorkspaceBackground

func (v *interfaceWm) GoChangeCurrentWorkspaceBackground(flags dbus.Flags, ch chan *dbus.Call, uri string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeCurrentWorkspaceBackground", flags, ch, uri)
}

func (v *interfaceWm) ChangeCurrentWorkspaceBackground(flags dbus.Flags, uri string) error {
	return (<-v.GoChangeCurrentWorkspaceBackground(flags, make(chan *dbus.Call, 1), uri).Done).Err
}

// method SwitchToWorkspace

func (v *interfaceWm) GoSwitchToWorkspace(flags dbus.Flags, ch chan *dbus.Call, backward bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchToWorkspace", flags, ch, backward)
}

func (v *interfaceWm) SwitchToWorkspace(flags dbus.Flags, backward bool) error {
	return (<-v.GoSwitchToWorkspace(flags, make(chan *dbus.Call, 1), backward).Done).Err
}

// method PresentWindows

func (v *interfaceWm) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call, xids []uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PresentWindows", flags, ch, xids)
}

func (v *interfaceWm) PresentWindows(flags dbus.Flags, xids []uint32) error {
	return (<-v.GoPresentWindows(flags, make(chan *dbus.Call, 1), xids).Done).Err
}

// method EnableZoneDetected

func (v *interfaceWm) GoEnableZoneDetected(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableZoneDetected", flags, ch, enabled)
}

func (v *interfaceWm) EnableZoneDetected(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableZoneDetected(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method GetIsShowDesktop

func (v *interfaceWm) GoGetIsShowDesktop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetIsShowDesktop", flags, ch)
}

func (*interfaceWm) StoreGetIsShowDesktop(call *dbus.Call) (isShowDesktop bool, err error) {
	err = call.Store(&isShowDesktop)
	return
}

func (v *interfaceWm) GetIsShowDesktop(flags dbus.Flags) (bool, error) {
	return v.StoreGetIsShowDesktop(
		<-v.GoGetIsShowDesktop(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetMultiTaskingStatus

func (v *interfaceWm) GoGetMultiTaskingStatus(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetMultiTaskingStatus", flags, ch)
}

func (*interfaceWm) StoreGetMultiTaskingStatus(call *dbus.Call) (isActive bool, err error) {
	err = call.Store(&isActive)
	return
}

func (v *interfaceWm) GetMultiTaskingStatus(flags dbus.Flags) (bool, error) {
	return v.StoreGetMultiTaskingStatus(
		<-v.GoGetMultiTaskingStatus(flags, make(chan *dbus.Call, 1)).Done)
}

// signal WorkspaceBackgroundChanged

func (v *interfaceWm) ConnectWorkspaceBackgroundChanged(cb func(index int32, newUri string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceBackgroundChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceBackgroundChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var index int32
		var newUri string
		err := dbus.Store(sig.Body, &index, &newUri)
		if err == nil {
			cb(index, newUri)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WorkspaceBackgroundChangedForMonitor

func (v *interfaceWm) ConnectWorkspaceBackgroundChangedForMonitor(cb func(index int32, monitor string, newUri string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceBackgroundChangedForMonitor", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceBackgroundChangedForMonitor",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var index int32
		var monitor string
		var newUri string
		err := dbus.Store(sig.Body, &index, &monitor, &newUri)
		if err == nil {
			cb(index, monitor, newUri)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal compositingEnabledChanged

func (v *interfaceWm) ConnectCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "compositingEnabledChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".compositingEnabledChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var enabled bool
		err := dbus.Store(sig.Body, &enabled)
		if err == nil {
			cb(enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal wmCompositingEnabledChanged

func (v *interfaceWm) ConnectWmCompositingEnabledChanged(cb func(enabled bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "wmCompositingEnabledChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".wmCompositingEnabledChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var enabled bool
		err := dbus.Store(sig.Body, &enabled)
		if err == nil {
			cb(enabled)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal workspaceCountChanged

func (v *interfaceWm) ConnectWorkspaceCountChanged(cb func(count int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "workspaceCountChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".workspaceCountChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var count int32
		err := dbus.Store(sig.Body, &count)
		if err == nil {
			cb(count)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal WorkspaceSwitched

func (v *interfaceWm) ConnectWorkspaceSwitched(cb func(from int32, to int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "WorkspaceSwitched", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".WorkspaceSwitched",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var from int32
		var to int32
		err := dbus.Store(sig.Body, &from, &to)
		if err == nil {
			cb(from, to)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property compositingEnabled b

func (v *interfaceWm) CompositingEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "compositingEnabled",
	}
}

// property compositingPossible b

func (v *interfaceWm) CompositingPossible() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "compositingPossible",
	}
}

// property compositingAllowSwitch b

func (v *interfaceWm) CompositingAllowSwitch() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "compositingAllowSwitch",
	}
}

// property zoneEnabled b

func (v *interfaceWm) ZoneEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "zoneEnabled",
	}
}

// property cursorTheme s

func (v *interfaceWm) CursorTheme() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "cursorTheme",
	}
}

// property cursorSize i

func (v *interfaceWm) CursorSize() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "cursorSize",
	}
}
