import { request } from 'umi';
import type { MemorialInfo } from './data';

export async function queryMemorialList(): Promise<{ data: MemorialInfo[] }> {
  return request('/v2/day/memorials');
}
