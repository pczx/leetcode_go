package solution45

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	res := []string{}
	for _, v := range nums {
		res = append(res, strconv.Itoa(v))
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i]+res[j] < res[j]+res[i]
	})
	return strings.Join(res, "")
}
