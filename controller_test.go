package thermostat

import "testing"

func TestThermostat_OnStartup_EverythingTurnedOff(t *testing.T) {
	_TestThermostat(t, AssertAllOff())
}
func TestThermostat_WhenTooCold_BlowerAndHeaterEngaged(t *testing.T) {
	_TestThermostat(t, MakeItTooCold(), AssertHeating(), AssertBlowing())
}

func MakeItTooCold() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		this.gauge.temperature = 70 - 5 - 1
		this.controller.Regulate()
	}
}
func AssertBlowing() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		if !this.hvac.IsBlowing() {
			this.Error("Not blowing!!")
		}
	}
}
func AssertHeating() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		if !this.hvac.IsHeating() {
			this.Error("Not heating!!")
		}
	}
}
func AssertAllOff() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		if this.hvac.IsBlowing() {
			this.Error("Blowing!!")
		}
		if this.hvac.IsCooling() {
			this.Error("Cooling!!")
		}
		if this.hvac.IsHeating() {
			this.Error("Heating!!")
		}
	}
}

func _TestThermostat(t *testing.T, options ...ThermostatFixtureOption) {
	t.Helper()
	t.Parallel()

	hvac := NewFakeHVAC()
	gauge := NewFakeGauge()
	controller := NewController(hvac, gauge)
	fixture := &ThermostatFixture{
		T:          t,
		hvac:       hvac,
		gauge:      gauge,
		controller: controller,
	}

	for _, option := range options {
		option(fixture)
	}
}

type (
	ThermostatFixtureOption func(this *ThermostatFixture)

	ThermostatFixture struct {
		*testing.T
		hvac       *FakeHVAC
		gauge      *FakeGauge
		controller *Controller
	}
)
