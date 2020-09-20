package thermostat

type Gauge interface {
	CurrentTemperature() int // Current ambient temperature rounded to the nearest degree (Fahrenheit).
}

type HVAC interface {
	SetBlower(state bool) // Turns the blower on or off.
	SetCooler(state bool) // Turns the cooler on or off.
	SetHeater(state bool) // Turns the heater on or off.

	IsBlowing() bool // Is the blower currently on or off?
	IsCooling() bool // Is the cooler currently on or off?
	IsHeating() bool // Is the heater currently on or off?
}
