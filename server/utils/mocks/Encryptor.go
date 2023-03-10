package mocks

import mock "github.com/stretchr/testify/mock"

type Encryptor struct {
	mock.Mock
}

func (_m *Encryptor) Encrypt(str string) (string, error) {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
