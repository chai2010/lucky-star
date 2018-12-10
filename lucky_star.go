// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package lucky_star

import (
	"hash/crc32"
	"sort"
)

func MakeLuckyStarFromList(seed string, luckyToken ...[2]string) string {
	luckyTokenMap := make(map[string]string)
	for _, v := range luckyToken {
		if v[1] != "" {
			luckyTokenMap[v[0]] = v[1]
		} else {
			luckyTokenMap[v[0]] = v[0]
		}
	}
	return MakeLuckyStar(seed, luckyTokenMap)
}

func MakeLuckyStar(seed string, luckyTokenMap map[string]string) string {
	var (
		crc32Sum   = uint64(crc32.ChecksumIEEE([]byte(seed)))
		sortedKeys = make([]string, 0, len(luckyTokenMap))
	)

	for id, tok := range luckyTokenMap {
		crc32Sum += uint64(crc32.ChecksumIEEE([]byte(tok)))
		sortedKeys = append(sortedKeys, id)
	}
	sort.Strings(sortedKeys)

	idx := int(crc32Sum) % len(luckyTokenMap)
	return sortedKeys[idx]
}
