import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { RoutineInfo } from './data';
import { upsertRoutine, deleteRoutine, queryRoutineList } from './service';

const EditRoutine = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<RoutineInfo[]>([]);

  const columns: ProColumns<RoutineInfo>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      editable: () => false,
      width: 50,
    },
    {
      title: '图标',
      dataIndex: 'icon',
      valueType: 'avatar',
      width: 120,
    },
    {
      title: '简称',
      dataIndex: 'short_name',
      width: 120,
    },
    {
      title: '例行事件内容',
      dataIndex: 'event',
    },
    {
      title: '预算 M',
      dataIndex: 'will_spend',
      valueType: 'digit',
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
            action?.startEditable?.(record.id);
          }}
        >
          编辑
        </a>,
      ],
    },
  ];

  return (
    <>
      <EditableProTable<RoutineInfo>
        rowKey="id"
        recordCreatorProps={
          {
            position: 'bottom',
            record: { id: 0, icon: '', short_name: '', event: '', will_spend: 0, today_spend: 0, total_spend: 0 },
          }
        }
        columns={columns}
        request={async () => {
          const { data } = await queryRoutineList("");
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
            await upsertRoutine(data);
          },
          onDelete: async (rowKey, data) => {
            await deleteRoutine(data.id);
          },
          onChange: setEditableRowKeys,
        }}
      />
    </>
  );
};

export {
  EditRoutine,
};
