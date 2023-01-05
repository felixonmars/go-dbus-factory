// Code generated by "./generator ./net.hadess.sensorproxy"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package sensorproxy

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockSensorProxy struct {
	MockInterfaceSensorProxy // interface net.hadess.SensorProxy
	proxy.MockObject
}

type MockInterfaceSensorProxy struct {
	mock.Mock
}

// method ClaimAccelerometer

func (v *MockInterfaceSensorProxy) GoClaimAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSensorProxy) ClaimAccelerometer(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ReleaseAccelerometer

func (v *MockInterfaceSensorProxy) GoReleaseAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSensorProxy) ReleaseAccelerometer(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ClaimLight

func (v *MockInterfaceSensorProxy) GoClaimLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSensorProxy) ClaimLight(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method ReleaseLight

func (v *MockInterfaceSensorProxy) GoReleaseLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSensorProxy) ReleaseLight(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// property HasAccelerometer b

func (v *MockInterfaceSensorProxy) HasAccelerometer() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property AccelerometerOrientation s

func (v *MockInterfaceSensorProxy) AccelerometerOrientation() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property HasAmbientLight b

func (v *MockInterfaceSensorProxy) HasAmbientLight() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LightLevelUnit s

func (v *MockInterfaceSensorProxy) LightLevelUnit() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property LightLevel d

func (v *MockInterfaceSensorProxy) LightLevel() proxy.PropDouble {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropDouble)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
