import { useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { DatePicker, Table } from 'antd';
import { Column } from '@ant-design/charts';
import moment from 'moment';
import { useRequest } from 'umi';
import type { StatInfo } from './data';
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

  const columns = [
    {
      title: 'routine',
      dataIndex: 'routine',
    },
    {
      title: 'month',
      dataIndex: 'month',
    },
    {
      title: 'spend',
      dataIndex: 'spend',
    },
  ];

  const chartConfig = {
    data: data || [],
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
    }
  };

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
      <Table<StatInfo>
        rowKey={(r: StatInfo) => r.routine + r.month}
        columns={columns}
        dataSource={data}
        size='small'
      />
    </PageContainer >
  );
};
