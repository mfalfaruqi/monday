package monday

import "strings"

// ============================================================
// Format rules for "ar_SA" locale: Arabic (Saudi Arabia)
// ============================================================

var longDayNamesArSA = map[string]string{
	"Sunday":    "الأحد",
	"Monday":    "الأثنين",
	"Tuesday":   "الثلاثاء",
	"Wednesday": "الأربعاء",
	"Thursday":  "الخميس",
	"Friday":    "الجمعه",
	"Saturday":  "السبت",
}

var shortDayNamesArSA = map[string]string{
	"Sun": "الأحد",
	"Mon": "الأثنين",
	"Tue": "الثلاثاء",
	"Wed": "الأربعاء",
	"Thu": "الخميس",
	"Fri": "الجمعه",
	"Sat": "السبت",
}

var longMonthNamesArSA = map[string]string{
	"January":   "يناير",
	"February":  "فبراير",
	"March":     "مارس",
	"April":     "أبريل",
	"May":       "مايو",
	"June":      "يونيو",
	"July":      "يوليو",
	"August":    "أغسطس",
	"September": "سبتمبر",
	"October":   "أكتوبر",
	"November":  "نوفمبر",
	"December":  "ديسمبر",
}

var shortMonthNamesArSA = map[string]string{
	"Jan": "يناير",
	"Feb": "فبراير",
	"Mar": "مارس",
	"Apr": "أبريل",
	"May": "مايو",
	"Jun": "يونيو",
	"Jul": "يوليو",
	"Aug": "أغسطس",
	"Sep": "سبتمبر",
	"Oct": "أكتوبر",
	"Nov": "نوفمبر",
	"Dec": "ديسمبر",
}

var periodsArSA = map[string]string{
	"am": "ص",
	"pm": "م",
	"AM": "ص",
	"PM": "م",
}

var numbersArSa = map[rune]rune{
	'0': '٠',
	'1': '١',
	'2': '٢',
	'3': '٣',
	'4': '٤',
	'5': '٥',
	'6': '٦',
	'7': '٧',
	'8': '٨',
	'9': '٩',
}

func formatFuncArCommon(locale Locale) internalFormatFunc {
	return func(value, layout string) (res string) {
		res = commonFormatFunc(value, layout,
			knownDaysShort[locale], knownDaysLong[locale], knownMonthsShort[locale], knownMonthsLong[locale], knownPeriods[locale])

		// replace number with arabic number
		newRes := &strings.Builder{}
		for _, r := range res {
			if v, ok := numbersArSa[r]; ok {
				newRes.WriteRune(v)
				continue
			}

			newRes.WriteRune(r)
		}
		return newRes.String()
	}
}

func parseFuncArCommon(locale Locale) internalParseFunc {
	return func(layout, value string) string {
		res := formatFuncArCommon(locale)
		return res(value, layout)
	}
}
