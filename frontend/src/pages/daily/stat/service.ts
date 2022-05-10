import { request } from 'umi';
import type { StatInfo } from './data';

export async function queryStat(firstMonth: string, lastMonth: string): Promise<{ data: StatInfo[] }> {
  return request('/v2/daily/stats', {
    params: { 'first_month': firstMonth, 'last_month': lastMonth },
  });
}
