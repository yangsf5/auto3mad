import { useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Button, DatePicker, Table, Avatar } from 'antd';
import moment from 'moment';
import { useRequest } from 'umi';
import { queryStat } from './service';

const { RangePicker } = DatePicker;
const monthFormat = 'YYYY-MM';

export default () => {
  const [queryDate, setQueryDate] = useState([moment('2021-02-25', 'YYYY-MM-DD'), moment()]);

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
      title: '',
      dataIndex: '',
    },
    {
      title: '',
      dataIndex: '',
    },
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <RangePicker picker='month' format='YYYY-MM' defaultValue={queryDate} onChange={onChangeDate} />,
        ],
      }}
    >
      <Table columns={columns} dataSource={data} pagination={false} />
    </PageContainer>
  );
};
