package json_hlp

import (
	"reflect"
	"testing"
)

func TestLoadjson(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Loadjson(tt.args.path, 10240)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Loadjson() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Loadjson() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name  string
		args  args
		want  map[string]interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMapToJson(t *testing.T) {
	type args struct {
		param map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToJson(tt.args.param); got != tt.want {
				t.Errorf("MapToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
