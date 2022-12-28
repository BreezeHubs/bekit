package sqlx

import "fmt"

var (
	DateFormatYM          = "%Y%m"
	DateFormatYMSpace     = "%Y %m"
	DateFormatYMLine      = "%Y-%m"
	DateFormatYMUnderLine = "%Y_%m"
	DateFormatYMZH        = "%Y年%m月"

	DateFormatYMD          = "%Y%m%d"
	DateFormatYMDSpace     = "%Y %m %d"
	DateFormatYMDLine      = "%Y-%m-%d"
	DateFormatYMDUnderLine = "%Y_%m_%d"
	DateFormatYMDZH        = "%Y年%m月%d日"
)

func DateFormat(field string, format string) string {
	// format %Y-%m-%d %H:%i:%s => DATE_FORMAT(`created_at`, "%Y-%m-%d %H:%i:%s")
	return fmt.Sprintf("DATE_FORMAT(`%s`, \"%s\")", field, format)
}
