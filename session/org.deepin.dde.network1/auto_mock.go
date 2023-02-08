// Code generated by "./generator ./session/org.deepin.dde.network1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package network1

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockNetwork struct {
	MockInterfaceNetwork // interface org.deepin.dde.Network1
	proxy.MockObject
}

type MockInterfaceNetwork struct {
	mock.Mock
}

// method ActivateAccessPoint

func (v *MockInterfaceNetwork) GoActivateAccessPoint(flags dbus.Flags, ch chan *dbus.Call, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, apPath, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) ActivateAccessPoint(flags dbus.Flags, uuid string, apPath dbus.ObjectPath, devPath dbus.ObjectPath) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, uuid, apPath, devPath)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ActivateConnection

func (v *MockInterfaceNetwork) GoActivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) ActivateConnection(flags dbus.Flags, uuid string, devPath dbus.ObjectPath) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, uuid, devPath)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method DeactivateConnection

func (v *MockInterfaceNetwork) GoDeactivateConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) DeactivateConnection(flags dbus.Flags, uuid string) error {
	mockArgs := v.Called(flags, uuid)

	return mockArgs.Error(0)
}

// method DeleteConnection

func (v *MockInterfaceNetwork) GoDeleteConnection(flags dbus.Flags, ch chan *dbus.Call, uuid string) *dbus.Call {
	mockArgs := v.Called(flags, ch, uuid)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) DeleteConnection(flags dbus.Flags, uuid string) error {
	mockArgs := v.Called(flags, uuid)

	return mockArgs.Error(0)
}

// method DisableWirelessHotspotMode

func (v *MockInterfaceNetwork) GoDisableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) DisableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method DisconnectDevice

func (v *MockInterfaceNetwork) GoDisconnectDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) DisconnectDevice(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method EnableDevice

func (v *MockInterfaceNetwork) GoEnableDevice(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath, enabled bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath, enabled)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) EnableDevice(flags dbus.Flags, devPath dbus.ObjectPath, enabled bool) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, devPath, enabled)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method EnableWirelessHotspotMode

func (v *MockInterfaceNetwork) GoEnableWirelessHotspotMode(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) EnableWirelessHotspotMode(flags dbus.Flags, devPath dbus.ObjectPath) error {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Error(0)
}

// method GetAccessPoints

func (v *MockInterfaceNetwork) GoGetAccessPoints(flags dbus.Flags, ch chan *dbus.Call, path dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, path)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetAccessPoints(flags dbus.Flags, path dbus.ObjectPath) (string, error) {
	mockArgs := v.Called(flags, path)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetActiveConnectionInfo

func (v *MockInterfaceNetwork) GoGetActiveConnectionInfo(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetActiveConnectionInfo(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetAutoProxy

func (v *MockInterfaceNetwork) GoGetAutoProxy(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetAutoProxy(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetProxy

func (v *MockInterfaceNetwork) GoGetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyType)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetProxy(flags dbus.Flags, proxyType string) (string, string, error) {
	mockArgs := v.Called(flags, proxyType)

	return mockArgs.String(0), mockArgs.String(1), mockArgs.Error(2)
}

// method GetProxyIgnoreHosts

func (v *MockInterfaceNetwork) GoGetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetProxyIgnoreHosts(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetProxyMethod

func (v *MockInterfaceNetwork) GoGetProxyMethod(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetProxyMethod(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method GetSupportedConnectionTypes

func (v *MockInterfaceNetwork) GoGetSupportedConnectionTypes(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) GetSupportedConnectionTypes(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method IsDeviceEnabled

func (v *MockInterfaceNetwork) GoIsDeviceEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) IsDeviceEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (bool, error) {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method IsWirelessHotspotModeEnabled

func (v *MockInterfaceNetwork) GoIsWirelessHotspotModeEnabled(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) IsWirelessHotspotModeEnabled(flags dbus.Flags, devPath dbus.ObjectPath) (bool, error) {
	mockArgs := v.Called(flags, devPath)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method ListDeviceConnections

func (v *MockInterfaceNetwork) GoListDeviceConnections(flags dbus.Flags, ch chan *dbus.Call, devPath dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPath)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) ListDeviceConnections(flags dbus.Flags, devPath dbus.ObjectPath) ([]dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, devPath)

	ret0, ok := mockArgs.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RequestWirelessScan

func (v *MockInterfaceNetwork) GoRequestWirelessScan(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) RequestWirelessScan(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method SetAutoProxy

func (v *MockInterfaceNetwork) GoSetAutoProxy(flags dbus.Flags, ch chan *dbus.Call, proxyAuto string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyAuto)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) SetAutoProxy(flags dbus.Flags, proxyAuto string) error {
	mockArgs := v.Called(flags, proxyAuto)

	return mockArgs.Error(0)
}

// method SetDeviceManaged

func (v *MockInterfaceNetwork) GoSetDeviceManaged(flags dbus.Flags, ch chan *dbus.Call, devPathOrIfc string, managed bool) *dbus.Call {
	mockArgs := v.Called(flags, ch, devPathOrIfc, managed)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) SetDeviceManaged(flags dbus.Flags, devPathOrIfc string, managed bool) error {
	mockArgs := v.Called(flags, devPathOrIfc, managed)

	return mockArgs.Error(0)
}

// method SetProxy

func (v *MockInterfaceNetwork) GoSetProxy(flags dbus.Flags, ch chan *dbus.Call, proxyType string, host string, port string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyType, host, port)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) SetProxy(flags dbus.Flags, proxyType string, host string, port string) error {
	mockArgs := v.Called(flags, proxyType, host, port)

	return mockArgs.Error(0)
}

// method SetProxyIgnoreHosts

func (v *MockInterfaceNetwork) GoSetProxyIgnoreHosts(flags dbus.Flags, ch chan *dbus.Call, ignoreHosts string) *dbus.Call {
	mockArgs := v.Called(flags, ch, ignoreHosts)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) SetProxyIgnoreHosts(flags dbus.Flags, ignoreHosts string) error {
	mockArgs := v.Called(flags, ignoreHosts)

	return mockArgs.Error(0)
}

// method SetProxyMethod

func (v *MockInterfaceNetwork) GoSetProxyMethod(flags dbus.Flags, ch chan *dbus.Call, proxyMode string) *dbus.Call {
	mockArgs := v.Called(flags, ch, proxyMode)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNetwork) SetProxyMethod(flags dbus.Flags, proxyMode string) error {
	mockArgs := v.Called(flags, proxyMode)

	return mockArgs.Error(0)
}

// signal AccessPointAdded

func (v *MockInterfaceNetwork) ConnectAccessPointAdded(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AccessPointRemoved

func (v *MockInterfaceNetwork) ConnectAccessPointRemoved(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal AccessPointPropertiesChanged

func (v *MockInterfaceNetwork) ConnectAccessPointPropertiesChanged(cb func(devPath string, apJSON string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal DeviceEnabled

func (v *MockInterfaceNetwork) ConnectDeviceEnabled(cb func(devPath string, enabled bool)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// property Connectivity u

func (v *MockInterfaceNetwork) Connectivity() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property NetworkingEnabled b

func (v *MockInterfaceNetwork) NetworkingEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property VpnEnabled b

func (v *MockInterfaceNetwork) VpnEnabled() proxy.PropBool {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropBool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Devices s

func (v *MockInterfaceNetwork) Devices() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Connections s

func (v *MockInterfaceNetwork) Connections() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ActiveConnections s

func (v *MockInterfaceNetwork) ActiveConnections() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property State u

func (v *MockInterfaceNetwork) State() proxy.PropUint32 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
