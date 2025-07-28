package draw

import (
	"strconv"
	"time"

	gitDraw "github.com/liankui/git-contributions-draw/go/git-contributions-draw"
)

func GetYear(year int) *gitDraw.Year {
	start := getFirstSunday(year)
	end := getLastSaturday(year)
	return &gitDraw.Year{
		Year:  strconv.Itoa(year),
		Total: 100,
		Range: &gitDraw.Range{
			Start: start.Format("2006-01-02"),
			End:   end.Format("2006-01-02"),
		},
	}
}

// 返回某年中的第一个周日
func getFirstSunday(year int) time.Time {
	date := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	// 加天数直到找到周日（Sunday == 0）
	for date.Weekday() != time.Sunday {
		date = date.AddDate(0, 0, 1)
	}
	return date
}

// 返回某年中的最后一个周六
func getLastSaturday(year int) time.Time {
	// 从该年的12月31日开始
	date := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	// 减去天数直到找到周六（Saturday == 6）
	for date.Weekday() != time.Saturday {
		date = date.AddDate(0, 0, -1)
	}
	return date
}
