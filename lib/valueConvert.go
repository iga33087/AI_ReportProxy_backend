package lib

import (
	"strconv"
)

type ConvertItem struct {
	AltText      string
	Total        int
	Table        map[string]string
	RestoreTable map[string]string
}

func ValueConvert(data map[string][]map[string]any) (map[string][]map[string]any, map[string]*ConvertItem) {
	resultMap := make(map[string][]map[string]any)
	var convertTable = map[string]*ConvertItem{
		"user":         {AltText: "USER", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"group":        {AltText: "GROUP", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service":      {AltText: "SERVICE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service_type": {AltText: "SERVICE_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"domain":       {AltText: "DOMAIN", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"web_type":     {AltText: "WEB_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
	}
	for key, list := range data {
		resultList := make([]map[string]any, 0, len(list))
		for _, val := range list {
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
		resultMap[key] = resultList
	}

	return resultMap, convertTable
}

func ValueConvert2(data map[string]map[string]any) (map[string]map[string]any, map[string]*ConvertItem) {
	// 初始化最終回傳的外部 map
	resultMap := make(map[string]map[string]any)

	var convertTable = map[string]*ConvertItem{
		"user":         {AltText: "USER", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"group":        {AltText: "GROUP", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service":      {AltText: "SERVICE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"service_type": {AltText: "SERVICE_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"domain":       {AltText: "DOMAIN", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
		"web_type":     {AltText: "WEB_TYPE", Total: 0, Table: map[string]string{}, RestoreTable: map[string]string{}},
	}

	for key, innerMap := range data {
		
		resultRow := make(map[string]any)

		for j, val2 := range innerMap {
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

		resultMap[key] = resultRow
	}

	return resultMap, convertTable
}