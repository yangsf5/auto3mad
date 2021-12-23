import { Table, Button } from 'antd';
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
      key: 'router_pattern',
      render: (text: any) => {
        return <Button href={`${target}${text}`} type='link' target='_blank'>
          {target}{text}
        </Button>;
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
