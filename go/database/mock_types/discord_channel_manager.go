// Code generated by MockGen. DO NOT EDIT.
// Source: go/database/discord_channel_manager.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ent "github.com/shifty11/dao-dao-notifier/ent"
	types "github.com/shifty11/dao-dao-notifier/types"
)

// MockIDiscordChannelManager is a mock of IDiscordChannelManager interface.
type MockIDiscordChannelManager struct {
	ctrl     *gomock.Controller
	recorder *MockIDiscordChannelManagerMockRecorder
}

// MockIDiscordChannelManagerMockRecorder is the mock recorder for MockIDiscordChannelManager.
type MockIDiscordChannelManagerMockRecorder struct {
	mock *MockIDiscordChannelManager
}

// NewMockIDiscordChannelManager creates a new mock instance.
func NewMockIDiscordChannelManager(ctrl *gomock.Controller) *MockIDiscordChannelManager {
	mock := &MockIDiscordChannelManager{ctrl: ctrl}
	mock.recorder = &MockIDiscordChannelManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDiscordChannelManager) EXPECT() *MockIDiscordChannelManagerMockRecorder {
	return m.recorder
}

// AddOrRemoveChain mocks base method.
func (m *MockIDiscordChannelManager) AddOrRemoveChain(tgChatId int64, chainId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrRemoveChain", tgChatId, chainId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrRemoveChain indicates an expected call of AddOrRemoveChain.
func (mr *MockIDiscordChannelManagerMockRecorder) AddOrRemoveChain(tgChatId, chainId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrRemoveChain", reflect.TypeOf((*MockIDiscordChannelManager)(nil).AddOrRemoveChain), tgChatId, chainId)
}

// AddOrRemoveContract mocks base method.
func (m *MockIDiscordChannelManager) AddOrRemoveContract(dChannelId int64, contractId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrRemoveContract", dChannelId, contractId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrRemoveContract indicates an expected call of AddOrRemoveContract.
func (mr *MockIDiscordChannelManagerMockRecorder) AddOrRemoveContract(dChannelId, contractId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrRemoveContract", reflect.TypeOf((*MockIDiscordChannelManager)(nil).AddOrRemoveContract), dChannelId, contractId)
}

// CountSubscriptions mocks base method.
func (m *MockIDiscordChannelManager) CountSubscriptions(channelId int64) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountSubscriptions", channelId)
	ret0, _ := ret[0].(int)
	return ret0
}

// CountSubscriptions indicates an expected call of CountSubscriptions.
func (mr *MockIDiscordChannelManagerMockRecorder) CountSubscriptions(channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSubscriptions", reflect.TypeOf((*MockIDiscordChannelManager)(nil).CountSubscriptions), channelId)
}

// CreateOrUpdateChannel mocks base method.
func (m *MockIDiscordChannelManager) CreateOrUpdateChannel(userId int64, userName string, channelId int64, name string, isGroup bool) (*ent.DiscordChannel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateChannel", userId, userName, channelId, name, isGroup)
	ret0, _ := ret[0].(*ent.DiscordChannel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CreateOrUpdateChannel indicates an expected call of CreateOrUpdateChannel.
func (mr *MockIDiscordChannelManagerMockRecorder) CreateOrUpdateChannel(userId, userName, channelId, name, isGroup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateChannel", reflect.TypeOf((*MockIDiscordChannelManager)(nil).CreateOrUpdateChannel), userId, userName, channelId, name, isGroup)
}

// Delete mocks base method.
func (m *MockIDiscordChannelManager) Delete(userId, channelId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, channelId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIDiscordChannelManagerMockRecorder) Delete(userId, channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIDiscordChannelManager)(nil).Delete), userId, channelId)
}

// DeleteMultiple mocks base method.
func (m *MockIDiscordChannelManager) DeleteMultiple(channelIds []int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteMultiple", channelIds)
}

// DeleteMultiple indicates an expected call of DeleteMultiple.
func (mr *MockIDiscordChannelManagerMockRecorder) DeleteMultiple(channelIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMultiple", reflect.TypeOf((*MockIDiscordChannelManager)(nil).DeleteMultiple), channelIds)
}

// GetChannelUsers mocks base method.
func (m *MockIDiscordChannelManager) GetChannelUsers(channelId int64) []*ent.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelUsers", channelId)
	ret0, _ := ret[0].([]*ent.User)
	return ret0
}

// GetChannelUsers indicates an expected call of GetChannelUsers.
func (mr *MockIDiscordChannelManagerMockRecorder) GetChannelUsers(channelId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelUsers", reflect.TypeOf((*MockIDiscordChannelManager)(nil).GetChannelUsers), channelId)
}

// GetSubscribedIds mocks base method.
func (m *MockIDiscordChannelManager) GetSubscribedIds(query *ent.DiscordChannelQuery) []types.DiscordChannelQueryResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribedIds", query)
	ret0, _ := ret[0].([]types.DiscordChannelQueryResult)
	return ret0
}

// GetSubscribedIds indicates an expected call of GetSubscribedIds.
func (mr *MockIDiscordChannelManagerMockRecorder) GetSubscribedIds(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribedIds", reflect.TypeOf((*MockIDiscordChannelManager)(nil).GetSubscribedIds), query)
}
