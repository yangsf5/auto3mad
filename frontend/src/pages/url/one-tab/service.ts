import { request } from 'umi';
import type { GroupInfo, UrlGroup } from './data';

export async function queryUrlList(): Promise<{ data: UrlGroup[] }> {
  return request('/v2/url/urls');
}

export async function queryGroupList(): Promise<{ data: GroupInfo[] }> {
  return request('/v2/url/groups');
}
