package task_repository

import (
	"reflect"
	"testing"
)

func TestNewInTaskRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *InTaskRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInTaskRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInTaskRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
