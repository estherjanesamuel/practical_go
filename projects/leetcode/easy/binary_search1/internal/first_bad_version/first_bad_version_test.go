package firstbadversion

import "testing"

var Result int

func TestFirstBadVersion(t *testing.T) {
	got := firstBadVersionMine(2147483647)
	want :=	1
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func BenchmarkMine(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersionMine(1000)
	}

	Result = r
}

func BenchmarkBadVer(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVer(1000)
	}

	Result = r
}
func BenchmarkV1(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersion(1000)
	}

	Result = r
}

func BenchmarkV2(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersionV2(1000)
	}

	Result = r
}

func BenchmarkV3(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersionV3(1000)
	}

	Result = r
}

func BenchmarkV4(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersionV4(1000)
	}

	Result = r
}

func BenchmarkDiscuss1756768(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = firstBadVersionDiscuss1756768(1000)
	}

	Result = r
}

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			want: 50,
			args: args{
				n: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}