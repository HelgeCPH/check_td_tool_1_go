// It seems as if Code Climate counts SLOC. For this file it reports 8 lines of
// code.
//
// This code violates the first maintainability check `argument-count`
package main

import ("fmt"; "strings")


func print_it(a string, b string, c string, d string, e string, f string, g string) {
    fmt.Println(a, b, c, d, e, f, g)
}

func main() {
	parts := strings.Split("Hej Code Climate, please check this code.", " ")
    a, b, c, d, e, f, g := parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6]
    print_it(a, b, c, d, e, f, g)
}
