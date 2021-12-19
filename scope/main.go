package main

import (
	"myapp/packageone"
)

var myVar = "Hello, I'm a package level variable =^.^= "

func main() {

	var blockVar = "Hello, I'm a block level variable =^.^= "
	packageone.PrintMe(myVar, blockVar)
	//packageone.PrintMe(blockVar)
	//packageone.PrintMe(packageone.PackageVar)
}

/*variables*/
// declare a package level variable fot the main package named myVar
// declare a block variable for the main function called blockVar
// declare a package level variable in the package.packageone named PackageVar
// create an exported function in packageone called PrintMe
// in the main function, print out the values of myVar, blockVar, and PackageVar on one line using the PrintMe function in Packageone
