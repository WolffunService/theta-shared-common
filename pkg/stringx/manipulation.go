package stringx

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase returns snake_case of the provided value.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// KebabCasePath format path : ranking-reward/user-rank
//   - Ex method name : UserRanking => user-ranking
//   - Ex next path: method name : UserRanking_RankingReward => user-ranking/ranking-reward
//
// Gắn vào SetCustomPathWordFunc ở cái đoạn mà gắn Hanlde controller mong muốn á
// rồi đặt tên hàm sao cho giống ví dụ trên là được
func KebabCasePath(path, w string, wordIndex int) string {
	const suffixNextPath = "_"

	var nextPath = false

	w = strings.ToLower(w)
	if wordIndex == 0 {
		//reset when start
		nextPath = false
	} else {
		if nextPath {
			//reset
			nextPath = false
		} else {
			w = "-" + w
		}
		//"Acbs" or "Abxs_"
		if strings.HasSuffix(w, suffixNextPath) {
			w = strings.Replace(w, suffixNextPath, "/", 1)
			nextPath = true
		}
	}
	path += w
	return path
}
