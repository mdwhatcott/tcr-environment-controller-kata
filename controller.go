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

func (this *Controller) Regulate() {
	temperature := this.gauge.CurrentTemperature()

	if temperature < 70-5 {
		this.hvac.SetBlower(true)
		this.hvac.SetHeater(true)
	}
}
