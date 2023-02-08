// Code generated by "./generator ./system/org.deepin.dde.authenticate1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package authenticate1

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockAuthenticate struct {
	MockInterfaceAuthenticate // interface org.deepin.dde.Authenticate1
	proxy.MockObject
}

type MockInterfaceAuthenticate struct {
	mock.Mock
}

// method Authenticate

func (v *MockInterfaceAuthenticate) GoAuthenticate(flags dbus.Flags, ch chan *dbus.Call, username string, authFlags int32, appType int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, authFlags, appType)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) Authenticate(flags dbus.Flags, username string, authFlags int32, appType int32) (string, error) {
	mockArgs := v.Called(flags, username, authFlags, appType)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetLimits

func (v *MockInterfaceAuthenticate) GoGetLimits(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) GetLimits(flags dbus.Flags, username string) (string, error) {
	mockArgs := v.Called(flags, username)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method PreOneKeyLogin

func (v *MockInterfaceAuthenticate) GoPreOneKeyLogin(flags dbus.Flags, ch chan *dbus.Call, flag int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, flag)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) PreOneKeyLogin(flags dbus.Flags, flag int32) (string, error) {
	mockArgs := v.Called(flags, flag)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method ResetLimits

func (v *MockInterfaceAuthenticate) GoResetLimits(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAuthenticate) ResetLimits(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// signal LimitUpdated

func (v *MockInterfaceAuthenticate) ConnectLimitUpdated(cb func(username string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property SupportEncrypts s

func (v *MockInterfaceAuthenticate) SupportEncrypts() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property FrameworkState i

func (v *MockInterfaceAuthenticate) FrameworkState() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SupportedFlags i

func (v *MockInterfaceAuthenticate) SupportedFlags() proxy.PropInt32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockFingerprint struct {
	MockInterfaceFingerprint // interface org.deepin.dde.Authenticate1.Fingerprint
	proxy.MockObject
}

type MockInterfaceFingerprint struct {
	mock.Mock
}

// method Claim

func (v *MockInterfaceFingerprint) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string, claimed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, claimed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Claim(flags dbus.Flags, username string, claimed bool) error {
	mockArgs := v.Called(flags, username, claimed)

	return mockArgs.Error(0)
}

// method DeleteAllFingers

func (v *MockInterfaceFingerprint) GoDeleteAllFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) DeleteAllFingers(flags dbus.Flags, username string) error {
	mockArgs := v.Called(flags, username)

	return mockArgs.Error(0)
}

// method DeleteFinger

func (v *MockInterfaceFingerprint) GoDeleteFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) DeleteFinger(flags dbus.Flags, username string, finger string) error {
	mockArgs := v.Called(flags, username, finger)

	return mockArgs.Error(0)
}

// method Enroll

func (v *MockInterfaceFingerprint) GoEnroll(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Enroll(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// method ListFingers

func (v *MockInterfaceFingerprint) GoListFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) ListFingers(flags dbus.Flags, username string) ([]string, error) {
	mockArgs := v.Called(flags, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetDefaultDevice

func (v *MockInterfaceFingerprint) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) SetDefaultDevice(flags dbus.Flags, device string) error {
	mockArgs := v.Called(flags, device)

	return mockArgs.Error(0)
}

// method StopEnroll

func (v *MockInterfaceFingerprint) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) StopEnroll(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceFingerprint) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) StopVerify(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Verify

func (v *MockInterfaceFingerprint) GoVerify(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	mockArgs := v.Called(flags, ch, finger)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFingerprint) Verify(flags dbus.Flags, finger string) error {
	mockArgs := v.Called(flags, finger)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceFingerprint) ConnectEnrollStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceFingerprint) ConnectVerifyStatus(cb func(id string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal Touch

func (v *MockInterfaceFingerprint) ConnectTouch(cb func(id string, pressed bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property DefaultDevice s

func (v *MockInterfaceFingerprint) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Devices s

func (v *MockInterfaceFingerprint) Devices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockUKey struct {
	MockInterfaceUkey // interface org.deepin.dde.Authenticate1.UKey
	proxy.MockObject
}

type MockInterfaceUkey struct {
	mock.Mock
}

// method ConstructVerification

func (v *MockInterfaceUkey) GoConstructVerification(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, useDefaultService bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, useDefaultService)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) ConstructVerification(flags dbus.Flags, serviceName string, username string, useDefaultService bool) (string, error) {
	mockArgs := v.Called(flags, serviceName, username, useDefaultService)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetPINLength

func (v *MockInterfaceUkey) GoGetPINLength(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, useDefaultDevice bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, useDefaultDevice)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetPINLength(flags dbus.Flags, serviceName string, username string, useDefaultDevice bool) (int32, error) {
	mockArgs := v.Called(flags, serviceName, username, useDefaultDevice)

	ret0, ok := mockArgs.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method GetUserList

func (v *MockInterfaceUkey) GoGetUserList(flags dbus.Flags, ch chan *dbus.Call, serviceName string, useDefaultDevice bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, useDefaultDevice)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) GetUserList(flags dbus.Flags, serviceName string, useDefaultDevice bool) ([]string, error) {
	mockArgs := v.Called(flags, serviceName, useDefaultDevice)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method SetDefaultDevice

func (v *MockInterfaceUkey) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetDefaultDevice(flags dbus.Flags, device string) error {
	mockArgs := v.Called(flags, device)

	return mockArgs.Error(0)
}

// method SetPin

func (v *MockInterfaceUkey) GoSetPin(flags dbus.Flags, ch chan *dbus.Call, id string, pin string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id, pin)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetPin(flags dbus.Flags, id string, pin string) error {
	mockArgs := v.Called(flags, id, pin)

	return mockArgs.Error(0)
}

// method SetSessionPath

func (v *MockInterfaceUkey) GoSetSessionPath(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) SetSessionPath(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method StartVerify

func (v *MockInterfaceUkey) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) StartVerify(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceUkey) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUkey) StopVerify(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

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

// signal State

func (v *MockInterfaceUkey) ConnectState(cb func(id string, state int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property ValidDevices s

func (v *MockInterfaceUkey) ValidDevices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultDevice s

func (v *MockInterfaceUkey) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockFace struct {
	MockInterfaceFace // interface org.deepin.dde.Authenticate1.Face
	proxy.MockObject
}

type MockInterfaceFace struct {
	mock.Mock
}

// method DeleteFace

func (v *MockInterfaceFace) GoDeleteFace(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, faceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, faceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) DeleteFace(flags dbus.Flags, serviceName string, username string, faceName string) error {
	mockArgs := v.Called(flags, serviceName, username, faceName)

	return mockArgs.Error(0)
}

// method DeleteFaces

func (v *MockInterfaceFace) GoDeleteFaces(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) DeleteFaces(flags dbus.Flags, serviceName string, username string) error {
	mockArgs := v.Called(flags, serviceName, username)

	return mockArgs.Error(0)
}

// method GetShareMemInfo

func (v *MockInterfaceFace) GoGetShareMemInfo(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) GetShareMemInfo(flags dbus.Flags, id string) (string, string, int32, error) {
	mockArgs := v.Called(flags, id)

	ret2, ok := mockArgs.Get(2).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 2, mockArgs.Get(2)))
	}

	return mockArgs.String(0), mockArgs.String(1), ret2, mockArgs.Error(3)
}

// method ListFaces

func (v *MockInterfaceFace) GoListFaces(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) ListFaces(flags dbus.Flags, serviceName string, username string) ([]string, error) {
	mockArgs := v.Called(flags, serviceName, username)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RenameFace

func (v *MockInterfaceFace) GoRenameFace(flags dbus.Flags, ch chan *dbus.Call, serviceName string, username string, oldFace string, newFace string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, username, oldFace, newFace)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) RenameFace(flags dbus.Flags, serviceName string, username string, oldFace string, newFace string) error {
	mockArgs := v.Called(flags, serviceName, username, oldFace, newFace)

	return mockArgs.Error(0)
}

// method SetDefaultDevice

func (v *MockInterfaceFace) GoSetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call, serviceName string, device string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName, device)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) SetDefaultDevice(flags dbus.Flags, serviceName string, device string) error {
	mockArgs := v.Called(flags, serviceName, device)

	return mockArgs.Error(0)
}

// method SetDefaultService

func (v *MockInterfaceFace) GoSetDefaultService(flags dbus.Flags, ch chan *dbus.Call, serviceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, serviceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) SetDefaultService(flags dbus.Flags, serviceName string) error {
	mockArgs := v.Called(flags, serviceName)

	return mockArgs.Error(0)
}

// method StartEnroll

func (v *MockInterfaceFace) GoStartEnroll(flags dbus.Flags, ch chan *dbus.Call, username string, serviceName string, faceName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, serviceName, faceName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StartEnroll(flags dbus.Flags, username string, serviceName string, faceName string) (string, error) {
	mockArgs := v.Called(flags, username, serviceName, faceName)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method StartVerify

func (v *MockInterfaceFace) GoStartVerify(flags dbus.Flags, ch chan *dbus.Call, username string, serviceName string, timeout int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, username, serviceName, timeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StartVerify(flags dbus.Flags, username string, serviceName string, timeout int32) (string, error) {
	mockArgs := v.Called(flags, username, serviceName, timeout)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method StopEnroll

func (v *MockInterfaceFace) GoStopEnroll(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StopEnroll(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method StopVerify

func (v *MockInterfaceFace) GoStopVerify(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceFace) StopVerify(flags dbus.Flags, id string) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceFace) ConnectEnrollStatus(cb func(id string, user string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal VerifyStatus

func (v *MockInterfaceFace) ConnectVerifyStatus(cb func(id string, user string, code int32, msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DeviceStatus

func (v *MockInterfaceFace) ConnectDeviceStatus(cb func(serviceName string, code int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property DefaultDevice s

func (v *MockInterfaceFace) DefaultDevice() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property DefaultService s

func (v *MockInterfaceFace) DefaultService() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ServiceList s

func (v *MockInterfaceFace) ServiceList() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockCharaManger struct {
	MockInterfaceCharaManger // interface org.deepin.dde.Authenticate1.CharaManger
	proxy.MockObject
}

type MockInterfaceCharaManger struct {
	mock.Mock
}

// method Delete

func (v *MockInterfaceCharaManger) GoDelete(flags dbus.Flags, ch chan *dbus.Call, charaType int32, charaName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, charaType, charaName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCharaManger) Delete(flags dbus.Flags, charaType int32, charaName string) error {
	mockArgs := v.Called(flags, charaType, charaName)

	return mockArgs.Error(0)
}

// method EnrollStart

func (v *MockInterfaceCharaManger) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, driverName string, charaType int32, charaName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, driverName, charaType, charaName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCharaManger) EnrollStart(flags dbus.Flags, driverName string, charaType int32, charaName string) (dbus.UnixFD, error) {
	mockArgs := v.Called(flags, driverName, charaType, charaName)

	ret0, ok := mockArgs.Get(0).(dbus.UnixFD)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method EnrollStop

func (v *MockInterfaceCharaManger) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCharaManger) EnrollStop(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method List

func (v *MockInterfaceCharaManger) GoList(flags dbus.Flags, ch chan *dbus.Call, driverName string, charaType int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, driverName, charaType)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCharaManger) List(flags dbus.Flags, driverName string, charaType int32) (string, error) {
	mockArgs := v.Called(flags, driverName, charaType)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Rename

func (v *MockInterfaceCharaManger) GoRename(flags dbus.Flags, ch chan *dbus.Call, charaType int32, oldName string, newName string) *dbus.Call {
	mockArgs := v.Called(flags, ch, charaType, oldName, newName)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceCharaManger) Rename(flags dbus.Flags, charaType int32, oldName string, newName string) error {
	mockArgs := v.Called(flags, charaType, oldName, newName)

	return mockArgs.Error(0)
}

// signal EnrollStatus

func (v *MockInterfaceCharaManger) ConnectEnrollStatus(cb func(Sender string, Code int32, Msg string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal CharaUpdated

func (v *MockInterfaceCharaManger) ConnectCharaUpdated(cb func(DriverName string, CharaType int32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DriverChanged

func (v *MockInterfaceCharaManger) ConnectDriverChanged(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property DriverInfo s

func (v *MockInterfaceCharaManger) DriverInfo() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
