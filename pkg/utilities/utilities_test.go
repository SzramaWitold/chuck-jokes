package utilities

import (
	"reflect"
	"testing"
	"time"
)

func TestGetToday(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "base test",
			want: getDay(time.Now()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetToday(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYesterday(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "base test",
			want: getDay(time.Now().Add(-24 * time.Hour)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYesterday(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYesterday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDay(t *testing.T) {
	now := time.Now()
	year, month, day := now.Date()
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "base test",
			args: args{t: time.Now()},
			want: time.Date(year, month, day, 0, 0, 0, 0, now.Location()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDay(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
