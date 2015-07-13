package main

import (
	"fmt"

	"github.com/dustywilson/dsadmin"
)

func main() {
	var fun dsadmin.AccountFlags
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
	fun.Set(dsadmin.AFNormalAccount)
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
	fun.Set(dsadmin.AFAccountDisable)
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
	fun.Unset(dsadmin.AFNormalAccount)
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
	fun.Unset(dsadmin.AFAccountDisable)
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
	fun.Set(dsadmin.AFNormalAccount)
	fmt.Println(fun.String())
	fmt.Println(fun.Has(dsadmin.AFNormalAccount), "\n")
}
