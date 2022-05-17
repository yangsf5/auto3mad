import { request } from 'umi';
import type { StatRet } from './data';

export async function queryStat(firstMonth: string, lastMonth: string): Promise<{ data: StatRet }> {
  return request('/v2/daily/stats', {
    params: { 'first_month': firstMonth, 'last_month': lastMonth },
  });
}
