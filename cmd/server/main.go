package main

import (
	"github.com/nifle3/goarenas-snowflakeClone/internal/config"
	"github.com/nifle3/goarenas-snowflakeClone/internal/domain/snowflakeid"
	"github.com/nifle3/goarenas-snowflakeClone/internal/logger"
	"github.com/nifle3/goarenas-snowflakeClone/internal/server"
)

func main() {
	cfg := config.MustNew()
	logger.MustSetup(cfg.EnvType)

	generator := snowflakeid.NewGenerator(snowflakeid.MachineIdMock{}, cfg.StartEpoch)

	service := snowflakeid.NewService(generator)

	server.MustStart(service)
}
