package packageone

var privateVar = "I'm a private var"
var PublicVar = "I'm public (or exported)"

// In GoLang variables starting with capital letter as first letter, means it's exported, it's available outside this package
// privateVar - with Lower case is available only inside this package
