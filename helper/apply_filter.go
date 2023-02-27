package helper

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplyFilter(tx *gorm.DB, filters *map[string]string) error {
	for key, value := range *filters {
		keySplitted := strings.Split(key, ".")
		column := keySplitted[0]
		operator, err := OperatorQuery(keySplitted[1])
		if err != nil {
			return err
		}
		query := fmt.Sprintf("%s %s ?", column, operator)
		if operator == "like" {
			value = "%" + value + "%"
		}
		tx = tx.Where(query, value)
	}
	return nil
}
