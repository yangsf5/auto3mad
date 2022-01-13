import { request } from 'umi';
import type { MemorialInfo, MemorialEditInfo } from './data';

export async function queryMemorialList(): Promise<{ data: MemorialInfo[] }> {
  return request('/v2/day/memorials', {
    params: { 'kind': 'full' },
  });
}

export async function queryEditList(): Promise<{ data: MemorialEditInfo[] }> {
  return request('/v2/day/memorials', {
    params: { 'kind': 'edit' },
  });
}

export async function upsertMemorial(info: MemorialEditInfo): Promise<{ data: any }> {
  return request('/v2/day/memorials', {
    method: 'post',
    data: info,
  });
}

export async function deleteMemorial(id: number): Promise<{ success: boolean }> {
  return request('/v2/day/memorials', {
    method: 'delete',
    params: { 'id': id },
  });
}

export async function queryMaxID(): Promise<{ data: number }> {
  return request('/v2/misc', {
    params: { 'kind': 'memorial' },
  });
}
