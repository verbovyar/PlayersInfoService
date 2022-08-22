// Code generated by MockGen. DO NOT EDIT.
// Source: playersService_grpc.pb.go

// Package mock_pbGoFiles is a generated GoMock package.
package mock_pbGoFiles

import (
	context "context"
	pbGoFiles "modules/internal/infrastructure/playersInfoServiceClient/api/pbGoFiles"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockPlayersServiceClient is a mock of PlayersServiceClient interface.
type MockPlayersServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockPlayersServiceClientMockRecorder
}

// MockPlayersServiceClientMockRecorder is the mock recorder for MockPlayersServiceClient.
type MockPlayersServiceClientMockRecorder struct {
	mock *MockPlayersServiceClient
}

// NewMockPlayersServiceClient creates a new mock instance.
func NewMockPlayersServiceClient(ctrl *gomock.Controller) *MockPlayersServiceClient {
	mock := &MockPlayersServiceClient{ctrl: ctrl}
	mock.recorder = &MockPlayersServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayersServiceClient) EXPECT() *MockPlayersServiceClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPlayersServiceClient) Add(ctx context.Context, in *pbGoFiles.AddRequest, opts ...grpc.CallOption) (*pbGoFiles.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*pbGoFiles.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockPlayersServiceClientMockRecorder) Add(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPlayersServiceClient)(nil).Add), varargs...)
}

// Delete mocks base method.
func (m *MockPlayersServiceClient) Delete(ctx context.Context, in *pbGoFiles.DeleteRequest, opts ...grpc.CallOption) (*pbGoFiles.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*pbGoFiles.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockPlayersServiceClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlayersServiceClient)(nil).Delete), varargs...)
}

// List mocks base method.
func (m *MockPlayersServiceClient) List(ctx context.Context, in *pbGoFiles.ListRequest, opts ...grpc.CallOption) (*pbGoFiles.ListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*pbGoFiles.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlayersServiceClientMockRecorder) List(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlayersServiceClient)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockPlayersServiceClient) Update(ctx context.Context, in *pbGoFiles.UpdateRequest, opts ...grpc.CallOption) (*pbGoFiles.UpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*pbGoFiles.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPlayersServiceClientMockRecorder) Update(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPlayersServiceClient)(nil).Update), varargs...)
}

// MockPlayersServiceServer is a mock of PlayersServiceServer interface.
type MockPlayersServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockPlayersServiceServerMockRecorder
}

// MockPlayersServiceServerMockRecorder is the mock recorder for MockPlayersServiceServer.
type MockPlayersServiceServerMockRecorder struct {
	mock *MockPlayersServiceServer
}

// NewMockPlayersServiceServer creates a new mock instance.
func NewMockPlayersServiceServer(ctrl *gomock.Controller) *MockPlayersServiceServer {
	mock := &MockPlayersServiceServer{ctrl: ctrl}
	mock.recorder = &MockPlayersServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayersServiceServer) EXPECT() *MockPlayersServiceServerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPlayersServiceServer) Add(arg0 context.Context, arg1 *pbGoFiles.AddRequest) (*pbGoFiles.AddResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(*pbGoFiles.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockPlayersServiceServerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPlayersServiceServer)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockPlayersServiceServer) Delete(arg0 context.Context, arg1 *pbGoFiles.DeleteRequest) (*pbGoFiles.DeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*pbGoFiles.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockPlayersServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlayersServiceServer)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockPlayersServiceServer) List(arg0 context.Context, arg1 *pbGoFiles.ListRequest) (*pbGoFiles.ListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*pbGoFiles.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlayersServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlayersServiceServer)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockPlayersServiceServer) Update(arg0 context.Context, arg1 *pbGoFiles.UpdateRequest) (*pbGoFiles.UpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*pbGoFiles.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockPlayersServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPlayersServiceServer)(nil).Update), arg0, arg1)
}

// mustEmbedUnimplementedPlayersServiceServer mocks base method.
func (m *MockPlayersServiceServer) mustEmbedUnimplementedPlayersServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPlayersServiceServer")
}

// mustEmbedUnimplementedPlayersServiceServer indicates an expected call of mustEmbedUnimplementedPlayersServiceServer.
func (mr *MockPlayersServiceServerMockRecorder) mustEmbedUnimplementedPlayersServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPlayersServiceServer", reflect.TypeOf((*MockPlayersServiceServer)(nil).mustEmbedUnimplementedPlayersServiceServer))
}

// MockUnsafePlayersServiceServer is a mock of UnsafePlayersServiceServer interface.
type MockUnsafePlayersServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePlayersServiceServerMockRecorder
}

// MockUnsafePlayersServiceServerMockRecorder is the mock recorder for MockUnsafePlayersServiceServer.
type MockUnsafePlayersServiceServerMockRecorder struct {
	mock *MockUnsafePlayersServiceServer
}

// NewMockUnsafePlayersServiceServer creates a new mock instance.
func NewMockUnsafePlayersServiceServer(ctrl *gomock.Controller) *MockUnsafePlayersServiceServer {
	mock := &MockUnsafePlayersServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePlayersServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePlayersServiceServer) EXPECT() *MockUnsafePlayersServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPlayersServiceServer mocks base method.
func (m *MockUnsafePlayersServiceServer) mustEmbedUnimplementedPlayersServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPlayersServiceServer")
}

// mustEmbedUnimplementedPlayersServiceServer indicates an expected call of mustEmbedUnimplementedPlayersServiceServer.
func (mr *MockUnsafePlayersServiceServerMockRecorder) mustEmbedUnimplementedPlayersServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPlayersServiceServer", reflect.TypeOf((*MockUnsafePlayersServiceServer)(nil).mustEmbedUnimplementedPlayersServiceServer))
}