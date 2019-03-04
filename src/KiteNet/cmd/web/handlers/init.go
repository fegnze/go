package handlers

import (
	"KiteNet/cmd/web/handlers/account"
	"KiteNet/cmd/web/handlers/region"
	"KiteNet/cmd/web/handlers/tourist"
	"KiteNet/httpkt"
)

//RoutMap -
var RoutMap = httpkt.RoutMap{
	"/worldship/version": VersionHandler,

	"/worldship/stepAnalytics": StepAnalyticsHandler,

	"/worldship/region/find": region.FindHandler,
	"/worldship/region/get": region.GetHandler,

	"/worldship/mix/login": account.MixLoginHandler,
	"/worldship/mix/account/bind": account.BindHandler,
	"/worldship/mix/account/register": account.RegisterHandler,
	"/worldship/mix/login/tourist/get": tourist.GetTouristHandler,
}
