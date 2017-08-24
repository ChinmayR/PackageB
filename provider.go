package PackageB

import "github.com/ChinmayR/PackageD"

func FuncInPackageB() string {
	str, _ := PackageD.FuncInPackageD()
	return "From PackageB: " + str
}
