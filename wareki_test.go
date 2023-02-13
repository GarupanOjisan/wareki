package wareki

import (
	"reflect"
	"testing"
	"time"
)

func TestIsMeiji(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1868年9月7日は慶応",
			args: args{
				t: time.Date(1868, 10, 22, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
		{
			name: "1868年10月23日は明治",
			args: args{
				t: time.Date(1868, 10, 23, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1912年7月30日は明治",
			args: args{
				t: time.Date(1912, 7, 30, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1912年7月31日は大正",
			args: args{
				t: time.Date(1912, 7, 31, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMeiji(tt.args.t); got != tt.want {
				t.Errorf("IsMeiji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTaisho(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1912年7月29日は明治",
			args: args{
				t: time.Date(1912, 7, 29, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
		{
			name: "1912年7月30日は大正",
			args: args{
				t: time.Date(1912, 7, 30, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1926年12月25日は大正",
			args: args{
				t: time.Date(1926, 12, 25, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1926年12月26日は昭和",
			args: args{
				t: time.Date(1926, 12, 26, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTaisho(tt.args.t); got != tt.want {
				t.Errorf("IsTaisho() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsShowa(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1926年12月24日は大正",
			args: args{
				t: time.Date(1926, 12, 24, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
		{
			name: "1926年12月25日は昭和",
			args: args{
				t: time.Date(1926, 12, 25, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1989年1月7日は昭和",
			args: args{
				t: time.Date(1989, 1, 7, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "1989年1月8日は平成",
			args: args{
				t: time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsShowa(tt.args.t); got != tt.want {
				t.Errorf("IsShowa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHeisei(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1989年1月7日は昭和",
			args: args{
				t: time.Date(1989, 1, 7, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
		{
			name: "1989年1月8日は平成",
			args: args{
				t: time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "2019年4月30日は平成",
			args: args{
				t: time.Date(2019, 4, 30, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
		{
			name: "2019年5月1日は令和",
			args: args{
				t: time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHeisei(tt.args.t); got != tt.want {
				t.Errorf("IsHeisei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsReiwa(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2019年4月30日は平成",
			args: args{
				t: time.Date(2019, 4, 30, 0, 0, 0, 0, locationJST),
			},
			want: false,
		},
		{
			name: "2019年5月1日は令和",
			args: args{
				t: time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReiwa(tt.args.t); got != tt.want {
				t.Errorf("IsReiwa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWareki(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Wareki
	}{
		{
			name: "1868年10月22日はサポート対象外",
			args: args{
				t: time.Date(1868, 10, 22, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{},
		},
		{
			name: "1868年10月23日は明治1年が返る",
			args: args{
				t: time.Date(1868, 10, 23, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "明治",
				Year:  1,
			},
		},
		{
			name: "1912年7月30日は明治45年と大正元年が返る",
			args: args{
				t: time.Date(1912, 7, 30, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "大正",
				Year:  1,
			},
		},
		{
			name: "1926年12月25日は大正15年と昭和元年が返る",
			args: args{
				t: time.Date(1926, 12, 25, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "昭和",
				Year:  1,
			},
		},
		{
			name: "1989年1月7日は昭和64年が返る",
			args: args{
				t: time.Date(1989, 1, 7, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "昭和",
				Year:  64,
			},
		},
		{
			name: "1989年1月8日は平成元年が返る",
			args: args{
				t: time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "平成",
				Year:  1,
			},
		},
		{
			name: "2019年4月30日は平成31年が返る",
			args: args{
				t: time.Date(2019, 4, 30, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "平成",
				Year:  31,
			},
		},
		{
			name: "2019年5月1日は令和元年が返る",
			args: args{
				t: time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST),
			},
			want: Wareki{
				Gengo: "令和",
				Year:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestWareki_Time は、Wareki型のTimeメソッドのテストです。
func TestWareki_Time(t *testing.T) {
	type fields struct {
		Gengo string
		Year  int
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "明治1年は1868年10月23日",
			fields: fields{
				Gengo: "明治",
				Year:  1,
			},
			want: time.Date(1868, 10, 23, 0, 0, 0, 0, locationJST),
		},
		{
			name: "大正1年は1912年7月30日",
			fields: fields{
				Gengo: "大正",
				Year:  1,
			},
			want: time.Date(1912, 7, 30, 0, 0, 0, 0, locationJST),
		},
		{
			name: "昭和1年は1926年12月25日",
			fields: fields{
				Gengo: "昭和",
				Year:  1,
			},
			want: time.Date(1926, 12, 25, 0, 0, 0, 0, locationJST),
		},
		{
			name: "平成1年は1989年1月8日",
			fields: fields{
				Gengo: "平成",
				Year:  1,
			},
			want: time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST),
		},
		{
			name: "令和1年は2019年5月1日",
			fields: fields{
				Gengo: "令和",
				Year:  1,
			},
			want: time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Wareki{
				Gengo: tt.fields.Gengo,
				Year:  tt.fields.Year,
			}
			if got := w.Time(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}
