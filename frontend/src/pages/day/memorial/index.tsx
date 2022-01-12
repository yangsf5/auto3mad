import { Table } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryMemorialList } from './service';

export default () => {
  const { data } = useRequest(() => {
    return queryMemorialList();
  });

  const columns = [
    {
      title: '纪念',
      dataIndex: 'desc',
    },
    {
      title: '日期',
      dataIndex: 'date',
    },
    {
      title: '已经过去',
      dataIndex: 'passed',
    },
    {
      title: '下个纪念还剩',
      dataIndex: 'next_left',
    },
    {
      title: '下个纪念日',
      dataIndex: 'next_date',
    },
    {
      title: '周期',
      dataIndex: 'remind_type_desc',
    },
    {
      title: '已过周期数',
      dataIndex: 'cycle_count',
    }
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
      }}
    >
      <div>
        <Table columns={columns} dataSource={data} />
      </div>
    </PageContainer>
  );
};
