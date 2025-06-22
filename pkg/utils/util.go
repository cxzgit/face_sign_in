package utils

import (
	"fmt"
	"strconv"
)

// ChangeStringToUint
// @Description: 将string类型的值转换成uint类型
// @Author wangyulong 2024-10-09 15:26:15
// @param        str string
// @return       uint
// @return       error
func ChangeStringToUint(str string) (uint, error) {
	// 将字符串转换成uint64, 基数为 10, 位大小为 64 位
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("ChangeType -> 将string类型的值转换成uint类型失败 -> %s", err)
	}

	// 将 uint64 转换成 uint 类型
	uintNum := uint(num)

	return uintNum, nil
}
