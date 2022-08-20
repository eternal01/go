package lru

import (
	"container/list"
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestCache_RemoveOldest(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.RemoveOldest()
		})
	}
}

func TestCache_Add(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	type args struct {
		key   string
		value Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "testkey1", fields: fields(*New(10, nil)), args: args{key: "testkey1", value: String("11222")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			c.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		maxBytes  int64
		nbytes    int64
		ll        *list.List
		cache     map[string]*list.Element
		OnEvicted func(key string, value Value)
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue Value
		wantOk    bool
	}{
		// TODO: Add test cases.
		{name: "undefind", fields: fields(*New(10, nil)), args: args{key: "testkey1"}, wantValue: nil, wantOk: false},
		{name: "getOne", fields: fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				maxBytes:  tt.fields.maxBytes,
				nbytes:    tt.fields.nbytes,
				ll:        tt.fields.ll,
				cache:     tt.fields.cache,
				OnEvicted: tt.fields.OnEvicted,
			}
			gotValue, gotOk := c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Cache.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Cache.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
