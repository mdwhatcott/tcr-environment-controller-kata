package thermostat

type FakeHVAC struct {
	blowing bool
	cooling bool
	heating bool
}

func NewFakeHVAC() *FakeHVAC {
	return &FakeHVAC{}
}

func (this *FakeHVAC) SetBlower(state bool) { this.blowing = state }
func (this *FakeHVAC) SetCooler(state bool) { this.cooling = state }
func (this *FakeHVAC) SetHeater(state bool) { this.heating = state }

func (this *FakeHVAC) IsBlowing() bool { return this.blowing }
func (this *FakeHVAC) IsCooling() bool { return this.cooling }
func (this *FakeHVAC) IsHeating() bool { return this.heating }

type FakeGauge struct {
	temperature int
}

func NewFakeGauge() *FakeGauge {
	return &FakeGauge{}
}

func (this *FakeGauge) CurrentTemperature() int {
	return this.temperature
}
