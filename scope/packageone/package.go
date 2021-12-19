package packageone

import "fmt"

var PackageVar = "Hello, I'm a package level variable in the package.packageone =^.^= "

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}
