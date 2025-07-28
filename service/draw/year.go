package draw

import (
	"strconv"
	"time"

	gitDraw "github.com/liankui/git-contributions-draw/go/git-contributions-draw"
)

func GetYears() *gitDraw.Year {
	year := time.Now().Year()
	start := getFirstSunday(year)
	end := getLastSaturday(year)
	return &gitDraw.Year{
		Year:  strconv.Itoa(year),
		Total: 100,
		Range: &gitDraw.Range{
			Start: start,
			End:   end,
		},
	}
}

// 返回某年中的第一个周日
func getFirstSunday(year int) string {
	date := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	// 加天数直到找到周日（Sunday == 0）
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}
	return date.Format("2006-01-02")
}

// 返回某年中的最后一个周六
func getLastSaturday(year int) string {
	// 从该年的12月31日开始
	date := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	// 减去天数直到找到周六（Saturday == 6）
	for date.Weekday() != time.Saturday {
		date = date.AddDate(0, 0, -1)
	}
	return date.Format("2006-01-02")
}
