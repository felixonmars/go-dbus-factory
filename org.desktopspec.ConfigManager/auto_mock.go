// Code generated by "./generator ./org.desktopspec.ConfigManager"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package ConfigManager

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockConfigManager struct {
	MockInterfaceConfigManager // interface org.desktopspec.ConfigManager
	proxy.MockObject
}

type MockInterfaceConfigManager struct {
	mock.Mock
}

// method acquireManager

func (v *MockInterfaceConfigManager) GoAcquireManager(flags dbus.Flags, ch chan *dbus.Call, appid string, name string, subpath string) *dbus.Call {
	mockArgs := v.Called(flags, ch, appid, name, subpath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceConfigManager) AcquireManager(flags dbus.Flags, appid string, name string, subpath string) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, appid, name, subpath)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method update

func (v *MockInterfaceConfigManager) GoUpdate(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call {
	mockArgs := v.Called(flags, ch, path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceConfigManager) Update(flags dbus.Flags, path string) error {
	mockArgs := v.Called(flags, path)

	return mockArgs.Error(0)
}

// method sync

func (v *MockInterfaceConfigManager) GoSync(flags dbus.Flags, ch chan *dbus.Call, path string) *dbus.Call {
	mockArgs := v.Called(flags, ch, path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceConfigManager) Sync(flags dbus.Flags, path string) error {
	mockArgs := v.Called(flags, path)

	return mockArgs.Error(0)
}

// method setDelayReleaseTime

func (v *MockInterfaceConfigManager) GoSetDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call, time int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, time)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceConfigManager) SetDelayReleaseTime(flags dbus.Flags, time int32) error {
	mockArgs := v.Called(flags, time)

	return mockArgs.Error(0)
}

// method delayReleaseTime

func (v *MockInterfaceConfigManager) GoDelayReleaseTime(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceConfigManager) DelayReleaseTime(flags dbus.Flags) (int32, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

type MockManager struct {
	MockInterfaceManager // interface org.desktopspec.ConfigManager.Manager
	proxy.MockObject
}

type MockInterfaceManager struct {
	mock.Mock
}

// method value

func (v *MockInterfaceManager) GoValue(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Value(flags dbus.Flags, key string) (dbus.Variant, error) {
	mockArgs := v.Called(flags, key)

	ret0, ok := mockArgs.Get(0).(dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method setValue

func (v *MockInterfaceManager) GoSetValue(flags dbus.Flags, ch chan *dbus.Call, key string, value dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, key, value)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) SetValue(flags dbus.Flags, key string, value dbus.Variant) error {
	mockArgs := v.Called(flags, key, value)

	return mockArgs.Error(0)
}

// method reset

func (v *MockInterfaceManager) GoReset(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Reset(flags dbus.Flags, key string) error {
	mockArgs := v.Called(flags, key)

	return mockArgs.Error(0)
}

// method name

func (v *MockInterfaceManager) GoName(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key, language)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Name(flags dbus.Flags, key string, language string) (string, error) {
	mockArgs := v.Called(flags, key, language)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method description

func (v *MockInterfaceManager) GoDescription(flags dbus.Flags, ch chan *dbus.Call, key string, language string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key, language)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Description(flags dbus.Flags, key string, language string) (string, error) {
	mockArgs := v.Called(flags, key, language)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method visibility

func (v *MockInterfaceManager) GoVisibility(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Visibility(flags dbus.Flags, key string) (string, error) {
	mockArgs := v.Called(flags, key)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method permissions

func (v *MockInterfaceManager) GoPermissions(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Permissions(flags dbus.Flags, key string) (string, error) {
	mockArgs := v.Called(flags, key)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method flags

func (v *MockInterfaceManager) GoFlags(flags dbus.Flags, ch chan *dbus.Call, key string) *dbus.Call {
	mockArgs := v.Called(flags, ch, key)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Flags(flags dbus.Flags, key string) (int32, error) {
	mockArgs := v.Called(flags, key)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method release

func (v *MockInterfaceManager) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceManager) Release(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// signal valueChanged

func (v *MockInterfaceManager) ConnectValueChanged(cb func(key string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property version s

func (v *MockInterfaceManager) Version() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property keyList as

func (v *MockInterfaceManager) KeyList() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}