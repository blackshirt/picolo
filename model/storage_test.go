package model

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

func TestGetRupItemWithKey(t *testing.T) {
	mock := r.NewMock()
	query := r.Table("rup_item").Get("020708eb-09b1-44de-a880-9dc769224b41")
	term := mock.On(query)
	term.Return([]interface{}{
		map[string]interface{}{
			"detilWaktu": "",
			"id":         "020708eb-09b1-44de-a880-9dc769224b41",
			"jenis":      "",
			"kategori":   "Penyedia",
			"kodeOpd":    "63408",
			"kodeRup":    "22836111",
			"metode":     "Tender",
			"namaPaket":  "Rehabilitasi TPI Pasir",
			"pagu":       "350000000",
			"state":      "",
			"sumberDana": "APBD",
			"tahun":      "2020",
			"waktu":      "April 2020"},
	}, nil)

	cursor, err := query.Run(mock)
	if err != nil {
		t.Errorf("err is: %v", err)
	}

	var rows []interface{}
	err = cursor.All(&rows)
	if err != nil {
		t.Errorf("err is: %v", err)
	}

	// Test result of rows

	mock.AssertExpectations(t)
	mock.AssertExecuted(t, term)
}

func TestNewService(t *testing.T) {
	var ctx context.Context = context.TODO()
	// var w waktu
	// Create a new instance of the mock store
	m := new(mockStorage)
	// In the "On" method, we assert that we want the "Get" method
	// to be called with one argument, that is 2
	// In the "Return" method, we define the return values to be 7, and nil (for the result and error values)
	m.On("Rup", ctx, "020708eb-09b1-44de-a880-9dc769224b41").Return(
		RupItem{
			DetilWaktu: nil,
			ID:         "020708eb-09b1-44de-a880-9dc769224b41",
			Jenis:      nil,
			Kategori:   KategoriPenyedia,
			KodeOpd:    "63408",
			KodeRup:    "22836111",
			Metode:     Tender,
			NamaPaket:  "Rehabilitasi TPI Pasir",
			Pagu:       "350000000",
			State:      nil,
			SumberDana: "APBD",
			Tahun:      "2020",
			Waktu:      "April 2020"}, nil)
	// Next, we create a new instance of our module with the mock store as its "store" dependency
	s := NewService(m)
	// The "Get" method call is then made
	_, err := s.Get(ctx, "020708eb-09b1-44de-a880-9dc769224b41")
	// The expectations that we defined for our mock store earlier are asserted here
	m.AssertExpectations(t)
	// Finally, we assert that we should'nt get any error
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}

// mockStorage is an autogenerated mock type for the Storage type
type mockStorage struct {
	mock.Mock
}

// AllRup provides a mock function with given fields: ctx
func (_m *mockStorage) AllRup(ctx context.Context) ([]RupItem, error) {
	ret := _m.Called(ctx)

	var r0 []RupItem
	if rf, ok := ret.Get(0).(func(context.Context) []RupItem); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RupItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: ctx, t, key
func (_m *mockStorage) Exists(ctx context.Context, t Type, key string) bool {
	ret := _m.Called(ctx, t, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, Type, string) bool); ok {
		r0 = rf(ctx, t, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FilteredRup provides a mock function with given fields: ctx, opt
func (_m *mockStorage) FilteredRup(ctx context.Context, opt RupOptions) ([]RupItem, error) {
	ret := _m.Called(ctx, opt)

	var r0 []RupItem
	if rf, ok := ret.Get(0).(func(context.Context, RupOptions) []RupItem); ok {
		r0 = rf(ctx, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RupItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, RupOptions) error); ok {
		r1 = rf(ctx, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Opd provides a mock function with given fields: ctx, id
func (_m *mockStorage) Opd(ctx context.Context, id string) (OpdItem, error) {
	ret := _m.Called(ctx, id)

	var r0 OpdItem
	if rf, ok := ret.Get(0).(func(context.Context, string) OpdItem); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(OpdItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Rup provides a mock function with given fields: ctx, id
func (_m *mockStorage) Rup(ctx context.Context, id string) (RupItem, error) {
	ret := _m.Called(ctx, id)

	var r0 RupItem
	if rf, ok := ret.Get(0).(func(context.Context, string) RupItem); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(RupItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RupForOpd provides a mock function with given fields: ctx, opd
func (_m *mockStorage) RupForOpd(ctx context.Context, opd OpdItem) ([]RupItem, error) {
	ret := _m.Called(ctx, opd)

	var r0 []RupItem
	if rf, ok := ret.Get(0).(func(context.Context, OpdItem) []RupItem); ok {
		r0 = rf(ctx, opd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RupItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, OpdItem) error); ok {
		r1 = rf(ctx, opd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOpd provides a mock function with given fields: ctx, objs
func (_m *mockStorage) SaveOpd(ctx context.Context, objs []OpdItem) error {
	ret := _m.Called(ctx, objs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []OpdItem) error); ok {
		r0 = rf(ctx, objs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveRup provides a mock function with given fields: ctx, objs
func (_m *mockStorage) SaveRup(ctx context.Context, objs []RupItem) error {
	ret := _m.Called(ctx, objs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []RupItem) error); ok {
		r0 = rf(ctx, objs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
