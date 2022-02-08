import { Table, Avatar } from 'antd';
import { useRequest } from 'umi';
import { queryRoutineList } from './service';

var RoutineRun: Function;

const RoutineTable = (props: any) => {
  const { date } = props;

  const { data, run } = useRequest(() => {
    return queryRoutineList(date);
  });

  RoutineRun = run;

  const columns = [
    {
      title: '简称',
      dataIndex: 'short_name',
      render: (text: any, record: any) => <div><Avatar size={16} src={record.icon} /> {text}</div>,
    },
    {
      title: '例行事件内容',
      dataIndex: 'event',
    },
    {
      title: '预算 M',
      dataIndex: 'will_spend',
    },
    {
      title: '今日已投入 M',
      dataIndex: 'today_spend',
    },
    {
      title: '累积投入 H',
      dataIndex: 'total_spend',
    },
  ];

  return (
    <div>
      <Table columns={columns} dataSource={data} pagination={false} size='small' />
    </div>
  );
};

export {
  RoutineTable,
  RoutineRun,
};
