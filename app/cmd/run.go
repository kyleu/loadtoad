package cmd

import (
	"github.com/kyleu/loadtoad/app"
	"github.com/kyleu/loadtoad/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return _logger, err
	}
	return _logger, nil
}
