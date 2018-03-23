package config

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

var qs = []*survey.Question{
	{
		Name: "Path",
		Prompt: &survey.Input{
			Message: "Path to robot?",
			Default: "/Robot",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "Port",
		Prompt: &survey.Input{
			Message: "UDP port?",
			Default: "8765",
		},
		Validate: survey.Required,
	},
	{
		Name: "Serial",
		Prompt: &survey.Input{
			Message: "Serial port?",
			Default: "/dev/ttyACM0",
		},
		Validate: survey.Required,
	},
}

var servoQs = []*survey.Question{
	{
		Name: "Path",
		Prompt: &survey.Input{
			Message: "Path to bodypart? ex. \"/Leg/Toe\"",
		},
		Validate: survey.Required,
	},
	{
		Name: "ServoID",
		Prompt: &survey.Input{
			Message: "ID of the Servo?",
		},
		Validate: survey.Required,
	},
	//		ReturnDelayTime
	{
		Name: "CwAngleLimit",
		Prompt: &survey.Input{
			Message: "Clockwise angle limit? Set to \"0\" for continuous",
			Default: "1023",
		},
		Validate: survey.Required,
	},
	{
		Name: "CcwAngleLimit",
		Prompt: &survey.Input{
			Message: "Counter clockwise angle limit? Set to \"0\" for continuous",
			Default: "0",
		},
		Validate: survey.Required,
	},
	{
		Name: "HighestLimitTemperature",
		Prompt: &survey.Input{Message: "HighestLimitTemperature",
			Default: "70",
		},
		Validate: survey.Required,
	},
	{
		Name: "LowestLimitVoltage",
		Prompt: &survey.Input{Message: "LowestLimitVoltage",
			Default: "60",
		},
		Validate: survey.Required,
	},

	/*


		HighestLimitVoltage
		MaxTorque
		StatusReturnLevel
		AlarmLed
		AlarmShutdown
		TorqueEnable
		Led
		CwComplianceMargin
		CcwComplianceMargin
		CwComplianceSlope
		CcwComplianceSlope
		GoalPosition
		MovingSpeed
		TorqueLimit
		PresentPosition
		PresentSpeed
		PresentLoad
		PresentVoltage
		PresentTemperature
		RegisteredInstruction
		Moving
		Lock
		Punch
	*/
}
