package thermostat

type Controller struct {
	hvac  HVAC
	gauge Gauge
}

func NewController(hvac HVAC, gauge Gauge) *Controller {
	hvac.SetBlower(false)
	hvac.SetCooler(false)
	hvac.SetHeater(false)
	return &Controller{
		hvac:  hvac,
		gauge: gauge,
	}
}

func (this *Controller) Regulate() {
	switch this.temperature() {
	case TooCold:
		this.hvac.SetBlower(true)
		this.hvac.SetCooler(false)
		this.hvac.SetHeater(true)
	case TooHot:
		this.hvac.SetBlower(true)
		this.hvac.SetCooler(true)
		this.hvac.SetHeater(false)
	}
}

func (this *Controller) temperature() EnvironmentState {
	temperature := this.gauge.CurrentTemperature()

	if temperature < IdealTemperature-AllowedTolerance {
		return TooCold
	} else if temperature > IdealTemperature+AllowedTolerance {
		return TooHot
	} else {
		return Comfy
	}
}

type EnvironmentState int

const (
	TooCold = iota
	TooHot
	Comfy
)

const AllowedTolerance = 5
const IdealTemperature = 70
