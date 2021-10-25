# run
$ gop run error-handling.gop
add("100", "23"): 123
add("10", "abc"): 0 strconv.Atoi: parsing "abc": invalid syntax
addSafe("10", "abc"): 10
