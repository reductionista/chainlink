// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	query "github.com/cosmos/cosmos-sdk/types/query"
	mock "github.com/stretchr/testify/mock"

	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

	tx "github.com/cosmos/cosmos-sdk/types/tx"

	types "github.com/cosmos/cosmos-sdk/types"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Account provides a mock function with given fields: address
func (_m *Reader) Account(address types.AccAddress) (uint64, uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(types.AccAddress) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(types.AccAddress) uint64); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(types.AccAddress) error); ok {
		r2 = rf(address)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Balance provides a mock function with given fields: addr, denom
func (_m *Reader) Balance(addr types.AccAddress, denom string) (*types.Coin, error) {
	ret := _m.Called(addr, denom)

	var r0 *types.Coin
	if rf, ok := ret.Get(0).(func(types.AccAddress, string) *types.Coin); ok {
		r0 = rf(addr, denom)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Coin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.AccAddress, string) error); ok {
		r1 = rf(addr, denom)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHeight provides a mock function with given fields: height
func (_m *Reader) BlockByHeight(height int64) (*tmservice.GetBlockByHeightResponse, error) {
	ret := _m.Called(height)

	var r0 *tmservice.GetBlockByHeightResponse
	if rf, ok := ret.Get(0).(func(int64) *tmservice.GetBlockByHeightResponse); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tmservice.GetBlockByHeightResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractStore provides a mock function with given fields: contractAddress, queryMsg
func (_m *Reader) ContractStore(contractAddress types.AccAddress, queryMsg []byte) ([]byte, error) {
	ret := _m.Called(contractAddress, queryMsg)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(types.AccAddress, []byte) []byte); ok {
		r0 = rf(contractAddress, queryMsg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.AccAddress, []byte) error); ok {
		r1 = rf(contractAddress, queryMsg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestBlock provides a mock function with given fields:
func (_m *Reader) LatestBlock() (*tmservice.GetLatestBlockResponse, error) {
	ret := _m.Called()

	var r0 *tmservice.GetLatestBlockResponse
	if rf, ok := ret.Get(0).(func() *tmservice.GetLatestBlockResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tmservice.GetLatestBlockResponse)
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

// Tx provides a mock function with given fields: hash
func (_m *Reader) Tx(hash string) (*tx.GetTxResponse, error) {
	ret := _m.Called(hash)

	var r0 *tx.GetTxResponse
	if rf, ok := ret.Get(0).(func(string) *tx.GetTxResponse); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tx.GetTxResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxsEvents provides a mock function with given fields: events, paginationParams
func (_m *Reader) TxsEvents(events []string, paginationParams *query.PageRequest) (*tx.GetTxsEventResponse, error) {
	ret := _m.Called(events, paginationParams)

	var r0 *tx.GetTxsEventResponse
	if rf, ok := ret.Get(0).(func([]string, *query.PageRequest) *tx.GetTxsEventResponse); ok {
		r0 = rf(events, paginationParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tx.GetTxsEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, *query.PageRequest) error); ok {
		r1 = rf(events, paginationParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
