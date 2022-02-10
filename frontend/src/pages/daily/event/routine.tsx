import { Table, Avatar, Progress } from 'antd';
import { useRequest } from 'umi';
import { queryRoutineList } from './service';

const RoutineTable = (props: any) => {
  const { date, refresh } = props;

  const { data } = useRequest(() => queryRoutineList(date), { refreshDeps: [date, refresh] });

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
      <Table rowKey='id' columns={columns} dataSource={data} pagination={false} size='small' />
    </div>
  );
};

export {
  RoutineTable,
};
