// Code generated by "./generator ./session/org.bluez.obex"; DO NOT EDIT.

package obex

import (
	"fmt"

	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-dbus-factory/object_manager"
	"github.com/linuxdeepin/go-lib/dbusutil/proxy"
	"github.com/stretchr/testify/mock"
)

type MockObjectManager struct {
	object_manager.MockInterfaceObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.MockObject
}

type MockManager struct {
	MockInterfaceAgentManager // interface org.bluez.obex.AgentManager1
	MockInterfaceClient       // interface org.bluez.obex.Client1
	proxy.MockObject
}

type MockInterfaceAgentManager struct {
	mock.Mock
}

// method RegisterAgent

func (v *MockInterfaceAgentManager) GoRegisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, agent)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAgentManager) RegisterAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	mockArgs := v.Called(flags, agent)

	return mockArgs.Error(0)
}

// method UnregisterAgent

func (v *MockInterfaceAgentManager) GoUnregisterAgent(flags dbus.Flags, ch chan *dbus.Call, agent dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, agent)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceAgentManager) UnregisterAgent(flags dbus.Flags, agent dbus.ObjectPath) error {
	mockArgs := v.Called(flags, agent)

	return mockArgs.Error(0)
}

type MockInterfaceClient struct {
	mock.Mock
}

// method CreateSession

func (v *MockInterfaceClient) GoCreateSession(flags dbus.Flags, ch chan *dbus.Call, destination string, args map[string]dbus.Variant) *dbus.Call {
	mockArgs := v.Called(flags, ch, destination, args)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceClient) CreateSession(flags dbus.Flags, destination string, args map[string]dbus.Variant) (dbus.ObjectPath, error) {
	mockArgs := v.Called(flags, destination, args)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method RemoveSession

func (v *MockInterfaceClient) GoRemoveSession(flags dbus.Flags, ch chan *dbus.Call, session dbus.ObjectPath) *dbus.Call {
	mockArgs := v.Called(flags, ch, session)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceClient) RemoveSession(flags dbus.Flags, session dbus.ObjectPath) error {
	mockArgs := v.Called(flags, session)

	return mockArgs.Error(0)
}

type MockSession struct {
	MockInterfaceSession    // interface org.bluez.obex.Session1
	MockInterfaceObjectPush // interface org.bluez.obex.ObjectPush1
	proxy.MockObject
}

type MockInterfaceSession struct {
	mock.Mock
}

// method GetCapabilities

func (v *MockInterfaceSession) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceSession) GetCapabilities(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// property Source s

func (v *MockInterfaceSession) Source() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Destination s

func (v *MockInterfaceSession) Destination() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Channel y

func (v *MockInterfaceSession) Channel() proxy.PropByte {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropByte)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Target s

func (v *MockInterfaceSession) Target() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

type MockInterfaceObjectPush struct {
	mock.Mock
}

// method SendFile

func (v *MockInterfaceObjectPush) GoSendFile(flags dbus.Flags, ch chan *dbus.Call, sourcefile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, sourcefile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceObjectPush) SendFile(flags dbus.Flags, sourcefile string) (dbus.ObjectPath, map[string]dbus.Variant, error) {
	mockArgs := v.Called(flags, sourcefile)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(map[string]dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	return ret0, ret1, mockArgs.Error(2)
}

// method PullBusinessCard

func (v *MockInterfaceObjectPush) GoPullBusinessCard(flags dbus.Flags, ch chan *dbus.Call, targetfile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, targetfile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceObjectPush) PullBusinessCard(flags dbus.Flags, targetfile string) (dbus.ObjectPath, map[string]dbus.Variant, error) {
	mockArgs := v.Called(flags, targetfile)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(map[string]dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	return ret0, ret1, mockArgs.Error(2)
}

// method ExchangeBusinessCards

func (v *MockInterfaceObjectPush) GoExchangeBusinessCards(flags dbus.Flags, ch chan *dbus.Call, clientfile string, targetfile string) *dbus.Call {
	mockArgs := v.Called(flags, ch, clientfile, targetfile)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceObjectPush) ExchangeBusinessCards(flags dbus.Flags, clientfile string, targetfile string) (dbus.ObjectPath, map[string]dbus.Variant, error) {
	mockArgs := v.Called(flags, clientfile, targetfile)

	ret0, ok := mockArgs.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	ret1, ok := mockArgs.Get(1).(map[string]dbus.Variant)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 1, mockArgs.Get(1)))
	}

	return ret0, ret1, mockArgs.Error(2)
}

type MockTransfer struct {
	MockInterfaceTransfer // interface org.bluez.obex.Transfer1
	proxy.MockObject
}

type MockInterfaceTransfer struct {
	mock.Mock
}

// method Cancel

func (v *MockInterfaceTransfer) GoCancel(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceTransfer) Cancel(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// property Status s

func (v *MockInterfaceTransfer) Status() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Session o

func (v *MockInterfaceTransfer) Session() proxy.PropObjectPath {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Name s

func (v *MockInterfaceTransfer) Name() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Type s

func (v *MockInterfaceTransfer) Type() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Size t

func (v *MockInterfaceTransfer) Size() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Time t

func (v *MockInterfaceTransfer) Time() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Filename s

func (v *MockInterfaceTransfer) Filename() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Transferred t

func (v *MockInterfaceTransfer) Transferred() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}