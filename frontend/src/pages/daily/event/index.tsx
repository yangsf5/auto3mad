import { useState } from 'react';
import { Table, Modal, Button } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryRoutineList } from './service';
import { EditRoutine } from './edit';

export default () => {
  const { data, run } = useRequest(() => {
    return queryRoutineList();
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
      title: 'ID',
      dataIndex: 'id',
    },
    {
      title: '图标',
      dataIndex: 'icon',
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
        title="Edit Routine"
        visible={isEditModalVisible}
        onCancel={onModalCancel}
        footer={null}
        width={1000}
        destroyOnClose={true}
      >
        <EditRoutine></EditRoutine>
      </Modal>
      <div>
        <Table columns={columns} dataSource={data} />
      </div>
    </PageContainer>
  );
};
