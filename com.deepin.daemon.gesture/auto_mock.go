// Code generated by "./generator ./com.deepin.daemon.gesture"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package gesture

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockGesture struct {
	MockInterfaceGesture // interface com.deepin.daemon.Gesture
	proxy.MockObject
}

type MockInterfaceGesture struct {
	mock.Mock
}

// method SetShortPressDuration

func (v *MockInterfaceGesture) GoSetShortPressDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, duration)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGesture) SetShortPressDuration(flags dbus.Flags, duration uint32) error {
	mockArgs := v.Called(flags, duration)

	return mockArgs.Error(0)
}

// method SetEdgeMoveStopDuration

func (v *MockInterfaceGesture) GoSetEdgeMoveStopDuration(flags dbus.Flags, ch chan *dbus.Call, duration uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, duration)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGesture) SetEdgeMoveStopDuration(flags dbus.Flags, duration uint32) error {
	mockArgs := v.Called(flags, duration)

	return mockArgs.Error(0)
}

// method SetInputIgnore

func (v *MockInterfaceGesture) GoSetInputIgnore(flags dbus.Flags, ch chan *dbus.Call, node string, ignore bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, node, ignore)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceGesture) SetInputIgnore(flags dbus.Flags, node string, ignore bool) error {
	mockArgs := v.Called(flags, node, ignore)

	return mockArgs.Error(0)
}

// signal Event

func (v *MockInterfaceGesture) ConnectEvent(cb func(name string, direction string, fingers int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DbclickDown

func (v *MockInterfaceGesture) ConnectDbclickDown(cb func(fingers int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal SwipeMoving

func (v *MockInterfaceGesture) ConnectSwipeMoving(cb func(fingers int32, accelX float64, accely float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal SwipeStop

func (v *MockInterfaceGesture) ConnectSwipeStop(cb func(fingers int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchEdgeEvent

func (v *MockInterfaceGesture) ConnectTouchEdgeEvent(cb func(direction string, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchSinglePressTimeout

func (v *MockInterfaceGesture) ConnectTouchSinglePressTimeout(cb func(time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchPressTimeout

func (v *MockInterfaceGesture) ConnectTouchPressTimeout(cb func(fingers int32, time int32, scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchUpOrCancel

func (v *MockInterfaceGesture) ConnectTouchUpOrCancel(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchEdgeMoveStop

func (v *MockInterfaceGesture) ConnectTouchEdgeMoveStop(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchEdgeMoveStopLeave

func (v *MockInterfaceGesture) ConnectTouchEdgeMoveStopLeave(cb func(direction string, scalex float64, scaley float64, duration int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchMoving

func (v *MockInterfaceGesture) ConnectTouchMoving(cb func(scalex float64, scaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal TouchMovementEvent

func (v *MockInterfaceGesture) ConnectTouchMovementEvent(cb func(duration string, fingers int32, startScalex float64, startScaley float64, endScalex float64, endScaley float64)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
