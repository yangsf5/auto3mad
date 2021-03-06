import { useState } from 'react';
import { Table, Avatar, Progress, Space, Select, Radio, Tooltip } from 'antd';
import { InfoCircleOutlined } from '@ant-design/icons';
import { RoutineInfo } from './data';


const { Option } = Select;

const RoutineTable = (props: { dataSource: RoutineInfo[] | undefined }) => {
  const { dataSource } = props;

  const totalSpend = (record: any) => {
    return record.total_spend + record.history_spend;
  };

  const progress = (record: any, current: number, total: number, steps: number, hidden: boolean) => {
    if (hidden && record.history_spend != 0) {
      return '--';
    }

    return <>
      <Progress type='line' percent={Math.floor(current / total * 100)} steps={steps} />
    </>;
  };

  const progressDigit = (record: any, current: number, total: number, hidden: boolean) => {
    if (hidden && record.history_spend != 0) {
      return '--';
    }

    return current + '/' + total;
  };

  const objectProgress = (record: any) => {
    if (record.object == 0) {
      return '--';
    }

    return <>{record.progress}/{record.object} {record.object_unit}</>;
  }

  const columns = [
    {
      title: '日拱家族',
      dataIndex: 'short_name',
      render: (text: any, record: any) => <>
        <Avatar size={16} src={record.icon} />
        <span></span> {text} <span></span>
        <Tooltip title={'日拱范围 - ' + record.event_scope}>
          <InfoCircleOutlined />
        </Tooltip>
      </>,
    },
    {
      title: '开始日期',
      dataIndex: 'start_date',
    },
    {
      title: '日投/预算 M',
      dataIndex: 'today_spend',
      render: (_: any, record: any) => progressDigit(record, record.today_spend, record.will_spend, false),
    },
    {
      title: '日进度',
      dataIndex: 'today_spend',
      key: 'today_progress',
      width: 100,
      render: (_: any, record: any) => progress(record, record.today_spend, record.will_spend, 0, true),
    },
    {
      title: '周投/应投 M',
      dataIndex: 'week_spend',
      render: (_: any, record: any) => progressDigit(record, record.week_spend, record.week_will_spend, true),
    },
    {
      title: '周进度',
      dataIndex: 'week_spend',
      key: 'week_progress',
      width: 100,
      render: (_: any, record: any) => progress(record, record.week_spend, record.week_will_spend, 5, true),
    },
    {
      title: '月投/应投 M',
      dataIndex: 'month_spend',
      render: (_: any, record: any) => progressDigit(record, record.month_spend, record.month_will_spend, true),
    },
    {
      title: '月进度',
      dataIndex: 'month_spend',
      key: 'month_progress',
      width: 100,
      render: (_: any, record: any) => progress(record, record.month_spend, record.month_will_spend, 4, true),
    },
    {
      title: '累投/应投 H',
      dataIndex: 'total_spend',
      render: (_: any, record: any) => progressDigit(record, totalSpend(record), record.total_will_spend, false),
    },
    {
      title: '至今健康度',
      dataIndex: 'total_spend',
      key: 'total_progress',
      width: 100,
      render: (_: any, record: any) => progress(record, totalSpend(record), record.total_will_spend, 0, false),
    },
    {
      title: '产出/目标',
      dataIndex: 'object',
      render: (_: any, record: any) => objectProgress(record),
    },
  ];

  return (
    <Table rowKey='id' columns={columns} dataSource={dataSource} pagination={false} size='small' bordered />
  );
};

const RoutineSelect = (props: { dataSource: RoutineInfo[] | undefined, onSelect: any }) => {
  const { dataSource, onSelect } = props;

  const [createType, setCreateType] = useState(0);

  const selOptions = dataSource?.map(val => (
    <Option key={val.id} value={val.id} disabled={val.history_spend != 0}>
      <Avatar size={16} src={val.icon}></Avatar> {val.short_name}
    </Option>
  ));

  return (
    <>
      <Space>
        <Radio.Group
          onChange={(e) => setCreateType(e.target.value)}
          value={createType}
        >
          <Radio.Button value={0}>新开</Radio.Button>
          <Radio.Button value={1}>补录</Radio.Button>
        </Radio.Group>
        <Select
          style={{ width: 140 }}
          placeholder='开始一项日拱'
          onSelect={(routineID) => {
            if (routineID != undefined) {
              onSelect(createType, routineID);
            }
          }}
          allowClear={true}
        >
          {selOptions}
        </Select>
      </Space>
    </>
  );
};

export {
  RoutineTable,
  RoutineSelect,
};
