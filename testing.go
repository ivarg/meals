package meals

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func ok(t testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		// for text color, also check out https://github.com/daviddengcn/go-colortext
		fmt.Print("\x1b[38;5;1m") // text color
		fname := filepath.Base(file)
		fmt.Printf("%s:%d: unexpected error: %s\n\n",
			append([]interface{}{fname, line}, err.Error())...)
		fmt.Print("\x1b[m") // reset color
		t.Fail()
	}
}

func nok(t testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		// for text color, also check out https://github.com/daviddengcn/go-colortext
		fmt.Print("\x1b[38;5;1m") // text color
		fname := filepath.Base(file)
		fmt.Printf("%s:%d: error expected but not received\n\n",
			append([]interface{}{fname, line})...)
		fmt.Print("\x1b[m") // reset color
		t.Fail()
	}
}

func assert(t testing.TB, cond bool, msg string, v ...interface{}) {
	if !cond {
		_, file, line, _ := runtime.Caller(1)
		// for text color, also check out https://github.com/daviddengcn/go-colortext
		fmt.Print("\x1b[38;5;1m") // text color
		fname := filepath.Base(file)
		fmt.Printf("%s:%d: "+msg+"\n\n",
			append([]interface{}{fname, line}, v...)...)
		fmt.Print("\x1b[m") // reset color
		t.Fail()
	}
}

func equals(tb testing.TB, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Print("\x1b[38;5;1m") // text color
		fname := filepath.Base(file)
		fmt.Printf("%s:%d:\n\texp: %#v\n\tact: %#v\n\n",
			append([]interface{}{fname, line}, exp, act)...)
		fmt.Print("\x1b[m") // reset color
		tb.Fail()
	}
}
