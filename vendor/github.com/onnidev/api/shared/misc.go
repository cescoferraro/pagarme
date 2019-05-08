package shared

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pressly/chi/render"
	"gopkg.in/mgo.v2/bson"
)

// Month TODO: NEEDS COMMENT INFO
func Month(mes int) string {
	switch mes {
	case 1:
		return "Janeiro"
	case 2:
		return "Fevereiro"
	case 3:
		return "Março"
	case 4:
		return "Abril"
	case 5:
		return "Maio"
	case 6:
		return "Junho"
	case 7:
		return "Julho"
	case 8:
		return "Agosto"
	case 9:
		return "Setembro"
	case 10:
		return "Outubro"
	case 11:
		return "Novembro"
	case 12:
		return "Dezembro"
	}
	return ""

}

// InTimeSpan TODO: NEEDS COMMENT INFO
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// Contains kjsdnf
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsTime  kjsdnf
func ContainsTime(s []time.Time, e time.Time) bool {
	for _, a := range s {
		if DateEqual(a, e) {
			return true
		}
	}
	return false
}

// DateEqual TODO: NEEDS COMMENT INFO
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	h1, z1, k1 := date1.Clock()
	y2, m2, d2 := date2.Date()
	h2, z2, k2 := date2.Clock()
	return y1 == y2 && m1 == m2 && d1 == d2 && h1 == h2 && z1 == z2 && k1 == k2
}

// ContainsObjectIDIndex TODO: NEEDS COMMENT INFO
func ContainsObjectIDIndex(s []bson.ObjectId, e bson.ObjectId) (bool, int) {
	for index, a := range s {
		if a.Hex() == e.Hex() {
			return true, index
		}
	}
	return false, 0
}

// ContainsObjectID TODO: NEEDS COMMENT INFO
func ContainsObjectID(s []bson.ObjectId, e bson.ObjectId) bool {
	for _, a := range s {
		if a.Hex() == e.Hex() {
			return true
		}
	}
	return false
}

// StandardONNiError TODO: NEEDS COMMENT INFO
func StandardONNiError(code int, msg string) struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
} {
	return struct {
		Code   int    `json:"code"`
		Reason string `json:"reason"`
	}{
		Code:   code,
		Reason: msg,
	}
}

// MakeONNiError TODO: NEEDS COMMENT INFO
func MakeONNiError(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Println("ERROR: ", err.Error())
	render.Status(r, code)
	result := StandardONNiError(code, err.Error())
	log.Println("ERROR JSON: ", result)
	render.JSON(w, r, result)
}

// RangeIn skdjfnsdkf
func RangeIn(low, hi int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa(low + rand.Intn(hi-low))
}

// EncryptPassword2 kdjfsdf
func EncryptPassword2(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}

// LatLongDistance sdkjfn
func LatLongDistance(lat1, lon1, lat2, lon2 float64) float64 {
	hsin := func(theta float64) float64 { return math.Pow(math.Sin(theta/2), 2) }
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180
	r = 6378100
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)
	return 2 * r * math.Asin(math.Sqrt(h))
}

// OrBlank TODO: NEEDS COMMENT INFO
func OrBlank(og, sent string) string {
	if sent == "" {
		return og
	}
	return sent
}

// OrBlanktime TODO: NEEDS COMMENT INFO
func OrBlanktime(og, sent time.Time) time.Time {
	if sent.IsZero() {
		return og
	}
	return sent
}

// AddCents TODO: NEEDS COMMENT INFO
func AddCents(text string) string {
	result := text
	if !strings.Contains(text, ".") {
		return strings.Replace(text+",00", ".", ",", -1)
	}
	log.Println("adding cents to", text)
	splt := strings.Split(text, ".")
	log.Println("adding cents to", splt)
	log.Println(len(splt[1]))
	if len(splt[1]) == 1 {
		result = text + "0"
	}
	return strings.Replace(result, ".", ",", -1)
}
