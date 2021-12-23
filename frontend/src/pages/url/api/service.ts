import { request } from 'umi';
import type { ApiInfo } from './data';

export async function queryApiList(): Promise<{ data: ApiInfo[] }> {
  return request('/v2/url/apis');
}
