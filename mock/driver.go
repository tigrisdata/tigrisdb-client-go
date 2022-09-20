// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tigrisdata/tigris-client-go/driver (interfaces: Driver,Tx,Database,Iterator,SearchResultIterator)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/tigrisdata/tigris-client-go/driver"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockDriver) BeginTx(arg0 context.Context, arg1 string, arg2 ...*driver.TxOptions) (driver.Tx, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BeginTx", varargs...)
	ret0, _ := ret[0].(driver.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDriverMockRecorder) BeginTx(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDriver)(nil).BeginTx), varargs...)
}

// Close mocks base method.
func (m *MockDriver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDriverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDriver)(nil).Close))
}

// CreateDatabase mocks base method.
func (m *MockDriver) CreateDatabase(arg0 context.Context, arg1 string, arg2 ...*driver.DatabaseOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDatabase", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDatabase indicates an expected call of CreateDatabase.
func (mr *MockDriverMockRecorder) CreateDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabase", reflect.TypeOf((*MockDriver)(nil).CreateDatabase), varargs...)
}

// DescribeDatabase mocks base method.
func (m *MockDriver) DescribeDatabase(arg0 context.Context, arg1 string) (*driver.DescribeDatabaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDatabase", arg0, arg1)
	ret0, _ := ret[0].(*driver.DescribeDatabaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDatabase indicates an expected call of DescribeDatabase.
func (mr *MockDriverMockRecorder) DescribeDatabase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDatabase", reflect.TypeOf((*MockDriver)(nil).DescribeDatabase), arg0, arg1)
}

// DropDatabase mocks base method.
func (m *MockDriver) DropDatabase(arg0 context.Context, arg1 string, arg2 ...*driver.DatabaseOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropDatabase", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropDatabase indicates an expected call of DropDatabase.
func (mr *MockDriverMockRecorder) DropDatabase(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropDatabase", reflect.TypeOf((*MockDriver)(nil).DropDatabase), varargs...)
}

// Info mocks base method.
func (m *MockDriver) Info(arg0 context.Context) (*driver.InfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", arg0)
	ret0, _ := ret[0].(*driver.InfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockDriverMockRecorder) Info(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockDriver)(nil).Info), arg0)
}

// ListDatabases mocks base method.
func (m *MockDriver) ListDatabases(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDatabases", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDatabases indicates an expected call of ListDatabases.
func (mr *MockDriverMockRecorder) ListDatabases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabases", reflect.TypeOf((*MockDriver)(nil).ListDatabases), arg0)
}

// UseDatabase mocks base method.
func (m *MockDriver) UseDatabase(arg0 string) driver.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseDatabase", arg0)
	ret0, _ := ret[0].(driver.Database)
	return ret0
}

// UseDatabase indicates an expected call of UseDatabase.
func (mr *MockDriverMockRecorder) UseDatabase(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseDatabase", reflect.TypeOf((*MockDriver)(nil).UseDatabase), arg0)
}

// MockTx is a mock of Tx interface.
type MockTx struct {
	ctrl     *gomock.Controller
	recorder *MockTxMockRecorder
}

// MockTxMockRecorder is the mock recorder for MockTx.
type MockTxMockRecorder struct {
	mock *MockTx
}

// NewMockTx creates a new mock instance.
func NewMockTx(ctrl *gomock.Controller) *MockTx {
	mock := &MockTx{ctrl: ctrl}
	mock.recorder = &MockTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTx) EXPECT() *MockTxMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTx) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTxMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTx)(nil).Commit), arg0)
}

// CreateOrUpdateCollection mocks base method.
func (m *MockTx) CreateOrUpdateCollection(arg0 context.Context, arg1 string, arg2 driver.Schema, arg3 ...*driver.CollectionOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateCollection", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateCollection indicates an expected call of CreateOrUpdateCollection.
func (mr *MockTxMockRecorder) CreateOrUpdateCollection(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateCollection", reflect.TypeOf((*MockTx)(nil).CreateOrUpdateCollection), varargs...)
}

// Delete mocks base method.
func (m *MockTx) Delete(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 ...*driver.DeleteOptions) (*driver.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*driver.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockTxMockRecorder) Delete(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTx)(nil).Delete), varargs...)
}

// DescribeCollection mocks base method.
func (m *MockTx) DescribeCollection(arg0 context.Context, arg1 string, arg2 ...*driver.CollectionOptions) (*driver.DescribeCollectionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCollection", varargs...)
	ret0, _ := ret[0].(*driver.DescribeCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCollection indicates an expected call of DescribeCollection.
func (mr *MockTxMockRecorder) DescribeCollection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCollection", reflect.TypeOf((*MockTx)(nil).DescribeCollection), varargs...)
}

// DropCollection mocks base method.
func (m *MockTx) DropCollection(arg0 context.Context, arg1 string, arg2 ...*driver.CollectionOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropCollection", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection.
func (mr *MockTxMockRecorder) DropCollection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockTx)(nil).DropCollection), varargs...)
}

// Insert mocks base method.
func (m *MockTx) Insert(arg0 context.Context, arg1 string, arg2 []driver.Document, arg3 ...*driver.InsertOptions) (*driver.InsertResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(*driver.InsertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockTxMockRecorder) Insert(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTx)(nil).Insert), varargs...)
}

// ListCollections mocks base method.
func (m *MockTx) ListCollections(arg0 context.Context, arg1 ...*driver.CollectionOptions) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCollections", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollections indicates an expected call of ListCollections.
func (mr *MockTxMockRecorder) ListCollections(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollections", reflect.TypeOf((*MockTx)(nil).ListCollections), varargs...)
}

// Publish mocks base method.
func (m *MockTx) Publish(arg0 context.Context, arg1 string, arg2 []driver.Message, arg3 ...*driver.PublishOptions) (*driver.PublishResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*driver.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockTxMockRecorder) Publish(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockTx)(nil).Publish), varargs...)
}

// Read mocks base method.
func (m *MockTx) Read(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 driver.Projection, arg4 ...*driver.ReadOptions) (driver.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(driver.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTxMockRecorder) Read(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTx)(nil).Read), varargs...)
}

// Replace mocks base method.
func (m *MockTx) Replace(arg0 context.Context, arg1 string, arg2 []driver.Document, arg3 ...*driver.ReplaceOptions) (*driver.ReplaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Replace", varargs...)
	ret0, _ := ret[0].(*driver.ReplaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Replace indicates an expected call of Replace.
func (mr *MockTxMockRecorder) Replace(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockTx)(nil).Replace), varargs...)
}

// Rollback mocks base method.
func (m *MockTx) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTxMockRecorder) Rollback(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTx)(nil).Rollback), arg0)
}

// Search mocks base method.
func (m *MockTx) Search(arg0 context.Context, arg1 string, arg2 *driver.SearchRequest) (driver.SearchResultIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1, arg2)
	ret0, _ := ret[0].(driver.SearchResultIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockTxMockRecorder) Search(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockTx)(nil).Search), arg0, arg1, arg2)
}

// Subscribe mocks base method.
func (m *MockTx) Subscribe(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 ...*driver.SubscribeOptions) (driver.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(driver.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockTxMockRecorder) Subscribe(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTx)(nil).Subscribe), varargs...)
}

// Update mocks base method.
func (m *MockTx) Update(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 driver.Update, arg4 ...*driver.UpdateOptions) (*driver.UpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*driver.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTxMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTx)(nil).Update), varargs...)
}

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateOrUpdateCollection mocks base method.
func (m *MockDatabase) CreateOrUpdateCollection(arg0 context.Context, arg1 string, arg2 driver.Schema, arg3 ...*driver.CollectionOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateCollection", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateCollection indicates an expected call of CreateOrUpdateCollection.
func (mr *MockDatabaseMockRecorder) CreateOrUpdateCollection(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateCollection", reflect.TypeOf((*MockDatabase)(nil).CreateOrUpdateCollection), varargs...)
}

// Delete mocks base method.
func (m *MockDatabase) Delete(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 ...*driver.DeleteOptions) (*driver.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*driver.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockDatabaseMockRecorder) Delete(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatabase)(nil).Delete), varargs...)
}

// DescribeCollection mocks base method.
func (m *MockDatabase) DescribeCollection(arg0 context.Context, arg1 string, arg2 ...*driver.CollectionOptions) (*driver.DescribeCollectionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCollection", varargs...)
	ret0, _ := ret[0].(*driver.DescribeCollectionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCollection indicates an expected call of DescribeCollection.
func (mr *MockDatabaseMockRecorder) DescribeCollection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCollection", reflect.TypeOf((*MockDatabase)(nil).DescribeCollection), varargs...)
}

// DropCollection mocks base method.
func (m *MockDatabase) DropCollection(arg0 context.Context, arg1 string, arg2 ...*driver.CollectionOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropCollection", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection.
func (mr *MockDatabaseMockRecorder) DropCollection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockDatabase)(nil).DropCollection), varargs...)
}

// Insert mocks base method.
func (m *MockDatabase) Insert(arg0 context.Context, arg1 string, arg2 []driver.Document, arg3 ...*driver.InsertOptions) (*driver.InsertResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(*driver.InsertResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockDatabaseMockRecorder) Insert(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockDatabase)(nil).Insert), varargs...)
}

// ListCollections mocks base method.
func (m *MockDatabase) ListCollections(arg0 context.Context, arg1 ...*driver.CollectionOptions) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCollections", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCollections indicates an expected call of ListCollections.
func (mr *MockDatabaseMockRecorder) ListCollections(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCollections", reflect.TypeOf((*MockDatabase)(nil).ListCollections), varargs...)
}

// Publish mocks base method.
func (m *MockDatabase) Publish(arg0 context.Context, arg1 string, arg2 []driver.Message, arg3 ...*driver.PublishOptions) (*driver.PublishResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*driver.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockDatabaseMockRecorder) Publish(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockDatabase)(nil).Publish), varargs...)
}

// Read mocks base method.
func (m *MockDatabase) Read(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 driver.Projection, arg4 ...*driver.ReadOptions) (driver.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(driver.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockDatabaseMockRecorder) Read(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockDatabase)(nil).Read), varargs...)
}

// Replace mocks base method.
func (m *MockDatabase) Replace(arg0 context.Context, arg1 string, arg2 []driver.Document, arg3 ...*driver.ReplaceOptions) (*driver.ReplaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Replace", varargs...)
	ret0, _ := ret[0].(*driver.ReplaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Replace indicates an expected call of Replace.
func (mr *MockDatabaseMockRecorder) Replace(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockDatabase)(nil).Replace), varargs...)
}

// Search mocks base method.
func (m *MockDatabase) Search(arg0 context.Context, arg1 string, arg2 *driver.SearchRequest) (driver.SearchResultIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1, arg2)
	ret0, _ := ret[0].(driver.SearchResultIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDatabaseMockRecorder) Search(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDatabase)(nil).Search), arg0, arg1, arg2)
}

// Subscribe mocks base method.
func (m *MockDatabase) Subscribe(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 ...*driver.SubscribeOptions) (driver.Iterator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(driver.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockDatabaseMockRecorder) Subscribe(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockDatabase)(nil).Subscribe), varargs...)
}

// Update mocks base method.
func (m *MockDatabase) Update(arg0 context.Context, arg1 string, arg2 driver.Filter, arg3 driver.Update, arg4 ...*driver.UpdateOptions) (*driver.UpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*driver.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDatabaseMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDatabase)(nil).Update), varargs...)
}

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Err mocks base method.
func (m *MockIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockIterator) Next(arg0 *driver.Document) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next), arg0)
}

// MockSearchResultIterator is a mock of SearchResultIterator interface.
type MockSearchResultIterator struct {
	ctrl     *gomock.Controller
	recorder *MockSearchResultIteratorMockRecorder
}

// MockSearchResultIteratorMockRecorder is the mock recorder for MockSearchResultIterator.
type MockSearchResultIteratorMockRecorder struct {
	mock *MockSearchResultIterator
}

// NewMockSearchResultIterator creates a new mock instance.
func NewMockSearchResultIterator(ctrl *gomock.Controller) *MockSearchResultIterator {
	mock := &MockSearchResultIterator{ctrl: ctrl}
	mock.recorder = &MockSearchResultIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchResultIterator) EXPECT() *MockSearchResultIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSearchResultIterator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSearchResultIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSearchResultIterator)(nil).Close))
}

// Err mocks base method.
func (m *MockSearchResultIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockSearchResultIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockSearchResultIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockSearchResultIterator) Next(arg0 *driver.SearchResponse) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockSearchResultIteratorMockRecorder) Next(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockSearchResultIterator)(nil).Next), arg0)
}
