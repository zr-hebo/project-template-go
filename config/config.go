package config

import "github.com/zr-hebo/utils/ezflag"

const (
	ProjectName = "template-go"
)

var (
	WebServicePort = ezflag.NewIntVar(
		"web_service_port", 7088,
		"web service port, default is 7088. It is not required.", false)
)
