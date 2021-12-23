import { Table } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryApiList } from './service';

export default () => {
  const { data } = useRequest(() => {
    return queryApiList();
  });

  const columns = [
    {
      title: 'Router Pattern',
      dataIndex: 'router_pattern',
      key: 'router_pattern',
    },
    {
      title: 'Controller',
      dataIndex: 'controller',
      key: 'controller',
    },
  ];

  const list = data || [];
  console.log(data);

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
      }}
    >
      <div>
        <Table columns={columns} dataSource={list} />
      </div>
    </PageContainer>
  );
};
