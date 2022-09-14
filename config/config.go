package config

import (
	//"github.com/adammck/dynamixel/servo"

	"io/ioutil"

	"github.com/adammck/dynamixel/servo"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Path              string   `json:"path",yaml:"path"`
	Port              string   `json:"port",yaml:"port"`
	Serial            string   `json:"serial",yaml:"serial"`
	Servos            []*Servo `json:"servos",yaml:"servos"`
	StatusReturnLevel int      `json:"statusreturnlevel",yaml:"statusreturnlevel"`
}
type Servo struct {
	Servo                   *servo.Servo
	Path                    string `json:"path",yaml:"path"`
	ServoID                 int    `json:"servoid",yaml:"servoid"`
	Baudrate                int    `json:"baudrate",yaml:"baudrate"`
	ReturnDelayTime         int    `json:"returndelaytime",yaml:"returndelaytime"`
	CwAngleLimit            int    `json:"cwanglelimit",yaml:"cwanglelimit"`
	CcwAngleLimit           int    `json:"ccwanglelimit",yaml:"ccwanglelimit"`
	HighestLimitTemperature int    `json:"highestlimittemperature",yaml:"highestlimittemperature"`
	LowestLimitVoltage      int    `json:"lowestlimitvoltage",yaml:"lowestlimitvoltage"`
	HighestLimitVoltage     int    `json:"highestlimitvoltage",yaml:"highestlimitvoltage"`
	MaxTorque               int    `json:"maxtorque",yaml:"maxtorque"`
	StatusReturnLevel       int    `json:"statusreturnlevel",yaml:"statusreturnlevel"`
	AlarmLed                int    `json:"alarmled",yaml:"alarmled"`
	AlarmShutdown           int    `json:"alarmshutdown",yaml:"alarmshutdown"`
	TorqueEnable            int    `json:"torqueenable",yaml:"torqueenable"`
	Led                     int    `json:"led",yaml:"led"`
	CwComplianceMargin      int    `json:"cwcompliancemargin",yaml:"cwcompliancemargin"`
	CcwComplianceMargin     int    `json:"ccwcompliancemargin",yaml:"ccwcompliancemargin"`
	CwComplianceSlope       int    `json:"cwcomplianceslope",yaml:"cwcomplianceslope"`
	CcwComplianceSlope      int    `json:"ccwcomplianceslope",yaml:"ccwcomplianceslope"`
	GoalPosition            int    `json:"goalposition",yaml:"goalposition"`
	MovingSpeed             int    `json:"movingspeed",yaml:"movingspeed"`
	TorqueLimit             int    `json:"torquelimit",yaml:"torquelimit"`
	PresentPosition         int    `json:"presentposition",yaml:"presentposition"`
	PresentSpeed            int    `json:"presentspeed",yaml:"presentspeed"`
	PresentLoad             int    `json:"presentload",yaml:"presentload"`
	PresentVoltage          int    `json:"presentvoltage",yaml:"presentvoltage"`
	PresentTemperature      int    `json:"presenttemperature",yaml:"presenttemperature"`
	RegisteredInstruction   int    `json:"registeredinstruction",yaml:"registeredinstruction"`
	Moving                  int    `json:"moving",yaml:"moving"`
	Lock                    int    `json:"lock",yaml:"lock"`
	Punch                   int    `json:"punch",yaml:"punch"`
}

func (c *Config) GetConfig(filename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func PromptMe() (*Config, error) {
	c := &Config{}
	c.Servos = make([]*Servo, 0)

	err := survey.Ask(qs, c)
	if err != nil {
		return nil, err
	}

	more := true

	for more {

		s := &Servo{}

		err := survey.Ask(servoQs, s)
		if err != nil {
			return nil, err
		} else {
			c.Servos = append(c.Servos, s)
		}

		prompt := &survey.Confirm{
			Message: "Add one more?",
		}
		survey.AskOne(prompt, &more, nil)
	}
	save := false
	prompt := &survey.Confirm{
		Message: "Want to save the configuration",
	}

	survey.AskOne(prompt, &save, nil)
	if save {

		ymlConfig, err := yaml.Marshal(c)
		if err != nil {
			return nil, err
		}

		filenameQ := &survey.Input{
			Message: "Filename?",
		}
		var filename string

		survey.AskOne(filenameQ, &filename, nil)

		err = ioutil.WriteFile(filename, ymlConfig, 0644)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
