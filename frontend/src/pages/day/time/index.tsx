import { useState } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { Table, Card } from 'antd';
import { queryTimeList } from './service';
import { TimeInfo } from './data';
import { Timestamp } from './timestamp';


export default () => {
  const [timeList, setTimeList] = useState<TimeInfo[]>();

  const columns = [
    {
      title: '地区',
      dataIndex: 'area',
    },
    {
      title: '时区',
      dataIndex: 'timezone',
    },
    {
      title: '当地时间',
      dataIndex: 'time',
    },
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
        ],
      }}
    >
      <Card>
        <Timestamp />
      </Card>

      <br />

      <Card>
        <ProForm<{ timestamp: number }>
          layout={'inline'}
          onFinish={async (values) => {
            const { data } = await queryTimeList(values.timestamp);
            setTimeList(data);
          }}
        >
          <ProFormText name='timestamp' label='Timestamp' />
        </ProForm>

        <br />

        <Table columns={columns} dataSource={timeList} />
      </Card>
    </PageContainer >
  );
};
