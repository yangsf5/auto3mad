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
      title: '日预算 M',
      dataIndex: 'will_spend',
    },
    {
      title: '今日已投 M',
      dataIndex: 'today_spend',
      render: (text: any, record: any) => <div><Progress type='line' percent={Math.floor(record.today_spend / record.will_spend * 100)} /></div>,
    },
    {
      title: '累投 H',
      dataIndex: 'total_spend',
    },
  ];

  return (
    <div>
      <Table rowKey='id' columns={columns} dataSource={dataSource} pagination={false} size='small' />
    </div>
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
    <Select style={{ width: 140 }} onChange={onChange}>
      {options}
    </Select>
  );
};

export {
  RoutineTable,
  RoutineSelect,
};
