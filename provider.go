package PackageB

import "github.com/ChinmayR/PackageD"

func FuncInPackageB() string {
	dString, _ := PackageD.FuncInPackageD()
	return "From PackageB: " + dString
}
