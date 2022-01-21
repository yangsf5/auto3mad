import { useState } from 'react';
import { Modal, Button } from 'antd';
import type { ProColumns } from '@ant-design/pro-table';
import { PageContainer } from '@ant-design/pro-layout';
import { EditableProTable } from '@ant-design/pro-table';
import { queryEventList, upsertEvent, deleteEvent, queryRoutineList } from './service';
import { RoutineTable, RoutineRun } from './routine';
import { EditRoutine } from './edit';
import { EventInfo } from './data';
import { useRequest } from 'umi';



export default () => {
  const [isEditModalVisible, setEditModalVisible] = useState(false);
  const showEditModal = () => {
    setEditModalVisible(true);
  };
  const onModalCancel = () => {
    setEditModalVisible(false);
    RoutineRun();
  }

  const { data } = useRequest(() => {
    return queryRoutineList();
  });
  var groupOptions: { label: string; value: number; }[] = [];
  data?.forEach(val => groupOptions.push({ label: val.short_name, value: val.id }));


  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<EventInfo[]>([]);
  const columns: ProColumns<EventInfo>[] = [
    {
      title: '日期',
      dataIndex: 'date',
      width: 120,
    },
    {
      title: '开始时间',
      dataIndex: 'start_time',
      width: 100,
    },
    {
      title: '结束时间',
      dataIndex: 'end_time',
      width: 100,
    },
    {
      title: '具体事件',
      dataIndex: 'specific_event',
    },
    {
      title: '例行种类',
      dataIndex: 'routine_id',
      valueType: 'select',
      fieldProps: { options: groupOptions },
      width: 160,
    },
    {
      title: '投入分钟',
      dataIndex: 'spend',
      width: 100,
    },
    {
      title: '操作',
      valueType: 'option',
      width: 140,
      render: (text, record, _, action) => [
        <a
          key="editable"
          onClick={() => {
            action?.startEditable?.(record.start_time);
          }}
        >
          编辑
        </a>,
      ],
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

      {/* <RoutineTable /> */}

      <EditableProTable<EventInfo>
        size='small'
        rowKey="start_time"
        recordCreatorProps={
          {
            position: 'top',
            record: { start_time: '', end_time: '', specific_event: '', routine_event: '', spend: 0 },
          }
        }
        columns={columns}
        request={async () => {
          const { data } = await queryEventList();
          return {
            data: data,
            success: true,
          };
        }}
        value={dataSource}
        onChange={setDataSource}
        editable={{
          type: 'multiple',
          editableKeys,
          onSave: async (rowKey, data, row) => {
            await upsertEvent(data);
          },
          onDelete: async (rowKey, data) => {
            await deleteEvent(data.start_time);
          },
          onChange: setEditableRowKeys,
        }}
      />
    </PageContainer>
  );
};
