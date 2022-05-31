import { request } from 'umi';
import type { UrlGroup, GroupInfo, ItemInfo } from './data';

export async function queryUrlList(): Promise<{ data: UrlGroup[] }> {
  return request('/v2/url/urls');
}

export async function queryGroupList(): Promise<{ data: GroupInfo[] }> {
  return request('/v2/url/groups');
}

export async function upsertGroup(info: GroupInfo): Promise<{ data: any }> {
  return request('/v2/url/groups', {
    method: 'post',
    data: info,
  });
}

export async function deleteGroup(id: number): Promise<{ success: boolean }> {
  return request('/v2/url/groups', {
    method: 'delete',
    params: { "id": id },
  });
}

export async function queryItemList(): Promise<{ data: ItemInfo[] }> {
  return request('/v2/url/items');
}

export async function upsertItem(info: ItemInfo): Promise<{ data: any }> {
  return request('/v2/url/items', {
    method: 'post',
    data: info,
  });
}

export async function deleteItem(id: number): Promise<{ success: boolean }> {
  return request('/v2/url/items', {
    method: 'delete',
    params: { "id": id },
  });
}
