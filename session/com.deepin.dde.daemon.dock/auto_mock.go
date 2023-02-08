// Code generated by "./generator ./session/com.deepin.dde.daemon.dock"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package dock

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockDock struct {
	MockInterfaceDock // interface com.deepin.dde.daemon.Dock
	proxy.MockObject
}

type MockInterfaceDock struct {
	mock.Mock
}

// method ActivateWindow

func (v *MockInterfaceDock) GoActivateWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) ActivateWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method CancelPreviewWindow

func (v *MockInterfaceDock) GoCancelPreviewWindow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) CancelPreviewWindow(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method CloseWindow

func (v *MockInterfaceDock) GoCloseWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) CloseWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method GetEntryIDs

func (v *MockInterfaceDock) GoGetEntryIDs(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) GetEntryIDs(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsDocked

func (v *MockInterfaceDock) GoIsDocked(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) IsDocked(flags dbus.Flags, desktopFile string) (bool, error) {
	mockArgs := v.Called(flags, desktopFile)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsOnDock

func (v *MockInterfaceDock) GoIsOnDock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) IsOnDock(flags dbus.Flags, desktopFile string) (bool, error) {
	mockArgs := v.Called(flags, desktopFile)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method MakeWindowAbove

func (v *MockInterfaceDock) GoMakeWindowAbove(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) MakeWindowAbove(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method MaximizeWindow

func (v *MockInterfaceDock) GoMaximizeWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) MaximizeWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method MinimizeWindow

func (v *MockInterfaceDock) GoMinimizeWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) MinimizeWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method MoveEntry

func (v *MockInterfaceDock) GoMoveEntry(flags dbus.Flags, ch chan *dbus.Call, index int32, newIndex int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, index, newIndex)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) MoveEntry(flags dbus.Flags, index int32, newIndex int32) error {
	mockArgs := v.Called(flags, index, newIndex)

	return mockArgs.Error(0)
}

// method MoveWindow

func (v *MockInterfaceDock) GoMoveWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) MoveWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method PreviewWindow

func (v *MockInterfaceDock) GoPreviewWindow(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) PreviewWindow(flags dbus.Flags, win uint32) error {
	mockArgs := v.Called(flags, win)

	return mockArgs.Error(0)
}

// method QueryWindowIdentifyMethod

func (v *MockInterfaceDock) GoQueryWindowIdentifyMethod(flags dbus.Flags, ch chan *dbus.Call, win uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, win)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) QueryWindowIdentifyMethod(flags dbus.Flags, win uint32) (string, error) {
	mockArgs := v.Called(flags, win)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RequestDock

func (v *MockInterfaceDock) GoRequestDock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string, index int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFile, index)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) RequestDock(flags dbus.Flags, desktopFile string, index int32) (bool, error) {
	mockArgs := v.Called(flags, desktopFile, index)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method RequestUndock

func (v *MockInterfaceDock) GoRequestUndock(flags dbus.Flags, ch chan *dbus.Call, desktopFile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, desktopFile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) RequestUndock(flags dbus.Flags, desktopFile string) (bool, error) {
	mockArgs := v.Called(flags, desktopFile)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method SetFrontendWindowRect

func (v *MockInterfaceDock) GoSetFrontendWindowRect(flags dbus.Flags, ch chan *dbus.Call, x int32, y int32, width uint32, height uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, x, y, width, height)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceDock) SetFrontendWindowRect(flags dbus.Flags, x int32, y int32, width uint32, height uint32) error {
	mockArgs := v.Called(flags, x, y, width, height)

	return mockArgs.Error(0)
}

// signal ServiceRestarted

func (v *MockInterfaceDock) ConnectServiceRestarted(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal EntryAdded

func (v *MockInterfaceDock) ConnectEntryAdded(cb func(path dbus.ObjectPath, index int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal EntryRemoved

func (v *MockInterfaceDock) ConnectEntryRemoved(cb func(entryId string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property ShowTimeout u

func (v *MockInterfaceDock) ShowTimeout() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HideTimeout u

func (v *MockInterfaceDock) HideTimeout() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropDockFrontendWindowRect struct {
	mock.Mock
}

func (p MockPropDockFrontendWindowRect) Get(flags dbus.Flags) (value FrontendWindowRect, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(FrontendWindowRect)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropDockFrontendWindowRect) Set(flags dbus.Flags, value FrontendWindowRect) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropDockFrontendWindowRect) ConnectChanged(cb func(hasValue bool, value FrontendWindowRect)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property FrontendWindowRect (iiuu)

func (v *MockInterfaceDock) FrontendWindowRect() PropDockFrontendWindowRect {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropDockFrontendWindowRect)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Entries ao

func (v *MockInterfaceDock) Entries() proxy.PropObjectPathArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPathArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HideMode i

func (v *MockInterfaceDock) HideMode() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DisplayMode i

func (v *MockInterfaceDock) DisplayMode() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HideState i

func (v *MockInterfaceDock) HideState() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Position i

func (v *MockInterfaceDock) Position() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IconSize u

func (v *MockInterfaceDock) IconSize() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DockedApps as

func (v *MockInterfaceDock) DockedApps() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockEntry struct {
	MockInterfaceEntry // interface com.deepin.dde.daemon.Dock.Entry
	proxy.MockObject
}

type MockInterfaceEntry struct {
	mock.Mock
}

// method Activate

func (v *MockInterfaceEntry) GoActivate(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, timestamp)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) Activate(flags dbus.Flags, timestamp uint32) error {
	mockArgs := v.Called(flags, timestamp)

	return mockArgs.Error(0)
}

// method Check

func (v *MockInterfaceEntry) GoCheck(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) Check(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ForceQuit

func (v *MockInterfaceEntry) GoForceQuit(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) ForceQuit(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method HandleDragDrop

func (v *MockInterfaceEntry) GoHandleDragDrop(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, files []string) *dbus.Call {
	mockArgs := v.Called(flags, ch, timestamp, files)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) HandleDragDrop(flags dbus.Flags, timestamp uint32, files []string) error {
	mockArgs := v.Called(flags, timestamp, files)

	return mockArgs.Error(0)
}

// method HandleMenuItem

func (v *MockInterfaceEntry) GoHandleMenuItem(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, timestamp, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) HandleMenuItem(flags dbus.Flags, timestamp uint32, id string) error {
	mockArgs := v.Called(flags, timestamp, id)

	return mockArgs.Error(0)
}

// method NewInstance

func (v *MockInterfaceEntry) GoNewInstance(flags dbus.Flags, ch chan *dbus.Call, timestamp uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, timestamp)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) NewInstance(flags dbus.Flags, timestamp uint32) error {
	mockArgs := v.Called(flags, timestamp)

	return mockArgs.Error(0)
}

// method PresentWindows

func (v *MockInterfaceEntry) GoPresentWindows(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) PresentWindows(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method RequestDock

func (v *MockInterfaceEntry) GoRequestDock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) RequestDock(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method RequestUndock

func (v *MockInterfaceEntry) GoRequestUndock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceEntry) RequestUndock(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// property Name s

func (v *MockInterfaceEntry) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Icon s

func (v *MockInterfaceEntry) Icon() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Id s

func (v *MockInterfaceEntry) Id() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsActive b

func (v *MockInterfaceEntry) IsActive() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentWindow u

func (v *MockInterfaceEntry) CurrentWindow() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property IsDocked b

func (v *MockInterfaceEntry) IsDocked() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockPropEntryWindowInfos struct {
	mock.Mock
}

func (p MockPropEntryWindowInfos) Get(flags dbus.Flags) (value map[uint32]WindowInfo, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(map[uint32]WindowInfo)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p MockPropEntryWindowInfos) Set(flags dbus.Flags, value map[uint32]WindowInfo) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p MockPropEntryWindowInfos) ConnectChanged(cb func(hasValue bool, value map[uint32]WindowInfo)) error {
	args := p.Called(cb)

	return args.Error(0)
}

// property WindowInfos a{u(sb)}

func (v *MockInterfaceEntry) WindowInfos() PropEntryWindowInfos {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*MockPropEntryWindowInfos)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Menu s

func (v *MockInterfaceEntry) Menu() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DesktopFile s

func (v *MockInterfaceEntry) DesktopFile() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
