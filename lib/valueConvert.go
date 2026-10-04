package lib

import (
	"strconv"
)

type ConvertItem struct {
	AltText string
	Total   int
	Table   map[string]string
	RestoreTable   map[string]string
}

func ValueConvert(data []map[string]any) ([]map[string]any,map[string]*ConvertItem) {
	resultList := make([]map[string]any, 0, len(data))

	var convertTable = map[string]*ConvertItem{
		"user":         {AltText: "USER", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"group":        {AltText: "GROUP", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service":      {AltText: "SERVICE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service_type": {AltText: "SERVICE_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"domain":       {AltText: "DOMAIN", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"web_type":     {AltText: "WEB_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
	}

	for _, val := range data {
		resultRow := make(map[string]any)

		for j, val2 := range val {
			valStr, isString := val2.(string)
			if !isString {
				resultRow[j] = val2
				continue
			}
			if item, ok := convertTable[j]; ok {
				if _, ok2 := item.Table[valStr]; !ok2 {
					item.Total += 1
					altText := item.AltText + strconv.Itoa(item.Total)
					item.Table[valStr] = altText
					item.RestoreTable[altText] = valStr
				}
				resultRow[j] = item.Table[valStr]
			} else {
				resultRow[j] = val2 
			}
		}

		resultList = append(resultList, resultRow)
	}

	return resultList,convertTable
}