package parser

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	type args struct {
		str []rune
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lex(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		word    rune
		wordMap map[rune]struct{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.word, tt.args.wordMap); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexString(t *testing.T) {
	type args struct {
		str []rune
	}
	tests := []struct {
		name    string
		args    args
		want    []rune
		want1   []rune
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				str: []rune("\"hoge\""),
			},
			want:    []rune("hoge"),
			want1:   []rune(""),
			wantErr: false,
		},
		{
			name: "normal case2",
			args: args{
				str: []rune("\"hoge\",\"fuga\""),
			},
			want:    []rune("hoge"),
			want1:   []rune(",\"fuga\""),
			wantErr: false,
		},
		{
			name: "it's not a string case",
			args: args{
				str: []rune("1231"),
			},
			want:    nil,
			want1:   []rune("1231"),
			wantErr: false,
		},
		{
			name: "error case",
			args: args{
				str: []rune("\"hoge"),
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := lexString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("lexString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexString() got = %v, want %v", string(got), string(tt.want))
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("lexString() got1 = %v, want %v", string(got1), string(tt.want1))
			}
		})
	}
}

func Test_lexNumber(t *testing.T) {
	type args struct {
		str []rune
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		want1   []rune
		wantErr bool
	}{
		{
			name: "integer case",
			args: args{
				str: []rune("1234"),
			},
			want:    int32(1234),
			want1:   []rune(""),
			wantErr: false,
		},
		{
			name: "minus case",
			args: args{
				str: []rune("-1234"),
			},
			want:    int32(-1234),
			want1:   []rune(""),
			wantErr: false,
		},
		{
			name: "integer case2",
			args: args{
				str: []rune("1234, {\"hoge\"}"),
			},
			want:    int32(1234),
			want1:   []rune(", {\"hoge\"}"),
			wantErr: false,
		},
		{
			name: "float case",
			args: args{
				str: []rune("12.31"),
			},
			want:    float32(12.31),
			want1:   []rune(""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := lexNumber(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("lexNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexNumber() got = %v %T, want %v %T", got, got, tt.want, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("lexNumber() got1 = %v, %T want %v, %T", got1, got1, tt.want1, tt.want1)
			}
		})
	}
}

func Test_lexBool(t *testing.T) {
	type args struct {
		str []rune
	}
	tests := []struct {
		name     string
		args     args
		wantRes  bool
		wantOk   bool
		wantRest []rune
		wantErr  bool
	}{
		{
			name: "true case",
			args: args{
				str: []rune("true"),
			},
			wantRes:  true,
			wantOk:   true,
			wantRest: []rune(""),
		},
		{
			name: "false case",
			args: args{
				str: []rune("false"),
			},
			wantRes:  false,
			wantOk:   true,
			wantRest: []rune(""),
		},
		{
			name: "not bool case",
			args: args{
				str: []rune("\"hoge\""),
			},
			wantRes:  false,
			wantOk:   false,
			wantRest: []rune("\"hoge\""),
		},
		{
			name: "falsee case",
			args: args{
				str: []rune("falsee"),
			},
			wantRes:  false,
			wantOk:   true,
			wantRest: []rune("e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk, gotRest := lexBool(tt.args.str)
			if gotRes != tt.wantRes {
				t.Errorf("lexBool() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("lexBool() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotRest, tt.wantRest) {
				t.Errorf("lexBool() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
		})
	}
}
