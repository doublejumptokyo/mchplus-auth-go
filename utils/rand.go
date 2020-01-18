package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

func RandNumberString(l int) string {
	r := rand.New(randSrc)
	ret := r.Intn(int(math.Pow10(l)))
	format := "%0" + strconv.Itoa(l) + "d"
	return fmt.Sprintf(format, ret)
}
