package main

var _headerTemplate = `
// Code generated by kratos tool genbts. DO NOT EDIT.

NEWLINE
/* 
  Package {{.PkgName}} is a generated cache proxy package.
  It is generated from:
  ARGS
*/
NEWLINE

package {{.PkgName}}

import (
	"context"
	{{if .EnableBatch }}"sync"{{end}}
NEWLINE
	"github.com/bilibili/kratos/pkg/stat/metric"
	{{if .EnableBatch }}"github.com/bilibili/kratos/pkg/sync/errgroup"{{end}}
	{{.ImportPackage}}
NEWLINE
	{{if .EnableSingleFlight}}	"golang.org/x/sync/singleflight" {{end}}
)

var (
	_ _bts
	_metricHits = metric.NewBusinessMetricCount("hits_total", "name")
	_metricMisses = metric.NewBusinessMetricCount("misses_total", "name")
)
{{if .EnableSingleFlight}}
var cacheSingleFlights = [SFCOUNT]*singleflight.Group{SFINIT} 
{{end }}
`
