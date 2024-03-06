package main

import (
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/cmd/api"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
)

func main() {
	pkg.InitConfig()

	database.InitDBs(pkg.ConfigVars.TaikoDBPath, pkg.ConfigVars.AuthDBPath)
	api.Run(pkg.ConfigVars.Port, pkg.ConfigVars.DistPath)
}
