import { Table, Avatar, Progress, Select } from 'antd';
import { RoutineInfo } from './data';

const { Option } = Select;

const RoutineTable = (props: { dataSource: RoutineInfo[] | undefined }) => {
  const { dataSource } = props;

  const columns = [
    {
      title: '日拱项',
      dataIndex: 'short_name',
      render: (text: any, record: any) => <div><Avatar size={16} src={record.icon} /> {text}</div>,
    },
    {
      title: '日拱范围',
      dataIndex: 'event_scope',
    },
    {
      title: '当前默认事件',
      dataIndex: 'event_default',
    },
    {
      title: '日投 / 目标 M',
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
      title: '累投 / 目标 H',
      dataIndex: 'total_spend',
      render: (_: any, record: any) => <>{record.total_spend} / {record.total_will_spend}</>,
    },
    {
      title: '累投进度',
      dataIndex: 'total_spend',
      key: 'total_progress',
      width: 100,
      render: (_: any, record: any) => <><Progress type='line' percent={Math.floor(record.total_spend / record.total_will_spend * 100)} /></>,
    },
    {
      title: '产出阶段',
      dataIndex: 'start_date',
      render: (_: any, record: any) => <>{record.start_date} ~ {record.end_date}</>,
    },
    {
      title: '产出 / 目标',
      dataIndex: 'object',
      render: (_: any, record: any) => <>{record.progress} / {record.object}</>,
    },
    {
      title: '产出进度',
      dataIndex: 'progress',
      width: 100,
      render: (_: any, record: any) => <><Progress type='line' percent={Math.floor(record.progress / record.object * 100)} /></>,
    },
  ];

  return (
    <Table rowKey='id' columns={columns} dataSource={dataSource} pagination={false} size='small' />
  );
};

const RoutineSelect = (props: { dataSource: RoutineInfo[] | undefined, onChange: any }) => {
  const { dataSource, onChange } = props;

  const options = dataSource?.map(val => (
    <Option key={val.id} value={val.id}>
      <Avatar size={16} src={val.icon}></Avatar> {val.short_name}
    </Option>
  ));

  return (
    <>
      <Select
        style={{ width: 140 }}
        placeholder='开始一项日拱'
        onChange={onChange}
      >
        {options}
      </Select>
    </>
  );
};

export {
  RoutineTable,
  RoutineSelect,
};
