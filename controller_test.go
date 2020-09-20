package thermostat

import "testing"

func TestThermostat_OnStartup_EverythingTurnedOff(t *testing.T) {
	_TestThermostat(t, AssertAllOff())
}
func TestThermostat_WhenTooCold_BlowerAndHeaterEngaged(t *testing.T) {
	_TestThermostat(t, MakeItTooCold(), AssertHeating())
}
func TestThermostat_WhenTooHot_BlowerAndCoolerEngaged(t *testing.T) {
	_TestThermostat(t, MakeItTooHot(), AssertCooling())
}
func TestThermostat_WhenComfy_AllOff(t *testing.T) {
	_TestThermostat(t, MakeItComfy(), AssertAllOff())
}
func TestThermostat_WhenTooHotThenTooCold_BlowerAndHeaterEngaged(t *testing.T) {
	_TestThermostat(t, MakeItTooHot(), MakeItTooCold(), AssertHeating())
}
func TestThermostat_WhenTooColdThenTooHot_BlowerAndCoolerEngaged(t *testing.T) {
	_TestThermostat(t, MakeItTooCold(), MakeItTooHot(), AssertCooling())
}
func TestThermostat_WhenTooHotThenComfy_EverythingOff(t *testing.T) {
	_TestThermostat(t, MakeItTooHot(), MakeItComfy(), AssertAllOff())
}
func TestThermostat_WhenTooColdThenComfy_BlowerShouldRemainOn(t *testing.T) {
	_TestThermostat(t,
		MakeItTooCold(),
		MakeItComfy(), AssertBlowing(),
		MakeItComfy(), AssertBlowing(),
		MakeItComfy(), AssertBlowing(),
		MakeItComfy(), AssertBlowing(),
		MakeItComfy(), AssertBlowing(),
		MakeItComfy(), AssertAllOff(),
	)
}
func TestThermostat_WhenTooHotThenComfyThenTooHotAgain_CoolerOnlyReengagesAfterDelay(t *testing.T) {
	_TestThermostat(t,
		MakeItTooHot(),
		MakeItComfy(),
		MakeItTooHot(), AssertBlowing(),
		MakeItTooHot(), AssertBlowing(),
		MakeItTooHot(), AssertCooling(),
	)
}
func TestThermostat_WhenTooHotThenComfyThenTooHotThenComfy_CoolerStaysOff(t *testing.T) {
	_TestThermostat(t,
		MakeItTooHot(),
		MakeItComfy(),
		MakeItTooHot(),
		MakeItComfy(), AssertAllOff(),
		MakeItComfy(), AssertAllOff(),
		MakeItComfy(), AssertAllOff(),
	)
}

func MakeItComfy() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		this.gauge.temperature = 70
		this.controller.Regulate()
	}
}
func MakeItTooHot() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		this.gauge.temperature = 70 + 5 + 1
		this.controller.Regulate()
	}
}
func MakeItTooCold() ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		this.gauge.temperature = 70 - 5 - 1
		this.controller.Regulate()
	}
}
func AssertBlowing() ThermostatFixtureOption { return AssertHVACState("BLOWING cooling heating") }
func AssertCooling() ThermostatFixtureOption { return AssertHVACState("BLOWING COOLING heating") }
func AssertHeating() ThermostatFixtureOption { return AssertHVACState("BLOWING cooling HEATING") }
func AssertAllOff() ThermostatFixtureOption  { return AssertHVACState("blowing cooling heating") }
func AssertHVACState(expected string) ThermostatFixtureOption {
	return func(this *ThermostatFixture) {
		state := this.hvac.String()
		if state == expected {
			return
		}
		this.Errorf("\n"+
			"Expected: %s\n"+
			"Actual:   %s",
			expected,
			state,
		)
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
