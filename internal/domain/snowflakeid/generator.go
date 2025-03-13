package snowflakeid

import (
	"log/slog"
	"sync/atomic"
	"time"
)

type TypeToGenerate int32

const (
	Text TypeToGenerate = iota
	Base64
	Binary
)

const (
	maxLocalCounter = 4_096_000
	maxMachineId    = 1_024

	firstBit = 0b0111111111111111111111111111111111111111111111111111111111111111
)

type MachineIdMock struct {
}

func (MachineIdMock) Get() int64 {
	return 10
}

type MachineId interface {
	Get() int64
}

type Generator struct {
	machineId    MachineId
	startEpoch   time.Time
	localCounter atomic.Int64

	isSetToReset bool
}

func NewGenerator(machine MachineId, startEpoch time.Time) *Generator {
	return &Generator{
		machineId:    machine,
		startEpoch:   startEpoch,
		localCounter: atomic.Int64{},
		isSetToReset: false,
	}
}

func (g *Generator) Generate(typeToGenerate TypeToGenerate) (Model, error) {
	if !g.isSetToReset {
		g.SetTimerToResetLocalVariable()
	}

	return g.generate()
}

func (g *Generator) generate() (Model, error) {
	var snowflakeid int64

	now := time.Since(g.startEpoch).Milliseconds()

	machineId := g.machineId.Get()
	if machineId > maxMachineId {
		return Model{}, ErrMachineIdOverflow
	}

	value := g.localCounter.Add(1)
	if value > maxLocalCounter {
		return Model{}, ErrLocalCounterOverflow
	}

	slog.Debug("Basic value generate", slog.Int64("LocalCounter", value), slog.Int64("MachineId", machineId), slog.Int64("now", now))

	now = now << 22
	machineId = machineId << 12
	snowflakeid = now + machineId + value
	snowflakeid = snowflakeid & firstBit

	slog.Debug("Generate id with values", slog.Int64("Id", snowflakeid), slog.Int64("LocalCounter", value), slog.Int64("MachineId", machineId), slog.Int64("now", now))

	return Model{
		Base: snowflakeid,
	}, nil
}

func (g *Generator) SetTimerToResetLocalVariable() {
	if g.isSetToReset {
		return
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond)
		for range ticker.C {
			g.localCounter.Store(0)
		}
	}()

	g.isSetToReset = true
}
