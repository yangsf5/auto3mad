import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';
import { MemorialEditInfo } from './data';
import { upsertMemorial, deleteMemorial, queryMaxID, queryEditList } from './service';
import { useRequest } from 'umi';

const EditMerial = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<MemorialEditInfo[]>([]);

  const columns: ProColumns<MemorialEditInfo>[] = [
    {
      title: 'ID',
      dataIndex: 'id',
      editable: () => false,
    },
    {
      title: '纪念',
      dataIndex: 'desc',
    },
    {
      title: '日期',
      dataIndex: 'date',
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
    return queryMaxID();
  });

  return (
    <>
      <EditableProTable<MemorialEditInfo>
        rowKey="id"
        maxLength={10}
        recordCreatorProps={
          {
            position: 'bottom',
            record: { id: max.data + 1, desc: '', date: '' },
          }
        }
        columns={columns}
        request={async () => {
          const { data } = await queryEditList();
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
            await upsertMemorial(data);
            max.run();
          },
          onDelete: async (rowKey, data) => {
            await deleteMemorial(data.id);
          },
          onChange: setEditableRowKeys,
        }}
      />
    </>
  );
};

export {
  EditMerial,
};
