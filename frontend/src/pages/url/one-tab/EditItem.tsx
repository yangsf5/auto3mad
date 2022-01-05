import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { ItemInfo } from './data';
import { queryItemList, upsertItem, deleteItem, queryGroupList } from './service';
import { useRequest } from 'umi';

const EditItem = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<ItemInfo[]>([]);

  const { data } = useRequest(() => {
    return queryGroupList();
  });
  var groupOptions: { label: string; value: number; }[] = [];
  data?.forEach(val => groupOptions.push({ label: val.title, value: val.id }));

  const defaultGroup: number = groupOptions?.length > 0 ? groupOptions[0].value : -1;

  const columns: ProColumns<ItemInfo>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      width: 50,
      editable: () => false,
    },
    {
      title: '分组',
      dataIndex: 'group_id',
      valueType: 'select',
      fieldProps: { options: groupOptions },
      width: 110,
    },
    {
      title: '图标',
      dataIndex: 'icon',
      valueType: 'avatar',
      width: 250,
    },
    {
      title: '链接名称',
      dataIndex: 'title',
      width: 120,
    },
    {
      title: 'URL',
      dataIndex: 'url',
      render: text => <a href={`${text}`} target='_blank'>{text}</a>,
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
      <EditableProTable<ItemInfo>
        rowKey="id"
        maxLength={50}
        recordCreatorProps={
          {
            position: 'bottom',
            record: () => ({ id: -1, title: '', icon: '', url: '', group_id: defaultGroup }),
          }
        }
        columns={columns}
        request={async () => {
          const { data } = await queryItemList();
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
            await upsertItem(data);
          },
          onDelete: async (rowKey, data) => {
            await deleteItem(data.id);
          },
          onChange: setEditableRowKeys,
        }}
      />
    </>
  );
};

export {
  EditItem,
}
