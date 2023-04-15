package util

import (
	"fmt"

	"github.com/dambaquyen96/smartivr-backend-go/pkg/setting"
)

var ExclusiveInputSlot []string = []string{
	"CALLBOT_TYPE",
	"CALLBOT_RASA_ENDPOINT",
}

func GetInputSlots(contentData map[string]interface{}) ([]interface{}, error) {
	defaultInputSlots, err := GetScenarioInputSlots(contentData)
	if err != nil {
		return nil, err
	}

	pos := make([]int, 0)
	for id, inputSlot := range defaultInputSlots {
		parsedInputSlot := inputSlot.(map[string]interface{})
		for _, callbotType := range setting.CallBotRASASetting.ExclusiveInputSlot {
			if parsedInputSlot["key"].(string) == callbotType {
				pos = append(pos, id)
				break
			}
		}
	}
	acceptedInputSlots := make([]interface{}, 0)
	for id, dis := range defaultInputSlots {
		flag := true
		for _, idRemoved := range pos {
			if idRemoved == id {
				flag = false
			}
		}
		if flag {
			acceptedInputSlots = append(acceptedInputSlots, dis)
		}
	}
	acceptedInputSlots = append(acceptedInputSlots, map[string]interface{}{
		"key":   "phone",
		"value": nil,
	})
	return acceptedInputSlots, nil
}

func GetFullInputSlots(contentData map[string]interface{}) ([]interface{}, error) {
	defaultInputSlots, err := GetScenarioInputSlots(contentData)
	if err != nil {
		return nil, err
	}
	return defaultInputSlots, nil
}

func GetScenarioInputSlots(contentData map[string]interface{}) ([]interface{}, error) {
	defaultInputSlots := make([]interface{}, 0)
	for _, state := range contentData {
		stateMap, ok := state.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("GetInputSlots | Error when parsing stateMap")
		}
		stateName, ok := stateMap["type"].(string)
		if !ok {
			return nil, fmt.Errorf("GetInputSlots | Error when parsing stateName")
		}
		if stateName != "trigger" {
			continue
		}
		properties, ok := stateMap["properties"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("GetInputSlots | Error when parsing properties")
		}
		if properties["variables"] == nil {
			return defaultInputSlots, nil
		}
		variables, ok := properties["variables"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("GetInputSlots | Error when parsing variables")
		}
		for _, variable := range variables {
			variableMap, ok := variable.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("GetInputSlots | Error when parsing variableMap")
			}
			defaultInputSlots = append(defaultInputSlots, variableMap)
		}
	}
	return defaultInputSlots, nil
}

func BindDefaultInputSlots(inputSlots []interface{}, campaignInputSlots []interface{}) ([]interface{}, error) {
	appendedInputSlots := make([]interface{}, 0)
	for _, campaignInputSlot := range campaignInputSlots {
		cis := campaignInputSlot.(map[string]interface{})
		isIn := false
		for _, inputSlot := range inputSlots {
			is := inputSlot.(map[string]interface{})
			if is["key"].(string) == cis["key"].(string) {
				isIn = true
			}
		}
		if !isIn {
			appendedInputSlots = append(appendedInputSlots, cis)
		}
	}
	inputSlots = append(inputSlots, appendedInputSlots...)
	return inputSlots, nil
}
