import { useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { DatePicker, Table } from 'antd';
import { Column } from '@ant-design/charts';
import moment from 'moment';
import { useRequest } from 'umi';
import { each, groupBy } from '@antv/util';

import { queryStat } from './service';

const { RangePicker } = DatePicker;
const monthFormat = 'YYYY-MM';

export default () => {
  const [queryDate, setQueryDate] = useState([moment('2022-01-01', 'YYYY-MM-DD'), moment()]);

  const { data, run } = useRequest(() => {
    var firstMonth: string = queryDate[0].format(monthFormat);
    var lastMonth: string = queryDate[1].format(monthFormat);
    return queryStat(firstMonth, lastMonth);
  });

  const onChangeDate = (dates: [moment.Moment, moment.Moment]) => {
    setQueryDate(dates);
    run();
  };

  const annotations: any = [];
  each(groupBy(data?.chart, 'month'), (values, k) => {
    const value = values.reduce((a: any, b: any) => a + b.spend, 0);
    annotations.push({
      type: 'text',
      position: [k, value],
      content: `${value}`,
      style: {
        textAlign: 'center',
        fontSize: 14,
        fill: 'rgba(0,0,0,0.85)',
      },
      offsetY: -20,
    });
  });

  const chartConfig = {
    data: data?.chart || [],
    isStack: true,
    xField: 'month',
    yField: 'spend',
    seriesField: 'routine',
    label: {
      position: 'middle',
      layout: [
        {
          type: 'interval-adjust-position',
        },
        {
          type: 'interval-hide-overlap',
        },
        {
          type: 'adjust-color'
        },
      ],
    },
    interactions: [
      {
        type: 'active-region',
        enable: false,
      },
    ],
    connectedArea: {
      style: (oldStyle: any) => {
        return {
          fill: 'rgba(0,0,0,0.25)',
          stroke: oldStyle.fill,
          lineWidth: 0.5,
        };
      },
    },
    color: ['#5B8FF9', '#F4664A', '#5D7092', '#FF9845', '#6DC8EC', '#F6BD16', '#FF99C3', '#30BF78'],
    annotations,
  };

  const tableData: any = [];

  const columns = [
    {
      title: 'routine',
      dataIndex: 'routine',
    },
  ];

  data?.months.forEach(v => {
    columns.push({ title: v, dataIndex: v });
  })

  data?.table.forEach(v => {
    var dataItem: any = { 'routine': v.routine };
    data?.months.forEach((month, i) => {
      dataItem[month] = v.spends[i];
    })

    tableData.push(dataItem);
  });

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <RangePicker key='month' picker='month' format='YYYY-MM' defaultValue={queryDate} onChange={onChangeDate} />,
        ],
      }}
    >
      <Column {...chartConfig} />
      <br />
      <Table
        rowKey='routine'
        columns={columns}
        dataSource={tableData}
        size='small'
      />
    </PageContainer >
  );
};
