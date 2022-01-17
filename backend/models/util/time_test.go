package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPeriodDates(t *testing.T) {
	type input struct {
		startDate, endDate string
	}
	inputs := []input{
		{"2017-04-20", "2017-04-26"},
		{"2017-04-20", "2017-04-15"},
		{"2017-04-20", "2017-04-20"},
		{"2017-04-20", "2017-04-21"},
		{"2017-04-20", "2017-04-19"},
	}
	should := [][]string{
		{"2017-04-20", "2017-04-21", "2017-04-22", "2017-04-23", "2017-04-24", "2017-04-25", "2017-04-26"},
		{"2017-04-20", "2017-04-19", "2017-04-18", "2017-04-17", "2017-04-16", "2017-04-15"},
		{"2017-04-20"},
		{"2017-04-20", "2017-04-21"},
		{"2017-04-20", "2017-04-19"},
	}

	assert := assert.New(t)
	for k, input := range inputs {
		ret := GetPeriodDates(input.startDate, input.endDate)
		assert.Equal(should[k], ret)
	}
}

func TestFixQueryPeriodDates(t *testing.T) {
	type DateSpan struct {
		startDate, endDate string
	}

	inputs := []DateSpan{
		{"2017-04-28", "2017-04-28"},
		{"2017-04-20", "2017-04-28"},
		{"2017-04-28", "2017-04-30"},
		{"", "2017-04-28"},
		{"2017-04-28", ""},
		{"2017-04-20", ""},
	}
	should := []DateSpan{
		{"2017-04-28", "2017-04-28"},
		{"2017-04-20", "2017-04-28"},
		{"2017-04-28", "2017-04-30"},
		{"2017-04-21", "2017-04-28"},
		{"2017-04-28", "2017-04-28"},
		{"2017-04-20", "2017-04-28"},
	}

	assert := assert.New(t)
	for k, input := range inputs {
		fixedStartDate, fixedEndDate := FixQueryPeriodDates(input.startDate, input.endDate)
		assert.Equal(should[k].startDate, fixedStartDate)
		assert.Equal(should[k].endDate, fixedEndDate)
	}
}
