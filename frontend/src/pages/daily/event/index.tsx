import { useState, useRef } from 'react';
import { Modal, Button, DatePicker, Avatar } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { PageContainer } from '@ant-design/pro-layout';
import { EditableProTable } from '@ant-design/pro-table';
import { queryEventList, upsertEvent, deleteEvent, queryRoutineList } from './service';
import { RoutineTable, RoutineSelect } from './routine';
import { EditRoutine } from './edit';
import { EventInfo } from './data';
import { useRequest } from 'umi';
import moment from 'moment';


export default () => {
  // Routine 的编辑模态框
  const [isEditModalVisible, setEditModalVisible] = useState(false);
  const showEditModal = () => {
    setEditModalVisible(true);
  };
  const onModalCancel = () => {
    setEditModalVisible(false);
    refreshRoutine();
  };

  // 日期选择框
  const [queryDate, setQueryDate] = useState(moment());
  const onChangeDate = (m: moment.Moment, _: string) => {
    setQueryDate(m);
  };

  // 引用 EventTable 的动作
  const refEventTableAction = useRef<ActionType>();

  // Routine 的刷新标记：RoutineTable、EventTable
  const [refreshRoutineTable, setRefreshRoutineTable] = useState(1);
  const refreshRoutine = () => {
    setRefreshRoutineTable(refreshRoutineTable + 1);
    // 刷新 EventTable 里的 Routine 内容
    refEventTableAction.current?.reload();
  };

  const { data } = useRequest(() => queryRoutineList(queryDate.format('YYYY-MM-DD')), { refreshDeps: [queryDate, refreshRoutineTable] });

  // 设置 Routine Map，方便找 event_default
  var routineMap = new Map();

  // 给 EventTable 设置 Routine 类型选择器
  var groupOptions: { label: any; value: number; }[] = [];

  data?.forEach(val => {
    routineMap.set(val.id, val);

    groupOptions.push({
      label: <div><Avatar size={16} src={val.icon}></Avatar> {val.short_name}</div>,
      value: val.id,
    });
  });

  const [editableKeys, setEditableRowKeys] = useState<React.Key[]>([]);
  const [dataSource, setDataSource] = useState<EventInfo[]>();
  const columns: ProColumns<EventInfo>[] = [
    {
      title: '日期',
      dataIndex: 'date',
      editable: false,
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
      title: '日拱具体事件',
      dataIndex: 'specific_event',
    },
    {
      title: '日拱项',
      key: 'routine_id',
      dataIndex: 'routine_id',
      valueType: 'select',
      fieldProps: { options: groupOptions },
      width: 160,
    },
    {
      title: '投入分钟',
      dataIndex: 'spend',
      editable: false,
      width: 100,
    },
    {
      title: '操作',
      valueType: 'option',
      width: 200,
      render: (text, record, _, action) => [
        <a key='editable' onClick={() => action?.startEditable?.(record.start_time)}>编辑</a>,
        <a key='refresh_end_time' onClick={async () => {
          record.end_time = moment().format('HH:mm');
          await upsertEvent(record);
          refreshRoutine();
        }}>刷新结束时间</a>,
      ],
    },
  ];

  const newEvent = (routineID: number) => {
    var event: EventInfo = {
      date: queryDate.format('YYYY-MM-DD'),
      start_time: moment().format('HH:mm'),
      end_time: moment().format('HH:mm'),
      specific_event: routineMap.get(routineID)?.event_default,
      routine_id: routineID,
      spend: 0,
    };
    return event;
  };

  return (
    <PageContainer
      header={{
        title: "",
        breadcrumb: {},
        extra: [
          <DatePicker key='date' defaultValue={queryDate} onChange={onChangeDate} />,
          <Button key='edit' onClick={showEditModal}>Edit</Button>
        ],
      }}
    >
      <Modal
        title='Edit Routine'
        visible={isEditModalVisible}
        onCancel={onModalCancel}
        footer={null}
        width={1000}
        destroyOnClose={true}
      >
        <EditRoutine></EditRoutine>
      </Modal>

      <RoutineTable dataSource={data} />

      <p />

      <EditableProTable<EventInfo>
        rowKey='start_time'
        size='small'
        recordCreatorProps={false}
        toolBarRender={() => {
          return [
            <div>开始一项日拱</div>,
            <RoutineSelect
              dataSource={data}
              onChange={(val: number) => {
                refEventTableAction.current?.addEditRecord?.(newEvent(val), { position: 'top' });
              }}
            />
          ];
        }}
        columns={columns}
        params={queryDate}
        request={async () => {
          const { data } = await queryEventList(queryDate.format('YYYY-MM-DD'));
          return {
            data: data.events,
            success: true,
          };
        }}
        actionRef={refEventTableAction}
        value={dataSource}
        onChange={setDataSource}
        editable={{
          type: 'multiple',
          editableKeys,
          onSave: async (rowKey, data, row) => {
            await upsertEvent(data);
            refreshRoutine();
          },
          onDelete: async (rowKey, data) => {
            await deleteEvent(data.date, data.start_time);
          },
          onChange: setEditableRowKeys,
        }}
      />

    </PageContainer>
  );
};
