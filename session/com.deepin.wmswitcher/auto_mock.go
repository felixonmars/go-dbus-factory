// Code generated by "./generator ./session/com.deepin.wmswitcher"; DO NOT EDIT.

package wmswitcher

import "fmt"
import "github.com/godbus/dbus"
import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockWMSwitcher struct {
	MockInterfaceWmSwitcher // interface com.deepin.WMSwitcher
	proxy.MockObject
}

type MockInterfaceWmSwitcher struct {
	mock.Mock
}

// method AllowSwitch

func (v *MockInterfaceWmSwitcher) GoAllowSwitch(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWmSwitcher) AllowSwitch(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method CurrentWM

func (v *MockInterfaceWmSwitcher) GoCurrentWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWmSwitcher) CurrentWM(flags dbus.Flags) (string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method RequestSwitchWM

func (v *MockInterfaceWmSwitcher) GoRequestSwitchWM(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceWmSwitcher) RequestSwitchWM(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// signal WMChanged

func (v *MockInterfaceWmSwitcher) ConnectWMChanged(cb func(wmName string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}