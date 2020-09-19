# environment-controller-kata

An implementation of Robert "Uncle Bob" Martin's Environment Controller kata in Go.

- [Original presentation](https://vimeo.com/71816368) 
- [Source code](https://github.com/unclebob/environmentcontroller)


(page 1/6, scroll for more...)





































## The scenario

You've been tasked with writing software for a thermostat that will control
HVAC hardware (heater, blower, air-conditioner). This hardware must be operated
according to the manufacturer's specifications in order to run safely and effectively.

To effectively regulate the environment the temperature should never deviate more than
5 degrees from 70F.

(page 2/6)


































## Bad news:
    
The various hardware devices have not yet arrived, so we can't connect directly to them yet.


## Good news: 

We have the hardware documentation! Here are the included hardware interfaces:

------

    package contracts
    
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
    
------

On and off are analogous to the go values `true` and `false`, respectively.


(page 3/6)




















## Rules specified by hardware manufacturer:

1. When the heater is on, the blower must also be on.
2. When the cooler is on, the blower must also be on.
3. The blower must continue to run for 5 minutes after the heater turns off.
    - Because the heat exchanger has latent heat that must not be allowed to
    accumulate lest it melt the sensitive vanes.
4. The cooler must not be turned on within 3 minutes of being turned off.
    - Because the freon must be given time to re-condense lest the compressor vapor lock.


(page 4/6)


































## Design Questions:

2. How can we test-drive the thermostat without the actual hardware drivers?
1. How will we mark the passage of time?


(page 5/6)







































## To-do Items:

Basic Logic:

1. [ ] on startup, everything is off
2. [ ] when too cold, turn on blower and heater
4. [ ] when too hot, turn on blower and cooler
3. [ ] when comfy, nothing turns on

Intermediate Logic:

1. [ ] when too cold then too hot, blower and cooler on
2. [ ] when too hot then too cold, blower and heater on
3. [ ] when too hot then comfy, everything off
4. [ ] when too cold then comfy, blower remains on

Advanced Logic:

1. [ ] when too cold then comfy, blower remains on for 5 minutes before turning off
2. [ ] when too hot then comfy then too hot, cooler must not turn on for 3 minutes even if it gets too hot
3. [ ] when too hot then comfy, cooler does not turn back on after delay if not needed


(THE END)


