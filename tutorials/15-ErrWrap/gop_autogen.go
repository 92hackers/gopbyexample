package main

import (
	fmt "fmt"
	strconv "strconv"
)

func add(x string, y string) (int, error) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	var _autoGo_1 int
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	{
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		_autoGo_1, _gop_err = strconv.Atoi(x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
			return 0, _gop_err
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		goto _autoGo_2
	_autoGo_2:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	var _autoGo_3 int
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	{
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		_autoGo_3, _gop_err = strconv.Atoi(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
			return 0, _gop_err
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
		goto _autoGo_4
	_autoGo_4:
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:6
	return _autoGo_1 + _autoGo_3, nil
}
func addSafe(x string, y string) int {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
	return func() (_gop_ret int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		_gop_ret, _gop_err = strconv.Atoi(x)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
			return 0
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		return
	}() + func() (_gop_ret int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		_gop_ret, _gop_err = strconv.Atoi(y)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
			return 0
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:10
		return
	}()
}
func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
	fmt.Println(`add("100", "23"):`, func() (_gop_ret int) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
		_gop_ret, _gop_err = add("100", "23")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
			panic(_gop_err)
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:13
		return
	}())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:15
	sum, err := add("10", "abc")
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:16
	fmt.Println(`add("10", "abc"):`, sum, err)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/15-ErrWrap/err_wrap.gop:18
	fmt.Println(`addSafe("10", "abc"):`, addSafe("10", "abc"))
}
