package not_allowed_to_use_tc_internal_packages

import "example.com/banana/teamcity/internal/logger"

func da() {
	logger.Log.Println("opa")
}
