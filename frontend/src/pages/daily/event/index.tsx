import { useState, useRef } from 'react';
import { Modal, Button, DatePicker } from 'antd';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import { PageContainer } from '@ant-design/pro-layout';
import { EditableProTable } from '@ant-design/pro-table';
import { queryEventList, upsertEvent, deleteEvent, queryRoutineList } from './service';
import { RoutineTable, RoutineRun } from './routine';
import { EditRoutine } from './edit';
import { EventInfo } from './data';
import { useRequest } from 'umi';
import {
  GlobalOutlined,
  FireOutlined,
  DollarOutlined,
  BookOutlined,
  BulbOutlined,
  CodeOutlined,
  CoffeeOutlined,
  HomeOutlined,
  TrophyOutlined,
  GithubOutlined,
} from '@ant-design/icons';
import moment from 'moment';


export default () => {
  const [isEditModalVisible, setEditModalVisible] = useState(false);
  const showEditModal = () => {
    setEditModalVisible(true);
  };
  const onModalCancel = () => {
    setEditModalVisible(false);
    // RoutineRun();
  }

  const [queryDate, setQueryDate] = useState(moment());
  const onChangeDate = (m: any, _: string) => {
    setQueryDate(m);
  }

  const { data } = useRequest(() => {
    return queryRoutineList();
  });
  var groupOptions: { label: string; value: number; }[] = [];
  data?.forEach(val => groupOptions.push({ label: val.short_name, value: val.id }));

  const RoutineIcon = {
    1: <GlobalOutlined style={{ color: 'deepskyblue' }} />,
    2: <FireOutlined style={{ color: 'red' }} />,
    3: <DollarOutlined style={{ color: 'red' }} />,
    4: <BookOutlined style={{ color: 'green' }} />,
    6: <BulbOutlined style={{ color: 'orange' }} />,
    7: <CodeOutlined style={{ color: 'red' }} />,
    8: <CoffeeOutlined style={{ color: 'blue' }} />,
    9: <HomeOutlined style={{ color: 'hotpink' }} />,
    10: <TrophyOutlined style={{ color: 'tomato' }} />,
    11: <CodeOutlined style={{ color: 'red' }} />,
    12: <GithubOutlined />,
  }


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
      key: 'routine_avatar',
      dataIndex: 'routine_id',
      valueType: 'avatar',
      editable: false,
      width: 160,
      render: (_, row) => RoutineIcon[row.routine_id],
    },
    {
      title: '种类简称',
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
      width: 140,
      render: (text, record, _, action) => [
        <a
          key='editable'
          onClick={() => {
            action?.startEditable?.(record.start_time);
          }}
        >
          编辑
        </a>,
      ],
    },
  ];

  const ref = useRef<ActionType>();

  const [maxEndTime, setMaxEndTime] = useState('');
  const newEventInfo = () => {
    var event: EventInfo = {
      date: queryDate.format('YYYY-MM-DD'),
      start_time: maxEndTime,
      end_time: '',
      specific_event: '',
      routine_id: 6,
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

      {/* <RoutineTable /> */}

      <EditableProTable<EventInfo>
        size='small'
        rowKey='start_time'
        recordCreatorProps={
          {
            position: 'top',
            record: newEventInfo(),
          }
        }
        columns={columns}
        params={queryDate}
        request={async () => {
          const { data } = await queryEventList(queryDate.format('YYYY-MM-DD'));
          setMaxEndTime(data.max_end_time);
          return {
            data: data.events,
            success: true,
          };
        }}
        actionRef={ref}
        value={dataSource}
        onChange={setDataSource}
        editable={{
          type: 'multiple',
          editableKeys,
          onSave: async (rowKey, data, row) => {
            await upsertEvent(data);
            ref.current.reload();
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
