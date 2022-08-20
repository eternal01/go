package gocache

import (
	"testing"
)

func TestByteView_Len(t *testing.T) {
	type fields struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{name: "test", fields: fields{b: []byte("123")}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ByteView{
				b: tt.fields.b,
			}
			if got := v.Len(); got != tt.want {
				t.Errorf("ByteView.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteView_String(t *testing.T) {
	type fields struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "test", fields: fields{b: []byte("123")}, want: "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ByteView{
				b: tt.fields.b,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("ByteView.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
