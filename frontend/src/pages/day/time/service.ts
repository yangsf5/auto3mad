import { request } from 'umi';
import type { TimeInfo, TimestampInfo } from './data';

export async function queryTimeList(timestamp: number): Promise<{ data: TimeInfo[] }> {
  return request('/v2/day/timestamp', {
    params: { 'type': 't2d', 'timestamp': timestamp },
  });
}

export async function queryTimestamp(date: string): Promise<{ data: TimestampInfo }> {
  return request('/v2/day/timestamp', {
    params: { 'type': 'd2t', 'date': date },
  });
}
