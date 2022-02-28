import { useState } from 'react';
import { Table, Avatar, Progress, Space, Select, Radio } from 'antd';
import { RoutineInfo } from './data';


const { Option } = Select;

const RoutineTable = (props: { dataSource: RoutineInfo[] | undefined }) => {
  const { dataSource } = props;

  const totalSpend = (record: any) => {
    return record.total_spend + record.history_spend;
  };

  const columns = [
    {
      title: '日拱家族',
      dataIndex: 'short_name',
      render: (text: any, record: any) => <div><Avatar size={16} src={record.icon} /> {text}</div>,
    },
    {
      title: '开始日期',
      dataIndex: 'start_date',
    },
    {
      title: '累投 / 预算 H',
      dataIndex: 'total_spend',
      render: (_: any, record: any) => (
        <>{totalSpend(record)} / {record.total_will_spend}</>
      ),
    },
    {
      title: '累投健康度',
      dataIndex: 'total_spend',
      key: 'total_progress',
      width: 100,
      render: (_: any, record: any) => <><Progress type='line' percent={Math.floor(totalSpend(record) / record.total_will_spend * 100)} /></>,
    },
    {
      title: '日拱周数',
      dataIndex: 'week_passed',
    },
    {
      title: '日投 / 预算 M',
      dataIndex: 'today_spend',
      render: (_: any, record: any) => <>{record.today_spend} / {record.will_spend}</>,
    },
    {
      title: '日投进度',
      dataIndex: 'today_spend',
      key: 'today_progress',
      width: 100,
      render: (_: any, record: any) => <><Progress type='line' percent={Math.floor(record.today_spend / record.will_spend * 100)} /></>,
    },
    {
      title: '产出 / 目标',
      dataIndex: 'object',
      render: (_: any, record: any) => <>{record.progress} / {record.object} {record.object_unit}</>,
    },
    {
      title: '产出进度',
      dataIndex: 'progress',
      width: 100,
      render: (_: any, record: any) => <><Progress type='line' percent={Math.floor(record.progress / record.object * 100)} /></>,
    },
    {
      title: '日拱范围',
      dataIndex: 'event_scope',
    },
    {
      title: '当前默认事件',
      dataIndex: 'event_default',
    },
  ];

  return (
    <Table rowKey='id' columns={columns} dataSource={dataSource} pagination={false} size='small' bordered />
  );
};

const RoutineSelect = (props: { dataSource: RoutineInfo[] | undefined, onChange: any }) => {
  const { dataSource, onChange } = props;

  const [createType, setCreateType] = useState(0);

  const selOptions = dataSource?.map(val => (
    <Option key={val.id} value={val.id}>
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
          onChange={(routineID) => {
            console.log(routineID);
            if (routineID != undefined) {
              onChange(createType, routineID);
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
