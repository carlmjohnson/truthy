package truthy_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/carlmjohnson/truthy"
)

type errT struct{}

func (_ *errT) Error() string { return "" }

func test[T any](t *testing.T, v T, ok bool) {
	t.Run(fmt.Sprintf("%T-%v-%v", v, v, ok), func(t *testing.T) {
		if got := truthy.Value(v); got != ok {
			t.Fatal()
		}
	})
}

func TestValue(t *testing.T) {
	var err error
	test(t, err, false)
	err = (*errT)(nil)
	test(t, err, true)
	var errp *errT
	test(t, errp, false)
	test(t, "", false)
	test(t, " ", true)
	test(t, []byte{}, false)
	test(t, []byte(" "), true)
	var f func()
	test(t, f, false)
	f = func() {}
	test(t, f, true)
	var s struct{}
	test(t, s, false)
	test(t, &s, true)
	test(t, (*struct{})(nil), false)
	var i interface{}
	test(t, i, false)
	i = (*errT)(nil)
	test(t, i, true)
	test(t, 10, true)
	test(t, 0, false)
	test(t, 1.1, true)
	test(t, 0.0, false)
	var ch chan int
	test(t, ch, false)
	ch = make(chan int)
	test(t, ch, true)
	m := map[string]string{}
	test(t, m, false)
	m["one"] = "one"
	test(t, m, true)
	var a [2]int
	test(t, a, false)
	a = [2]int{0, 1}
	test(t, a, true)
	var cron time.Time
	test(t, cron, false)
	test(t, cron.In(time.Local), false)
	test(t, time.Now(), true)
}

func BenchmarkValue_error(b *testing.B) {
	fillVal := errors.New("something")
	fill := false
	for i := 0; i < b.N; i++ {
		var value error
		if fill {
			value = fillVal
		}
		if truthy.Value(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkValue_string(b *testing.B) {
	fillVal := "something"
	fill := false
	for i := 0; i < b.N; i++ {
		var value string
		if fill {
			value = fillVal
		}
		if truthy.Value(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}

func BenchmarkValue_int(b *testing.B) {
	fillVal := 1
	fill := false
	for i := 0; i < b.N; i++ {
		var value int
		if fill {
			value = fillVal
		}
		if truthy.Value(value) != fill {
			b.FailNow()
		}
		fill = !fill
	}
}
