package main

import (
	fmt "fmt"
	goptest "github.com/goplus/gop/ast/goptest"
	gopq "github.com/goplus/gop/ast/gopq"
)

func main() {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:5
	script := `
    import (
        gio "io"
    )

    func New() (*Bar, error) {
        return nil, gio.EOF
    }

    bar, err := New()
    if err != nil {
        log.Println(err)
    }
`
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
	doc := func() (_gop_ret gopq.NodeSet) {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
		var _gop_err error
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
		_gop_ret, _gop_err = goptest.New(script)
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
		if _gop_err != nil {
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
			panic(_gop_err)
		}
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:20
		return
	}()
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:22
	fmt.Println(doc.Any().FuncDecl().Name())
//line /home/cy/projects/open-source-projects/golang/gop/tutorial/36-Auto-Property/autoprop.gop:23
	fmt.Println(doc.Any().ImportSpec().Name())
}
