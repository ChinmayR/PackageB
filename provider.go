package PackageB

import "github.com/ChinmayR/PackageD"

func FuncInPackageB() string {
	dOut, _ := PackageD.FuncInPackageD()
	return "From PackageB: " + dOut
}


