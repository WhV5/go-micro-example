/**
* @Author : henry
* @Data: 2020-09-04 10:04
* @Note:
**/

package service

func GetSum(inputs ...int64) (ret int64) {
	for _, v := range inputs {
		ret += v
	}
	return ret
}
