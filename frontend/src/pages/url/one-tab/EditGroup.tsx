import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { GroupInfo } from './data';
import { queryGroupList, upsertGroup, deleteGroup, queryMaxID } from './service';
import { useRequest } from 'umi';

const EditGroup = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<GroupInfo[]>([]);

  const columns: ProColumns<GroupInfo>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      editable: () => false,
    },
    {
      title: '分组',
      dataIndex: 'title',
    },
    {
      title: '图标',
      dataIndex: 'icon',
      valueType: 'avatar',
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

  const max = useRequest(() => {
    return queryMaxID("group");
  });

  return (
    <>
      <EditableProTable<GroupInfo>
        rowKey="id"
        maxLength={10}
        recordCreatorProps={
          {
            position: 'bottom',
            record: { id: max.data + 1, title: '', icon: '' },
          }
        }
        columns={columns}
        request={async () => {
          const { data } = await queryGroupList();
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
            await upsertGroup(data);
            max.run();
          },
          onDelete: async (rowKey, data) => {
            await deleteGroup(data.id);
          },
          onChange: setEditableRowKeys,
        }}
      />
    </>
  );
};

export {
  EditGroup,
}
