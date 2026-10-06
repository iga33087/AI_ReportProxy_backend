package lib

import (
	"fmt"
	"regexp"
	"encoding/json"
)

// ToRegexp 接收明確的 map[string]*ConvertItem 型態
func ToRegexp(data map[string]any, table map[string]*ConvertItem) map[string]any {
	// 1. 安全解析大模型返回的 content (保持原本的安全檢查)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil // 序列化失敗
	}
	
	var copiedData map[string]any
	if err := json.Unmarshal(jsonBytes, &copiedData); err != nil {
		return nil // 反序列化失敗
	}

	choices, ok := copiedData["choices"].([]any)
	if !ok || len(choices) == 0 {
		return nil
	}
	firstChoice, ok := choices[0].(map[string]any)
	if !ok {
		return nil
	}
	message, ok := firstChoice["message"].(map[string]any)
	if !ok {
		return nil
	}
	content, ok := message["content"].(string)
	if !ok {
		return nil
	}

	var result string = content

	// 2. 因為型態明確，直接遍歷 table，此時 item 的型態直接就是 *ConvertItem
	for _, item := range table {
		if item == nil {
			continue
		}

		// 3. 直接存取 RestoreTable，不需任何型態斷言
		// 這裡假設 RestoreTable 的型態是 map[string]any 或 map[string]string
		for index2, val2 := range item.RestoreTable {
			replStr := fmt.Sprintf("%v", val2)
			re := regexp.MustCompile(regexp.QuoteMeta(index2))
			
			// 使用累積的 result 進行取代
			result = re.ReplaceAllString(result, replStr)
		}
	}

    message["content"] = result

	return copiedData
}