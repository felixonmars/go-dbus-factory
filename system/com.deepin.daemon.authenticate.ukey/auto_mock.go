// Code generated by "./generator ./system/com.deepin.daemon.authenticate.ukey"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package ukey

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockUKey struct {
	MockInterfaceUkey // interface com.deepin.daemon.Authenticate.UKey.Device
	proxy.MockObject
}

type MockInterfaceUkey struct {
	mock.Mock
}

func (v *MockInterfaceUkey) SetInterfaceName_(string) {
}

// method GetPINLength

func (v *MockInterfaceUkey) GoGetPINLength(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetPINLength(flags dbus.Flags, uuid string) (int32, error) {
	mockArgs := v.Called(flags, uuid)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetUserList

func (v *MockInterfaceUkey) GoGetUserList(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetUserList(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetPin

func (v *MockInterfaceUkey) GoSetPin(flags dbus.Flags, ch chan *dbus.Call, uuid string, gid string, pin string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, gid, pin)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetPin(flags dbus.Flags, uuid string, gid string, pin string) error {
	mockArgs := v.Called(flags, uuid, gid, pin)

	return mockArgs.Error(0)
}

// method SetSessionPath

func (v *MockInterfaceUkey) GoSetSessionPath(flags dbus.Flags, ch chan *dbus.Call, uuid string, gid string, path string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, gid, path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetSessionPath(flags dbus.Flags, uuid string, gid string, path string) error {
	mockArgs := v.Called(flags, uuid, gid, path)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceUkey) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, uuid string, gid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, gid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) StopVerify(flags dbus.Flags, uuid string, gid string) error {
	mockArgs := v.Called(flags, uuid, gid)

	return mockArgs.Error(0)
}

// method Verify

func (v *MockInterfaceUkey) GoVerify(flags dbus.Flags, ch chan *dbus.Call, uuid string, gid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, gid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) Verify(flags dbus.Flags, uuid string, gid string) error {
	mockArgs := v.Called(flags, uuid, gid)

	return mockArgs.Error(0)
}

// signal VerifyResult

func (v *MockInterfaceUkey) ConnectVerifyResult(cb func(id string, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Name s

func (v *MockInterfaceUkey) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State i

func (v *MockInterfaceUkey) State() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type i

func (v *MockInterfaceUkey) Type() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Capability i

func (v *MockInterfaceUkey) Capability() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
