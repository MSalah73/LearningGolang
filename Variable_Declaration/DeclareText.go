package main

import "fmt"

/*
######### Declairations ##########
Declairing Multiple Variable At Once
var V1. V2, V3 type
Multiple Variable Declairations with Initializations
var V1, V2, V3 type = E1, E2, E3
Using Short Hand Notaion
V1, V2, V3 := E1, E2, E3
*/
var lightSwitchIsOn bool = false

func main() {
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
}

func printLightSwitchState() {
	fmt.Println("The light switch is off:", lightSwitchIsOn)
}

func toggleLightSwitch() {
	lightSwitchIsOn = !lightSwitchIsOn
}
