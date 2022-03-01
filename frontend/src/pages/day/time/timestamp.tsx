import { useState } from 'react';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import { queryTimestamp } from './service';
import { TimestampInfo } from './data';


const Timestamp = () => {
  const [timestamp, setTimestamp] = useState<TimestampInfo>();

  return (
    <>
      <ProForm<{ date: string }>
        layout={'inline'}
        onFinish={async (values) => {
          const { data } = await queryTimestamp(values.date);
          setTimestamp(data);
        }}
      >
        <ProFormText name='date' label='Date' />
      </ProForm>

      <br />
      <div>First Second: {timestamp?.first_second}</div>
      <div>Last Second: {timestamp?.last_second}</div>
    </>
  );
};

export {
  Timestamp,
};
