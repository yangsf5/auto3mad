import { request } from 'umi';
import type { TimeInfo } from './data';

export async function queryTimeList(timestamp: number): Promise<{ data: TimeInfo[] }> {
  return request('/v2/day/timestamp', {
    params: { 'timestamp': timestamp },
  });
}
