import { request } from 'umi';
import type { UrlGroup } from './data';

export async function queryFakeList(): Promise<{ data: UrlGroup[] }> {
  return request('/v2/url/urls');
}
