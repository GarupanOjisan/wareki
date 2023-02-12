// Package wareki 和暦を扱う機能を提供します
package wareki

import "time"

type period struct {
	Gengo string
	Start time.Time
	End   time.Time
}

func (p period) contains(t time.Time) bool {
	if t.Before(p.Start) {
		return false
	}
	if t.After(p.End) {
		return false
	}
	return true
}

func (p period) yearFromStart(t time.Time) int {
	return t.Year() - p.Start.Year() + 1
}

var locationJST = time.FixedZone("Asia/Tokyo", 9*60*60)

var (
	meiji = period{
		Gengo: "明治",
		Start: time.Date(1868, 10, 23, 0, 0, 0, 0, locationJST),
		End:   time.Date(1912, 7, 31, 0, 0, 0, 0, locationJST).Add(-time.Nanosecond),
	}
	taisho = period{
		Gengo: "大正",
		Start: time.Date(1912, 7, 30, 0, 0, 0, 0, locationJST),
		End:   time.Date(1926, 12, 26, 0, 0, 0, 0, locationJST).Add(-time.Nanosecond),
	}
	showa = period{
		Gengo: "昭和",
		Start: time.Date(1926, 12, 25, 0, 0, 0, 0, locationJST),
		End:   time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST).Add(-time.Nanosecond),
	}
	heisei = period{
		Gengo: "平成",
		Start: time.Date(1989, 1, 8, 0, 0, 0, 0, locationJST),
		End:   time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST).Add(-time.Nanosecond),
	}
	reiwa = period{
		Gengo: "令和",
		Start: time.Date(2019, 5, 1, 0, 0, 0, 0, locationJST),
		End:   time.Date(2100, 12, 31, 0, 0, 0, 0, locationJST),
	}
)

// IsMeiji は指定した時刻が明治時代かどうかを返します
func IsMeiji(t time.Time) bool {
	return meiji.contains(t)
}

// IsTaisho は指定した時刻が大正時代かどうかを返します
func IsTaisho(t time.Time) bool {
	return taisho.contains(t)
}

// IsShowa は指定した時刻が昭和時代かどうかを返します
func IsShowa(t time.Time) bool {
	return showa.contains(t)
}

// IsHeisei は指定した時刻が平成時代かどうかを返します
func IsHeisei(t time.Time) bool {
	return heisei.contains(t)
}

// IsReiwa は指定した時刻が令和時代かどうかを返します
func IsReiwa(t time.Time) bool {
	return reiwa.contains(t)
}

type Wareki struct {
	Gengo string
	Year  int
}

// Get は指定した時刻の和暦を返します
// 昭和以前は即日改元なので改元日は2つの要素を返します
func Get(t time.Time) []Wareki {
	if IsReiwa(t) {
		return []Wareki{
			{Gengo: reiwa.Gengo, Year: reiwa.yearFromStart(t)},
		}
	}

	if IsHeisei(t) {
		return []Wareki{
			{Gengo: heisei.Gengo, Year: heisei.yearFromStart(t)},
		}
	}

	// 昭和までは即日改元なので複数の元号が返る
	result := make([]Wareki, 0, 2)
	if IsMeiji(t) {
		result = append(result, Wareki{
			Gengo: meiji.Gengo,
			Year:  meiji.yearFromStart(t),
		})
	}

	if IsTaisho(t) {
		result = append(result, Wareki{
			Gengo: taisho.Gengo,
			Year:  taisho.yearFromStart(t),
		})
	}

	if IsShowa(t) {
		result = append(result, Wareki{
			Gengo: showa.Gengo,
			Year:  showa.yearFromStart(t),
		})
	}

	return result
}
