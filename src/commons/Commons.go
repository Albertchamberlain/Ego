package commons

import (
	"math/rand"
	"strconv"
	"time"
)

func GenId() int {
	rand.Seed(time.Now().UnixNano())
	id, _ := strconv.Atoi(strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(int(time.Now().Unix())))
	return id
}
