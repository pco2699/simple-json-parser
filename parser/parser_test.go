package parser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		tokens []interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 []interface{}
	}{
		{
			name: "normal case",
			args: args{
				tokens: []interface{}{'{', "hoge", ':', int32(10), '}'},
			},
			want:  []interface{}{"hoge", int32(10)},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Parse(tt.args.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseArray(t *testing.T) {
	type args struct {
		tokens []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		want1   []interface{}
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				tokens: []interface{}{"hoge", int32(10), ']'},
			},
			want: []interface{}{"hoge", int32(10)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseArray(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArray() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseArray() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseObject(t *testing.T) {
	type args struct {
		tokens []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		want1   []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseObject(tt.args.tokens)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseObject() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseObject() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
