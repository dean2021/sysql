//go:build darwin && !cgo
// +build darwin,!cgo

package host

import (
	"context"
	"github.com/dean2021/sysql/extend/tables/common"
)

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
