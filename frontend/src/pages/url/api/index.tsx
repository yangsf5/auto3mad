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
        return <a href={`${apiURL}`} target='_blank'>{apiURL}</a>;
      }
    },
    {
      title: 'Controller',
      dataIndex: 'controller',
      key: 'controller',
    },
  ];

  const list = data || [];

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
