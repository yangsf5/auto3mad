import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { GroupInfo } from './data';
import { queryGroupList } from './service';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

const EditGroup = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<GroupInfo[]>([]);
  const [position] = useState<'top' | 'bottom' | 'hidden'>('bottom');

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
      width: 100,
      render: (text, record, _, action) => [
        <a
          key="editable"
          onClick={() => {
            action?.startEditable?.(record.id);
          }}
        >
          编辑
        </a>,
        <a
          key="delete"
          onClick={() => {
            setDataSource(dataSource.filter((item) => item.id !== record.id));
          }}
        >
          删除
        </a>,
      ],
    },
  ];

  return (
    <>
      <EditableProTable<GroupInfo>
        rowKey="id"
        headerTitle="Groups"
        maxLength={10}
        recordCreatorProps={
          {
            position: position as 'top',
            record: () => ({ id: (Math.random() * 1000000).toFixed(0) }),
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
            console.log(rowKey, data, row);
            await waitTime(2000);
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
