package thermostat

type Controller struct {
	hvac  HVAC
	gauge Gauge

	blowerDelay int
	coolerDelay int
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
	this.decrementDelays()

	switch this.temperature() {
	case TooCold:
		this.heat()
	case TooHot:
		this.cool()
	default:
		this.idle()
	}
}

func (this *Controller) decrementDelays() {
	if this.blowerDelay > 0 {
		this.blowerDelay--
	}
	if this.coolerDelay > 0 {
		this.coolerDelay--
	}
}

func (this *Controller) heat() {
	this.engageBlower()
	this.disengageCooler()
	this.engageHeater()
}
func (this *Controller) cool() {
	this.engageBlower()
	this.engageCooler()
	this.disengageHeater()
}
func (this *Controller) idle() {
	this.disengageBlower()
	this.disengageCooler()
	this.disengageHeater()
}

func (this *Controller) engageCooler() {
	if this.coolerDelay == 0 {
		this.hvac.SetCooler(true)
	}
}
func (this *Controller) engageBlower() {
	this.hvac.SetBlower(true)
}
func (this *Controller) engageHeater() {
	this.blowerDelay = 5 + 1
	this.hvac.SetHeater(true)
}

func (this *Controller) disengageHeater() {
	this.hvac.SetHeater(false)
}
func (this *Controller) disengageCooler() {
	if this.hvac.IsCooling() {
		this.coolerDelay = 3
	}
	this.hvac.SetCooler(false)
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
