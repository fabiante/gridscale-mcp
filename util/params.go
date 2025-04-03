package util

import "fmt"

func GetIntParam(args map[string]any, key string) (int, error) {
	if value, ok := args[key]; ok {
		if intValue, ok := value.(int); ok {
			return intValue, nil
		} else if floatValue, ok := value.(float64); ok {
			// FIXME: This is a workaround for the fact that mcp does not support int parameters. This should be removed asap
			return int(floatValue), nil
		}
	}
	return 0, fmt.Errorf("parameter %s not found or not an integer", key)
}
