// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import big "math/big"
import bson "gopkg.in/mgo.v2/bson"
import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/tomochain/backend-matching-engine/types"

// AccountDao is an autogenerated mock type for the AccountDao type
type AccountDao struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountDao) Create(account *types.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Drop provides a mock function with given fields:
func (_m *AccountDao) Drop() {
	_m.Called()
}

// GetAll provides a mock function with given fields:
func (_m *AccountDao) GetAll() ([]types.Account, error) {
	ret := _m.Called()

	var r0 []types.Account
	if rf, ok := ret.Get(0).(func() []types.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAddress provides a mock function with given fields: owner
func (_m *AccountDao) GetByAddress(owner common.Address) (*types.Account, error) {
	ret := _m.Called(owner)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(common.Address) *types.Account); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *AccountDao) GetByID(id bson.ObjectId) (*types.Account, error) {
	ret := _m.Called(id)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalance provides a mock function with given fields: owner, token
func (_m *AccountDao) GetTokenBalance(owner common.Address, token common.Address) (*types.TokenBalance, error) {
	ret := _m.Called(owner, token)

	var r0 *types.TokenBalance
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *types.TokenBalance); ok {
		r0 = rf(owner, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(owner, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalances provides a mock function with given fields: owner
func (_m *AccountDao) GetTokenBalances(owner common.Address) (map[common.Address]*types.TokenBalance, error) {
	ret := _m.Called(owner)

	var r0 map[common.Address]*types.TokenBalance
	if rf, ok := ret.Get(0).(func(common.Address) map[common.Address]*types.TokenBalance); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Address]*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAllowance provides a mock function with given fields: owner, token, allowance
func (_m *AccountDao) UpdateAllowance(owner common.Address, token common.Address, allowance *big.Int) error {
	ret := _m.Called(owner, token, allowance)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, common.Address, *big.Int) error); ok {
		r0 = rf(owner, token, allowance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBalance provides a mock function with given fields: owner, token, balance
func (_m *AccountDao) UpdateBalance(owner common.Address, token common.Address, balance *big.Int) error {
	ret := _m.Called(owner, token, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, common.Address, *big.Int) error); ok {
		r0 = rf(owner, token, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTokenBalance provides a mock function with given fields: owner, token, tokenBalance
func (_m *AccountDao) UpdateTokenBalance(owner common.Address, token common.Address, tokenBalance *types.TokenBalance) error {
	ret := _m.Called(owner, token, tokenBalance)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, common.Address, *types.TokenBalance) error); ok {
		r0 = rf(owner, token, tokenBalance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}