// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/henvic/pgxtutorial/internal/inventory (interfaces: DB)

// Package inventory is a generated GoMock package.
package inventory

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockDB) CreateProduct(arg0 context.Context, arg1 CreateProductParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockDBMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockDB)(nil).CreateProduct), arg0, arg1)
}

// CreateProductReview mocks base method.
func (m *MockDB) CreateProductReview(arg0 context.Context, arg1 CreateProductReviewDBParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProductReview indicates an expected call of CreateProductReview.
func (mr *MockDBMockRecorder) CreateProductReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductReview", reflect.TypeOf((*MockDB)(nil).CreateProductReview), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockDB) DeleteProduct(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockDBMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockDB)(nil).DeleteProduct), arg0, arg1)
}

// DeleteProductReview mocks base method.
func (m *MockDB) DeleteProductReview(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductReview indicates an expected call of DeleteProductReview.
func (mr *MockDBMockRecorder) DeleteProductReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductReview", reflect.TypeOf((*MockDB)(nil).DeleteProductReview), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockDB) GetProduct(arg0 context.Context, arg1 string) (*Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(*Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockDBMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockDB)(nil).GetProduct), arg0, arg1)
}

// GetProductReview mocks base method.
func (m *MockDB) GetProductReview(arg0 context.Context, arg1 string) (*ProductReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductReview", arg0, arg1)
	ret0, _ := ret[0].(*ProductReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductReview indicates an expected call of GetProductReview.
func (mr *MockDBMockRecorder) GetProductReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductReview", reflect.TypeOf((*MockDB)(nil).GetProductReview), arg0, arg1)
}

// GetProductReviews mocks base method.
func (m *MockDB) GetProductReviews(arg0 context.Context, arg1 ProductReviewsParams) (*ProductReviewsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductReviews", arg0, arg1)
	ret0, _ := ret[0].(*ProductReviewsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductReviews indicates an expected call of GetProductReviews.
func (mr *MockDBMockRecorder) GetProductReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductReviews", reflect.TypeOf((*MockDB)(nil).GetProductReviews), arg0, arg1)
}

// SearchProducts mocks base method.
func (m *MockDB) SearchProducts(arg0 context.Context, arg1 SearchProductsParams) (*SearchProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProducts", arg0, arg1)
	ret0, _ := ret[0].(*SearchProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockDBMockRecorder) SearchProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockDB)(nil).SearchProducts), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockDB) UpdateProduct(arg0 context.Context, arg1 UpdateProductParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockDBMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockDB)(nil).UpdateProduct), arg0, arg1)
}

// UpdateProductReview mocks base method.
func (m *MockDB) UpdateProductReview(arg0 context.Context, arg1 UpdateProductReviewParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductReview", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductReview indicates an expected call of UpdateProductReview.
func (mr *MockDBMockRecorder) UpdateProductReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductReview", reflect.TypeOf((*MockDB)(nil).UpdateProductReview), arg0, arg1)
}