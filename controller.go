package thermostat

type Controller struct {
	hvac  HVAC
	gauge Gauge

	blowerDelay int
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
	if this.blowerDelay > 0 {
		this.blowerDelay--
	}

	switch this.temperature() {
	case TooCold:
		this.hvac.SetBlower(true)
		this.hvac.SetCooler(false)
		this.engageHeater()
	case TooHot:
		this.hvac.SetBlower(true)
		this.hvac.SetCooler(true)
		this.hvac.SetHeater(false)
	default:
		this.disengageBlower()
		this.hvac.SetCooler(false)
		this.hvac.SetHeater(false)
	}
}

func (this *Controller) engageHeater() {
	this.blowerDelay = 5 + 1
	this.hvac.SetHeater(true)
}

func (this *Controller) disengageBlower() {
	if this.blowerDelay == 0 {
		this.hvac.SetBlower(false)
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
