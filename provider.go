package PackageB

import "github.com/ChinmayR/PackageD"

func FuncInPackageB() string {
	return "From PackageB: " + PackageD.FuncInPackageD()
}
