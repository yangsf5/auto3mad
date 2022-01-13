import { useState } from 'react';
import { Table, Modal, Button } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryMemorialList } from './service';
import { EditMerial } from './edit';

export default () => {
  const { data, run } = useRequest(() => {
    return queryMemorialList();
  });

  const [isEditModalVisible, setEditModalVisible] = useState(false);
  const showEditModal = () => {
    setEditModalVisible(true);
  };
  const onModalCancel = () => {
    setEditModalVisible(false);
    run();
  }

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
      title: '已过周期数',
      dataIndex: 'cycle_count',
    }
  ];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <Button onClick={showEditModal}>Edit</Button>
        ],
      }}
    >
      <Modal
        title="Edit Memorial"
        visible={isEditModalVisible}
        onCancel={onModalCancel}
        footer={null}
        width={1000}
        destroyOnClose={true}
      >
        <EditMerial></EditMerial>
      </Modal>
      <div>
        <Table columns={columns} dataSource={data} />
      </div>
    </PageContainer>
  );
};
