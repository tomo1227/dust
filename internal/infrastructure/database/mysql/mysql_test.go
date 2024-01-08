package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_start(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start()
		})
	}
}
