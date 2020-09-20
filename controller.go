package thermostat

type Controller struct {
	hvac  HVAC
	gauge Gauge
}

func NewController(hvac HVAC, gauge Gauge) *Controller {
	hvac.SetBlower(false)
	hvac.SetCooler(false)
	hvac.SetHeater(false)
	return &Controller{hvac: hvac, gauge: gauge}
}
