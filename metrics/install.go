package metrics

import (
	"github.com/PMoneda/telemetry/registry"
	"github.com/mundipagg/boleto-api/config"
)

//Install Inicia telemetria
func Install() {
	if config.Get().EnableMetrics {
		cnf := registry.Config{}
		cnf.Host = config.Get().InfluxDBHost
		cnf.Port = config.Get().InfluxDBPort
		InstallRuntime(cnf)
		InstallTimingMetrics(cnf)
		InstallBusinessMetrics(cnf)
	}
}
