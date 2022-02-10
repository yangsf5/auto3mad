import { useState, useRef } from 'react';
import { Row, Col, Card, Progress, Modal, Button, DatePicker, Avatar } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { PageContainer } from '@ant-design/pro-layout';
import { EditableProTable } from '@ant-design/pro-table';
import { queryEventList, upsertEvent, deleteEvent, queryRoutineList } from './service';
import { RoutineTable } from './routine';
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

  // 给 Event 设置 Routine 类型选择器
  const { data } = useRequest(() => {
    return queryRoutineList(queryDate.format('YYYY-MM-DD'));
  });
  var groupOptions: { label: any; value: number; }[] = [];
  data?.forEach(val => groupOptions.push({
    label: <div><Avatar size={16} src={val.icon}></Avatar> {val.short_name}</div>,
    value: val.id,
  }));

  // 通知 RoutineTable 刷新、EventTable 刷新
  const [refreshRoutineTable, setRefreshRoutineTable] = useState(1);
  const refreshRoutine = () => {
    setRefreshRoutineTable(refreshRoutineTable + 1);
    refEventTableAction.current?.reload();
  };

  const refEventTableAction = useRef<ActionType>();

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
      title: '具体事件',
      dataIndex: 'specific_event',
    },
    {
      title: '种类',
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

  const newEventInfo = () => {
    var event: EventInfo = {
      date: queryDate.format('YYYY-MM-DD'),
      start_time: moment().format('HH:mm'),
      end_time: moment().format('HH:mm'),
      specific_event: '',
      routine_id: 12,
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

      <Row gutter={16}>
        <Col span={12}>
          <Card>
            <Progress percent={222 / 720 * 100} success={{ percent: 30 }} type='circle'></Progress>
          </Card>
        </Col>
        <Col span={12}>
          <RoutineTable date={queryDate.format('YYYY-MM-DD')} refresh={refreshRoutineTable} />
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <Card>
            <EditableProTable<EventInfo>
              rowKey='start_time'
              recordCreatorProps={
                {
                  position: 'top',
                  record: newEventInfo(),
                  creatorButtonText: '开始一项日拱',
                }
              }
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
          </Card>
        </Col>
      </Row>
    </PageContainer>
  );
};
