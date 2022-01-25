import { PageContainer } from '@ant-design/pro-layout';
import { Table } from 'antd';
import { useRequest } from 'umi';
import { queryTimeList } from './service';


export default () => {
  const { data } = useRequest(() => {
    return queryTimeList(1451404800);
  });

  const columns = [
    {
      title: '地区',
      dataIndex: 'area',
    },
    {
      title: '时区',
      dataIndex: 'timezone',
    },
    {
      title: '当地时间',
      dataIndex: 'time',
    },
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
        ],
      }}
    >
      <Table columns={columns} dataSource={data} />
    </PageContainer>
  );
};
