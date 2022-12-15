// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tigrisdata/tigris-client-go/api/server/v1 (interfaces: TigrisServer,AuthServer,ManagementServer,ObservabilityServer)

// Package api is a generated GoMock package.
package api

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/tigrisdata/tigris-client-go/api/server/v1"
)

// MockTigrisServer is a mock of TigrisServer interface.
type MockTigrisServer struct {
	ctrl     *gomock.Controller
	recorder *MockTigrisServerMockRecorder
}

// MockTigrisServerMockRecorder is the mock recorder for MockTigrisServer.
type MockTigrisServerMockRecorder struct {
	mock *MockTigrisServer
}

// NewMockTigrisServer creates a new mock instance.
func NewMockTigrisServer(ctrl *gomock.Controller) *MockTigrisServer {
	mock := &MockTigrisServer{ctrl: ctrl}
	mock.recorder = &MockTigrisServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTigrisServer) EXPECT() *MockTigrisServerMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockTigrisServer) BeginTransaction(arg0 context.Context, arg1 *api.BeginTransactionRequest) (*api.BeginTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", arg0, arg1)
	ret0, _ := ret[0].(*api.BeginTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockTigrisServerMockRecorder) BeginTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockTigrisServer)(nil).BeginTransaction), arg0, arg1)
}

// CommitTransaction mocks base method.
func (m *MockTigrisServer) CommitTransaction(arg0 context.Context, arg1 *api.CommitTransactionRequest) (*api.CommitTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction", arg0, arg1)
	ret0, _ := ret[0].(*api.CommitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockTigrisServerMockRecorder) CommitTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockTigrisServer)(nil).CommitTransaction), arg0, arg1)
}

// CreateAppKey mocks base method.
func (m *MockTigrisServer) CreateAppKey(arg0 context.Context, arg1 *api.CreateAppKeyRequest) (*api.CreateAppKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppKey", arg0, arg1)
	ret0, _ := ret[0].(*api.CreateAppKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAppKey indicates an expected call of CreateAppKey.
func (mr *MockTigrisServerMockRecorder) CreateAppKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppKey", reflect.TypeOf((*MockTigrisServer)(nil).CreateAppKey), arg0, arg1)
}

// CreateBranch mocks base method.
func (m *MockTigrisServer) CreateBranch(arg0 context.Context, arg1 *api.CreateBranchRequest) (*api.CreateBranchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBranch", arg0, arg1)
	ret0, _ := ret[0].(*api.CreateBranchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBranch indicates an expected call of CreateBranch.
func (mr *MockTigrisServerMockRecorder) CreateBranch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBranch", reflect.TypeOf((*MockTigrisServer)(nil).CreateBranch), arg0, arg1)
}

// CreateOrUpdateCollection mocks base method.
func (m *MockTigrisServer) CreateOrUpdateCollection(arg0 context.Context, arg1 *api.CreateOrUpdateCollectionRequest) (*api.CreateOrUpdateCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateCollection", arg0, arg1)
	ret0, _ := ret[0].(*api.CreateOrUpdateCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdateCollection indicates an expected call of CreateOrUpdateCollection.
func (mr *MockTigrisServerMockRecorder) CreateOrUpdateCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateCollection", reflect.TypeOf((*MockTigrisServer)(nil).CreateOrUpdateCollection), arg0, arg1)
}

// CreateProject mocks base method.
func (m *MockTigrisServer) CreateProject(arg0 context.Context, arg1 *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*api.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockTigrisServerMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockTigrisServer)(nil).CreateProject), arg0, arg1)
}

// Delete mocks base method.
func (m *MockTigrisServer) Delete(arg0 context.Context, arg1 *api.DeleteRequest) (*api.DeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*api.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTigrisServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTigrisServer)(nil).Delete), arg0, arg1)
}

// DeleteAppKey mocks base method.
func (m *MockTigrisServer) DeleteAppKey(arg0 context.Context, arg1 *api.DeleteAppKeyRequest) (*api.DeleteAppKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppKey", arg0, arg1)
	ret0, _ := ret[0].(*api.DeleteAppKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAppKey indicates an expected call of DeleteAppKey.
func (mr *MockTigrisServerMockRecorder) DeleteAppKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppKey", reflect.TypeOf((*MockTigrisServer)(nil).DeleteAppKey), arg0, arg1)
}

// DeleteBranch mocks base method.
func (m *MockTigrisServer) DeleteBranch(arg0 context.Context, arg1 *api.DeleteBranchRequest) (*api.DeleteBranchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBranch", arg0, arg1)
	ret0, _ := ret[0].(*api.DeleteBranchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBranch indicates an expected call of DeleteBranch.
func (mr *MockTigrisServerMockRecorder) DeleteBranch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBranch", reflect.TypeOf((*MockTigrisServer)(nil).DeleteBranch), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockTigrisServer) DeleteProject(arg0 context.Context, arg1 *api.DeleteProjectRequest) (*api.DeleteProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(*api.DeleteProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockTigrisServerMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockTigrisServer)(nil).DeleteProject), arg0, arg1)
}

// DescribeCollection mocks base method.
func (m *MockTigrisServer) DescribeCollection(arg0 context.Context, arg1 *api.DescribeCollectionRequest) (*api.DescribeCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCollection", arg0, arg1)
	ret0, _ := ret[0].(*api.DescribeCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCollection indicates an expected call of DescribeCollection.
func (mr *MockTigrisServerMockRecorder) DescribeCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCollection", reflect.TypeOf((*MockTigrisServer)(nil).DescribeCollection), arg0, arg1)
}

// DescribeDatabase mocks base method.
func (m *MockTigrisServer) DescribeDatabase(arg0 context.Context, arg1 *api.DescribeDatabaseRequest) (*api.DescribeDatabaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDatabase", arg0, arg1)
	ret0, _ := ret[0].(*api.DescribeDatabaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDatabase indicates an expected call of DescribeDatabase.
func (mr *MockTigrisServerMockRecorder) DescribeDatabase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDatabase", reflect.TypeOf((*MockTigrisServer)(nil).DescribeDatabase), arg0, arg1)
}

// DropCollection mocks base method.
func (m *MockTigrisServer) DropCollection(arg0 context.Context, arg1 *api.DropCollectionRequest) (*api.DropCollectionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropCollection", arg0, arg1)
	ret0, _ := ret[0].(*api.DropCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DropCollection indicates an expected call of DropCollection.
func (mr *MockTigrisServerMockRecorder) DropCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockTigrisServer)(nil).DropCollection), arg0, arg1)
}

// Import mocks base method.
func (m *MockTigrisServer) Import(arg0 context.Context, arg1 *api.ImportRequest) (*api.ImportResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", arg0, arg1)
	ret0, _ := ret[0].(*api.ImportResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Import indicates an expected call of Import.
func (mr *MockTigrisServerMockRecorder) Import(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockTigrisServer)(nil).Import), arg0, arg1)
}

// Insert mocks base method.
func (m *MockTigrisServer) Insert(arg0 context.Context, arg1 *api.InsertRequest) (*api.InsertResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(*api.InsertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTigrisServerMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTigrisServer)(nil).Insert), arg0, arg1)
}

// ListAppKeys mocks base method.
func (m *MockTigrisServer) ListAppKeys(arg0 context.Context, arg1 *api.ListAppKeysRequest) (*api.ListAppKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAppKeys", arg0, arg1)
	ret0, _ := ret[0].(*api.ListAppKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAppKeys indicates an expected call of ListAppKeys.
func (mr *MockTigrisServerMockRecorder) ListAppKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAppKeys", reflect.TypeOf((*MockTigrisServer)(nil).ListAppKeys), arg0, arg1)
}

// ListCollections mocks base method.
func (m *MockTigrisServer) ListCollections(arg0 context.Context, arg1 *api.ListCollectionsRequest) (*api.ListCollectionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCollections", arg0, arg1)
	ret0, _ := ret[0].(*api.ListCollectionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollections indicates an expected call of ListCollections.
func (mr *MockTigrisServerMockRecorder) ListCollections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollections", reflect.TypeOf((*MockTigrisServer)(nil).ListCollections), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockTigrisServer) ListProjects(arg0 context.Context, arg1 *api.ListProjectsRequest) (*api.ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(*api.ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockTigrisServerMockRecorder) ListProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockTigrisServer)(nil).ListProjects), arg0, arg1)
}

// Read mocks base method.
func (m *MockTigrisServer) Read(arg0 *api.ReadRequest, arg1 api.Tigris_ReadServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *MockTigrisServerMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTigrisServer)(nil).Read), arg0, arg1)
}

// Replace mocks base method.
func (m *MockTigrisServer) Replace(arg0 context.Context, arg1 *api.ReplaceRequest) (*api.ReplaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", arg0, arg1)
	ret0, _ := ret[0].(*api.ReplaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Replace indicates an expected call of Replace.
func (mr *MockTigrisServerMockRecorder) Replace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockTigrisServer)(nil).Replace), arg0, arg1)
}

// RollbackTransaction mocks base method.
func (m *MockTigrisServer) RollbackTransaction(arg0 context.Context, arg1 *api.RollbackTransactionRequest) (*api.RollbackTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTransaction", arg0, arg1)
	ret0, _ := ret[0].(*api.RollbackTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockTigrisServerMockRecorder) RollbackTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockTigrisServer)(nil).RollbackTransaction), arg0, arg1)
}

// RotateAppKeySecret mocks base method.
func (m *MockTigrisServer) RotateAppKeySecret(arg0 context.Context, arg1 *api.RotateAppKeyRequest) (*api.RotateAppKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateAppKeySecret", arg0, arg1)
	ret0, _ := ret[0].(*api.RotateAppKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateAppKeySecret indicates an expected call of RotateAppKeySecret.
func (mr *MockTigrisServerMockRecorder) RotateAppKeySecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateAppKeySecret", reflect.TypeOf((*MockTigrisServer)(nil).RotateAppKeySecret), arg0, arg1)
}

// Search mocks base method.
func (m *MockTigrisServer) Search(arg0 *api.SearchRequest, arg1 api.Tigris_SearchServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Search indicates an expected call of Search.
func (mr *MockTigrisServerMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockTigrisServer)(nil).Search), arg0, arg1)
}

// Update mocks base method.
func (m *MockTigrisServer) Update(arg0 context.Context, arg1 *api.UpdateRequest) (*api.UpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*api.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTigrisServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTigrisServer)(nil).Update), arg0, arg1)
}

// UpdateAppKey mocks base method.
func (m *MockTigrisServer) UpdateAppKey(arg0 context.Context, arg1 *api.UpdateAppKeyRequest) (*api.UpdateAppKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppKey", arg0, arg1)
	ret0, _ := ret[0].(*api.UpdateAppKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAppKey indicates an expected call of UpdateAppKey.
func (mr *MockTigrisServerMockRecorder) UpdateAppKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppKey", reflect.TypeOf((*MockTigrisServer)(nil).UpdateAppKey), arg0, arg1)
}

// MockAuthServer is a mock of AuthServer interface.
type MockAuthServer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServerMockRecorder
}

// MockAuthServerMockRecorder is the mock recorder for MockAuthServer.
type MockAuthServerMockRecorder struct {
	mock *MockAuthServer
}

// NewMockAuthServer creates a new mock instance.
func NewMockAuthServer(ctrl *gomock.Controller) *MockAuthServer {
	mock := &MockAuthServer{ctrl: ctrl}
	mock.recorder = &MockAuthServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthServer) EXPECT() *MockAuthServerMockRecorder {
	return m.recorder
}

// GetAccessToken mocks base method.
func (m *MockAuthServer) GetAccessToken(arg0 context.Context, arg1 *api.GetAccessTokenRequest) (*api.GetAccessTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken", arg0, arg1)
	ret0, _ := ret[0].(*api.GetAccessTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken.
func (mr *MockAuthServerMockRecorder) GetAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockAuthServer)(nil).GetAccessToken), arg0, arg1)
}

// MockManagementServer is a mock of ManagementServer interface.
type MockManagementServer struct {
	ctrl     *gomock.Controller
	recorder *MockManagementServerMockRecorder
}

// MockManagementServerMockRecorder is the mock recorder for MockManagementServer.
type MockManagementServerMockRecorder struct {
	mock *MockManagementServer
}

// NewMockManagementServer creates a new mock instance.
func NewMockManagementServer(ctrl *gomock.Controller) *MockManagementServer {
	mock := &MockManagementServer{ctrl: ctrl}
	mock.recorder = &MockManagementServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagementServer) EXPECT() *MockManagementServerMockRecorder {
	return m.recorder
}

// CreateNamespace mocks base method.
func (m *MockManagementServer) CreateNamespace(arg0 context.Context, arg1 *api.CreateNamespaceRequest) (*api.CreateNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0, arg1)
	ret0, _ := ret[0].(*api.CreateNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockManagementServerMockRecorder) CreateNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockManagementServer)(nil).CreateNamespace), arg0, arg1)
}

// DescribeNamespaces mocks base method.
func (m *MockManagementServer) DescribeNamespaces(arg0 context.Context, arg1 *api.DescribeNamespacesRequest) (*api.DescribeNamespacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeNamespaces", arg0, arg1)
	ret0, _ := ret[0].(*api.DescribeNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNamespaces indicates an expected call of DescribeNamespaces.
func (mr *MockManagementServerMockRecorder) DescribeNamespaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNamespaces", reflect.TypeOf((*MockManagementServer)(nil).DescribeNamespaces), arg0, arg1)
}

// GetNamespaceMetadata mocks base method.
func (m *MockManagementServer) GetNamespaceMetadata(arg0 context.Context, arg1 *api.GetNamespaceMetadataRequest) (*api.GetNamespaceMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.GetNamespaceMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceMetadata indicates an expected call of GetNamespaceMetadata.
func (mr *MockManagementServerMockRecorder) GetNamespaceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceMetadata", reflect.TypeOf((*MockManagementServer)(nil).GetNamespaceMetadata), arg0, arg1)
}

// GetUserMetadata mocks base method.
func (m *MockManagementServer) GetUserMetadata(arg0 context.Context, arg1 *api.GetUserMetadataRequest) (*api.GetUserMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.GetUserMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserMetadata indicates an expected call of GetUserMetadata.
func (mr *MockManagementServerMockRecorder) GetUserMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserMetadata", reflect.TypeOf((*MockManagementServer)(nil).GetUserMetadata), arg0, arg1)
}

// InsertNamespaceMetadata mocks base method.
func (m *MockManagementServer) InsertNamespaceMetadata(arg0 context.Context, arg1 *api.InsertNamespaceMetadataRequest) (*api.InsertNamespaceMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNamespaceMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.InsertNamespaceMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNamespaceMetadata indicates an expected call of InsertNamespaceMetadata.
func (mr *MockManagementServerMockRecorder) InsertNamespaceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNamespaceMetadata", reflect.TypeOf((*MockManagementServer)(nil).InsertNamespaceMetadata), arg0, arg1)
}

// InsertUserMetadata mocks base method.
func (m *MockManagementServer) InsertUserMetadata(arg0 context.Context, arg1 *api.InsertUserMetadataRequest) (*api.InsertUserMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.InsertUserMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserMetadata indicates an expected call of InsertUserMetadata.
func (mr *MockManagementServerMockRecorder) InsertUserMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserMetadata", reflect.TypeOf((*MockManagementServer)(nil).InsertUserMetadata), arg0, arg1)
}

// ListNamespaces mocks base method.
func (m *MockManagementServer) ListNamespaces(arg0 context.Context, arg1 *api.ListNamespacesRequest) (*api.ListNamespacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0, arg1)
	ret0, _ := ret[0].(*api.ListNamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockManagementServerMockRecorder) ListNamespaces(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockManagementServer)(nil).ListNamespaces), arg0, arg1)
}

// UpdateNamespaceMetadata mocks base method.
func (m *MockManagementServer) UpdateNamespaceMetadata(arg0 context.Context, arg1 *api.UpdateNamespaceMetadataRequest) (*api.UpdateNamespaceMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNamespaceMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.UpdateNamespaceMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNamespaceMetadata indicates an expected call of UpdateNamespaceMetadata.
func (mr *MockManagementServerMockRecorder) UpdateNamespaceMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNamespaceMetadata", reflect.TypeOf((*MockManagementServer)(nil).UpdateNamespaceMetadata), arg0, arg1)
}

// UpdateUserMetadata mocks base method.
func (m *MockManagementServer) UpdateUserMetadata(arg0 context.Context, arg1 *api.UpdateUserMetadataRequest) (*api.UpdateUserMetadataResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserMetadata", arg0, arg1)
	ret0, _ := ret[0].(*api.UpdateUserMetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserMetadata indicates an expected call of UpdateUserMetadata.
func (mr *MockManagementServerMockRecorder) UpdateUserMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserMetadata", reflect.TypeOf((*MockManagementServer)(nil).UpdateUserMetadata), arg0, arg1)
}

// MockObservabilityServer is a mock of ObservabilityServer interface.
type MockObservabilityServer struct {
	ctrl     *gomock.Controller
	recorder *MockObservabilityServerMockRecorder
}

// MockObservabilityServerMockRecorder is the mock recorder for MockObservabilityServer.
type MockObservabilityServerMockRecorder struct {
	mock *MockObservabilityServer
}

// NewMockObservabilityServer creates a new mock instance.
func NewMockObservabilityServer(ctrl *gomock.Controller) *MockObservabilityServer {
	mock := &MockObservabilityServer{ctrl: ctrl}
	mock.recorder = &MockObservabilityServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObservabilityServer) EXPECT() *MockObservabilityServerMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockObservabilityServer) GetInfo(arg0 context.Context, arg1 *api.GetInfoRequest) (*api.GetInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", arg0, arg1)
	ret0, _ := ret[0].(*api.GetInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockObservabilityServerMockRecorder) GetInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockObservabilityServer)(nil).GetInfo), arg0, arg1)
}

// QueryTimeSeriesMetrics mocks base method.
func (m *MockObservabilityServer) QueryTimeSeriesMetrics(arg0 context.Context, arg1 *api.QueryTimeSeriesMetricsRequest) (*api.QueryTimeSeriesMetricsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTimeSeriesMetrics", arg0, arg1)
	ret0, _ := ret[0].(*api.QueryTimeSeriesMetricsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTimeSeriesMetrics indicates an expected call of QueryTimeSeriesMetrics.
func (mr *MockObservabilityServerMockRecorder) QueryTimeSeriesMetrics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTimeSeriesMetrics", reflect.TypeOf((*MockObservabilityServer)(nil).QueryTimeSeriesMetrics), arg0, arg1)
}

// QuotaLimits mocks base method.
func (m *MockObservabilityServer) QuotaLimits(arg0 context.Context, arg1 *api.QuotaLimitsRequest) (*api.QuotaLimitsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaLimits", arg0, arg1)
	ret0, _ := ret[0].(*api.QuotaLimitsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotaLimits indicates an expected call of QuotaLimits.
func (mr *MockObservabilityServerMockRecorder) QuotaLimits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaLimits", reflect.TypeOf((*MockObservabilityServer)(nil).QuotaLimits), arg0, arg1)
}

// QuotaUsage mocks base method.
func (m *MockObservabilityServer) QuotaUsage(arg0 context.Context, arg1 *api.QuotaUsageRequest) (*api.QuotaUsageResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuotaUsage", arg0, arg1)
	ret0, _ := ret[0].(*api.QuotaUsageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotaUsage indicates an expected call of QuotaUsage.
func (mr *MockObservabilityServerMockRecorder) QuotaUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaUsage", reflect.TypeOf((*MockObservabilityServer)(nil).QuotaUsage), arg0, arg1)
}
