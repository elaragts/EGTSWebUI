package main

import (
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/cmd/api"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/internal/database"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/pkg"
	"github.com/keitannunes/KeifunsTaikoWebUI/backend/updater"
)

func main() {
	pkg.InitConfig()
	updater.InitConfig()
	database.InitDBs(pkg.ConfigVars.TaikoDBPath, pkg.ConfigVars.AuthDBPath)
	pkg.InitDatatable(pkg.ConfigVars.DatatablePath)
	api.Run(pkg.ConfigVars.Port, pkg.ConfigVars.LegacyWebUIURL, pkg.ConfigVars.DistPath)
}
