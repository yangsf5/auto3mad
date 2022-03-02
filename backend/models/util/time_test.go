package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDateTimestamp(t *testing.T) {
	type shouldT struct {
		first, last int64
	}

	inputs := []string{
		"2021-01-01",
		"2022-02-27",
		"2022-03-01",
	}

	shoulds := []shouldT{
		{1609430400, 1609516799},
		{1645891200, 1645977599},
		{1646064000, 1646150399},
	}

	assert := assert.New(t)
	for k, input := range inputs {
		retFirst, retLast := GetDateTimestamp(input)
		assert.Equal(shoulds[k].first, retFirst)
		assert.Equal(shoulds[k].last, retLast)
	}
}

func TestGetPeriodDates(t *testing.T) {
	type inputT struct {
		startDate, endDate string
	}
	inputs := []inputT{
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

func TestGetWeekTimestamp(t *testing.T) {
	type shouldT struct {
		first, last int64
	}

	inputs := []string{
		"2022-02-27",
		"2022-02-28",
		"2022-03-01",
		"2022-03-04",
		"2022-03-06",
		"2022-03-07",
	}

	shoulds := []shouldT{
		{1645372800, 1645977599},
		{1645977600, 1646582399},
		{1645977600, 1646582399},
		{1645977600, 1646582399},
		{1645977600, 1646582399},
		{1646582400, 1647187199},
	}

	assert := assert.New(t)
	for k, input := range inputs {
		retFirst, retLast := GetWeekTimestamp(input)
		assert.Equal(shoulds[k].first, retFirst)
		assert.Equal(shoulds[k].last, retLast)
	}
}

func TestGetMonthTimestamp(t *testing.T) {
	type shouldT struct {
		first, last int64
	}

	inputs := []string{
		"2022-02-15",
		"2022-02-28",
		"2022-03-01",
		"2022-03-15",
		"2022-03-31",
	}

	shoulds := []shouldT{
		{1643644800, 1646063999},
		{1643644800, 1646063999},
		{1646064000, 1648742399},
		{1646064000, 1648742399},
		{1646064000, 1648742399},
	}

	assert := assert.New(t)
	for k, input := range inputs {
		retFirst, retLast := GetMonthTimestamp(input)
		assert.Equal(shoulds[k].first, retFirst)
		assert.Equal(shoulds[k].last, retLast)
	}
}
