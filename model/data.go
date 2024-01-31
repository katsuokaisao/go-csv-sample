package model

import (
	"time"
)

type Data struct {
	ID      int
	Name    string
	Age     int
	LoginAt time.Time
}
