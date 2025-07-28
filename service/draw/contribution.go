package draw

import (
	"errors"

	gitDraw "github.com/liankui/git-contributions-draw/go/git-contributions-draw"
)

func GetContributions(year int, pattern []string) ([]*gitDraw.Contribution, error) {
	if len(pattern[0]) > 53 {
		return nil, errors.New("pattern too long")
	}

	var contributions []*gitDraw.Contribution
	startDate := getFirstSunday(year)
	for x := 0; x < len(pattern[0]); x++ { // 每列一周
		for y := 0; y < len(pattern); y++ { // 每行一周中的某天
			char := pattern[y][x]
			date := startDate.AddDate(0, 0, x*7+y)
			count := 0
			color := "#ebedf0"
			intensity := "0"

			if char == '*' {
				count = 5
				color = "#196127"
				intensity = "4"
			}

			contributions = append(contributions, &gitDraw.Contribution{
				Date:      date.Format("2006-01-02"),
				Count:     int32(count),
				Color:     color,
				Intensity: intensity,
			})
		}
	}

	endDate := getFirstSunday(year)
	for x := len(pattern[0]); x <= 53; x++ {
		for y := 0; y < len(pattern); y++ {
			date := startDate.AddDate(0, 0, x*7+y)
			if date.After(endDate) {
				break
			}
			contributions = append(contributions, &gitDraw.Contribution{
				Date:      date.Format("2006-01-02"),
				Count:     0,
				Color:     "#ebedf0",
				Intensity: "0",
			})
		}
	}

	return contributions, nil

	//jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	//fmt.Println(string(jsonBytes))

	//builder := strings.Builder{}
	//builder.WriteString("DATES=(")
	//for _, cont := range contributions {
	//	builder.WriteString("\"")
	//	builder.WriteString(cont.Date)
	//	builder.WriteString("\"")
	//	builder.WriteString(" ")
	//}
	//builder.WriteString(")")
	//fmt.Println(builder.String())
}
