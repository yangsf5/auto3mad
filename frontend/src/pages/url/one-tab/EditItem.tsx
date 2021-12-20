import React, { useState } from 'react';
import type { ProColumns } from '@ant-design/pro-table';
import { EditableProTable } from '@ant-design/pro-table';

const waitTime = (time: number = 100) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(true);
    }, time);
  });
};

type DataSourceType = {
  id: React.Key;
  group_id: number;
  group_title: string;
  icon: string;
  url: string;
  title: string;
};

const defaultData: DataSourceType[] = [
  {
    id: 1000,
    group_id: 1,
    group_title: 'Code',
    icon: 'https://github.githubassets.com/favicons/favicon.png',
    url: 'https://github.com/yangsf5',
    title: 'GitHub 3Mad',
  },
  {
    id: 1001,
    group_id: 1,
    group_title: 'Code',
    icon: 'https://beego.vip/static/img/favicon.png',
    url: 'https://beego.vip/docs/intro/',
    title: 'Beego',
  },
];

const EditGroup = () => {
  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<DataSourceType[]>([]);
  const [position] = useState<'top' | 'bottom' | 'hidden'>('bottom');

  const columns: ProColumns<DataSourceType>[] = [
    {
      title: '分组',
      dataIndex: 'group_title',
    },
    {
      title: '图标',
      dataIndex: 'icon',
      valueType: 'avatar',
    },
    {
      title: '链接名称',
      dataIndex: 'title',
      formItemProps: (form, { rowIndex }) => {
        return {
          rules: [{ required: true, message: '此项为必填项' }],
        };
      },
    },
    {
      title: 'URL',
      dataIndex: 'url',
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
      <EditableProTable<DataSourceType>
        rowKey="id"
        headerTitle="URLs"
        maxLength={5}
        recordCreatorProps={
          position !== 'hidden'
            ? {
              position: position as 'top',
              record: () => ({ id: (Math.random() * 1000000).toFixed(0) }),
            }
            : false
        }
        columns={columns}
        request={async () => ({
          data: defaultData,
          total: 3,
          success: true,
        })}
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
