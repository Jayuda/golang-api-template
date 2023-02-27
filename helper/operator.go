package helper

import (
	"errors"
)

func OperatorQuery(symbol string) (string, error) {
	var operator string
	switch symbol {
	case "eq":
		operator = "="
		break
	case "like":
		operator = "like"
		break
	case "lt":
		operator = "<"
		break
	case "lte":
		operator = "<="
		break
	case "gt":
		operator = ">"
		break
	case "gte":
		operator = ">="
		break
	default:
		return operator, errors.New("operator symbol paramater is not valid")
	}
	return operator, nil
}
