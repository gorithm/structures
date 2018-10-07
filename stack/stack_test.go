package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			s.Push(tt.args.val)
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			got, err := s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Stack.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
