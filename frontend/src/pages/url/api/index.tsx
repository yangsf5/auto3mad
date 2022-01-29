import { Table } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryApiList } from './service';
import proxy from '../../../../config/proxy';

export default () => {
  const { data } = useRequest(() => {
    return queryApiList();
  });

  const target = proxy.dev['/v2/'].target;

  const columns = [
    {
      title: 'Router Pattern',
      dataIndex: 'router_pattern',
      key: 'router_pattern',
    },
    {
      title: 'API URL',
      dataIndex: 'router_pattern',
      key: 'url',
      render: (text: any) => {
        const apiURL = target + text;
        return <a key={`${apiURL}`} href={`${apiURL}`} target='_blank'>{apiURL}</a>;
      }
    },
    {
      title: 'Controller',
      dataIndex: 'controller',
      key: 'controller',
    },
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
      }}
    >
      <Table columns={columns} dataSource={data} />
    </PageContainer>
  );
};
