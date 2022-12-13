// Code generated by "./generator ./system/org.deepin.dde.systeminfo1"; DO NOT EDIT.

package systeminfo1

import "fmt"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockSystemInfo struct {
	MockInterfaceSystemInfo // interface org.deepin.dde.SystemInfo1
	proxy.MockObject
}

type MockInterfaceSystemInfo struct {
	mock.Mock
}

// property MemorySizeHuman s

func (v *MockInterfaceSystemInfo) MemorySizeHuman() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property MemorySize t

func (v *MockInterfaceSystemInfo) MemorySize() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property CurrentSpeed t

func (v *MockInterfaceSystemInfo) CurrentSpeed() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}