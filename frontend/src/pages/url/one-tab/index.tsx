import { useState } from 'react';
import { Card, List, Image, Button, Modal } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryFakeList } from './service';
import { EditGroup } from './EditGroup';

const CardList = () => {
  const [isEditModalVisible, setEditModalVisible] = useState(false);
  const showEditModal = () => {
    setEditModalVisible(true);
  };
  const handleOk = () => {
    setEditModalVisible(false);
  }
  const handleCancel = () => {
    setEditModalVisible(false);
  }


  const { data } = useRequest(() => {
    return queryFakeList();
  });

  const list = data || [];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <Button onClick={showEditModal}>修改</Button>
        ],
      }}
    >
      <Modal
        title="Modify OneTab"
        visible={isEditModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <EditGroup></EditGroup>
      </Modal>
      <div>
        <List
          grid={{
            gutter: 16,
            xs: 1,
            sm: 2,
            md: 4,
            lg: 4,
            xl: 6,
            xxl: 6,
          }}
          dataSource={list}
          renderItem={item => (
            <List.Item>
              <Card title={item.title}>
                <List
                  dataSource={item.urls}
                  renderItem={itemUrl => (
                    <List.Item>
                      <Image
                        width={16}
                        height={16}
                        src={itemUrl.icon}
                      />
                      <a href={itemUrl.url} target="_blank">{itemUrl.title}</a>
                    </List.Item>
                  )}
                >
                </List>
              </Card>
            </List.Item>
          )}
        />,
      </div>
    </PageContainer>
  );
};

export default CardList;
