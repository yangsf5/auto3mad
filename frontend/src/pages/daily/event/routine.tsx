import { Table, Avatar } from 'antd';
import { useRequest } from 'umi';
import { queryRoutineList } from './service';

var RoutineRun: Function;

const RoutineTable = () => {
  const { data, run } = useRequest(() => {
    return queryRoutineList();
  });

  RoutineRun = run;

  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: '图标',
      dataIndex: 'icon',
      render: (text: any) => <Avatar size={16} src={text} />
    },
    {
      title: '简称',
      dataIndex: 'short_name',
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
      dataIndex: 'spend',
    },
    {
      title: '累积投入 H',
      dataIndex: 'total_spend',
    },
  ];

  return (
    <div>
      <Table columns={columns} dataSource={data} size='small' pagination={false} />
    </div>
  );
};

export {
  RoutineTable,
  RoutineRun,
};
