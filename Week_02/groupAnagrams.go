package main

import (
	"fmt"
	"sort"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
type bytes []byte

func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := make(map[string]int) //结构是：[字符串]ret的下标，用于输出题解规定的结构
	for _, str := range strs {
		kByte := bytes(str) //将字符串转换成切片
		sort.Slice(kByte, func(i, j int) bool { return kByte[i] < kByte[j] })
		k := string(kByte)
		if idx, ok := m[k]; !ok {
			m[k] = len(ret) // 记录拍完序的字符串以及字符串在ret的位置
			ret = append(ret, []string{str})
		} else { // 已经出现过，将其放入出现在ret中，在ret的位置为idx
			ret[idx] = append(ret[idx], str)
		}
	}
	return ret
}

func main3() {
	ss := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	ret := groupAnagrams(ss)

	fmt.Println(ret)

}
