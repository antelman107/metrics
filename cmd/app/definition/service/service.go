// Package service provide dependency injection definitions.
package service

import "github.com/antelman107/metrics/cmd/app/service"

// DefServiceTag service tag name.
const DefServiceTag = "service"

// Service type alias of service.Service.
type Service = service.Service
