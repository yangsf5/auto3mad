import { useState } from 'react';
import { Card, List, Image, Button, Modal } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryUrlList } from './service';
import { EditGroup } from './EditGroup';
import { EditItem } from './EditItem';

const CardList = () => {
  const [isEditGroupModalVisible, setEditGroupModalVisible] = useState(false);
  const showEditGroupModal = () => {
    setEditGroupModalVisible(true);
  };
  const onGroupModalCancel = () => {
    setEditGroupModalVisible(false);
  }

  const [isEditItemModalVisible, setEditItemModalVisible] = useState(false);
  const showEditItemModal = () => {
    setEditItemModalVisible(true);
  };
  const onItemModalCancel = () => {
    setEditItemModalVisible(false);
  }

  const { data } = useRequest(() => {
    return queryUrlList();
  });

  const list = data || [];

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <Button onClick={showEditGroupModal}>Edit Group</Button>,
          <Button onClick={showEditItemModal}>Edit Item</Button>
        ],
      }}
    >
      <Modal
        title="Modify URL Group"
        visible={isEditGroupModalVisible}
        onCancel={onGroupModalCancel}
        footer={null}
        width={1000}
        destroyOnClose={true}
      >
        <EditGroup></EditGroup>
      </Modal>
      <Modal
        title="Modify URL Item"
        visible={isEditItemModalVisible}
        onCancel={onItemModalCancel}
        footer={null}
        width={1000}
        destroyOnClose={true}
      >
        <EditItem></EditItem>
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
