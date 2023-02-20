package app

import (
	v1 "bm-support/routes/v1"
)

func setUpRoutes() {
	// Routing v1
	routerV1 := router.Group("api/v1/")

	// Auth Routing
	v1.SetupAuthRoute(routerV1)

	// User Routing
	v1.SetupUsersRoute(routerV1)
}
