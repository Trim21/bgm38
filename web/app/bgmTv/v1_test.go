package bgmTv

import (
	"github.com/magiconair/properties/assert"
	//"gotest.tools/assert"

	"testing"
	"time"
)

func TestGetNextAirDate(t *testing.T) {
	assert.Equal(t, getAirDayOffset(time.Sunday, 7),
		0, "Sunday and Sunday")

	assert.Equal(t, getAirDayOffset(time.Monday, 1),
		0, "Monday and Monday")

	assert.Equal(t, getAirDayOffset(time.Monday, 2),
		1, "Monday and Tuesday")

	assert.Equal(t, getAirDayOffset(time.Tuesday, 3),
		1, "Tuesday and Wednesday")

	assert.Equal(t, getAirDayOffset(time.Tuesday, 4),
		2, "Tuesday and Thursday")

	assert.Equal(t, getAirDayOffset(time.Thursday, 7),
		3, "Thursday and Sunday")

	assert.Equal(t, getAirDayOffset(time.Tuesday, 4),
		2, "Tuesday and Thursday")

	assert.Equal(t, getAirDayOffset(time.Sunday, 1),
		1, "Sunday and Monday")

}
