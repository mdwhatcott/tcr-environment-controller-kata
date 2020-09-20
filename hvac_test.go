package thermostat

type FakeHVAC struct {
	blowing bool
	cooling bool
	heating bool
}

func NewFakeHVAC() *FakeHVAC {
	return &FakeHVAC{
		blowing: true,
		cooling: true,
		heating: true,
	}
}

func (this *FakeHVAC) SetBlower(state bool) { this.blowing = state }
func (this *FakeHVAC) SetCooler(state bool) { this.cooling = state }
func (this *FakeHVAC) SetHeater(state bool) { this.heating = state }

func (this *FakeHVAC) IsBlowing() bool { return this.blowing }
func (this *FakeHVAC) IsCooling() bool { return this.cooling }
func (this *FakeHVAC) IsHeating() bool { return this.heating }

func (this *FakeHVAC) String() string {
	result := ""
	if this.blowing {
		result += "BLOWING "
	} else {
		result += "blowing "
	}
	if this.cooling {
		result += "COOLING "
	} else {
		result += "cooling "
	}
	if this.heating {
		result += "HEATING"
	} else {
		result += "heating"
	}
	return result
}

type FakeGauge struct {
	temperature int
}

func NewFakeGauge() *FakeGauge {
	return &FakeGauge{}
}

func (this *FakeGauge) CurrentTemperature() int {
	return this.temperature
}
