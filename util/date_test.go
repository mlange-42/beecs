package util_test

import (
	"testing"
	"time"

	"github.com/mlange-42/beecs/util"
	"github.com/stretchr/testify/assert"
)

func TestTickToDate(t *testing.T) {
	assert.Equal(t, date(0, time.January, 1), util.TickToDate(0))

	assert.Equal(t, date(0, time.February, 28), util.TickToDate(58))
	assert.Equal(t, date(0, time.March, 1), util.TickToDate(59))

	assert.Equal(t, date(1, time.January, 1), util.TickToDate(365))
	assert.Equal(t, date(2, time.January, 1), util.TickToDate(730))
}

func date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
