package daily

import (
	"strconv"
	"time"

	"github.com/yangsf5/auto3mad/backend/controllers/base"
	"github.com/yangsf5/auto3mad/backend/models/db/daily"
)

const format = "2006-01"

type StatController struct {
	base.Controller
	mr *daily.RoutineModel
	me *daily.EventModel
}

func (c *StatController) Prepare() {
	c.Controller.Prepare()

	c.mr = daily.NewRoutineModel()
	c.me = daily.NewEventModel(c.GetMyUserID())
}

type statChartItem struct {
	Routine string `json:"routine"`
	Month   string `json:"month"`
	Spend   int    `json:"spend"`
}

type statTableItem struct {
	Routine string `json:"routine"`
	Spends  []int  `json:"spends"`
}

func (c *StatController) Get() {
	firstMonth := c.GetString("first_month")
	lastMonth := c.GetString("last_month")

	if firstMonth == "" || lastMonth == "" {
		firstMonth = time.Now().Format(format)
		lastMonth = firstMonth
	}

	months, err := getPeriodMonths(firstMonth, lastMonth)
	c.JSONErrorAbort(err)

	var rrs []daily.Routine
	err = c.mr.GetAllOrderBy(&rrs, "sort")
	c.JSONErrorAbort(err)

	chartRets := []statChartItem{}
	tableRets := []statTableItem{}

	for _, r := range rrs {
		spends, err := c.me.GetRoutineSpendGroupByMonth(r.ID, firstMonth, lastMonth)
		c.JSONErrorAbort(err)

		tableItem := statTableItem{
			Routine: r.ShortName,
			Spends:  []int{},
		}

		for _, month := range months {
			chartRet := statChartItem{
				Routine: r.ShortName,
				Month:   month,
			}

			if spend, ok := spends[month]; ok {
				chartRet.Month = month

				spendStr := spend.(string) // nolint
				chartRet.Spend, _ = strconv.Atoi(spendStr)
			}

			chartRets = append(chartRets, chartRet)
			tableItem.Spends = append(tableItem.Spends, chartRet.Spend)
		}

		tableRets = append(tableRets, tableItem)
	}

	c.JSONOK(map[string]interface{}{
		"chart":  chartRets,
		"table":  tableRets,
		"months": months,
	})
}

func getPeriodMonths(firstMonth, lastMonth string) ([]string, error) {
	first, err := time.Parse(format, firstMonth)
	if err != nil {
		return nil, err
	}

	last, err := time.Parse(format, lastMonth)
	if err != nil {
		return nil, err
	}

	if first.Unix() > last.Unix() {
		return nil, nil
	}

	months := []string{}
	for t := first; t.Unix() <= last.Unix(); t = t.AddDate(0, 1, 0) {
		months = append(months, t.Format(format))
	}

	return months, nil
}
