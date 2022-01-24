import { request } from 'umi';
import type { RoutineInfo, EventInfo } from './data';

export async function queryRoutineList(): Promise<{ data: RoutineInfo[] }> {
  return request('/v2/daily/routines');
}

export async function upsertRoutine(info: RoutineInfo): Promise<{ data: any }> {
  return request('/v2/daily/routines', {
    method: 'post',
    data: info,
  });
}

export async function deleteRoutine(id: number): Promise<{ success: boolean }> {
  return request('/v2/daily/routines', {
    method: 'delete',
    params: { 'id': id },
  });
}

export async function queryEventList(date: string): Promise<{ data: EventInfo[] }> {
  return request('/v2/daily/events', {
    params: { 'date': date }
  });
}

export async function upsertEvent(info: EventInfo): Promise<{ data: any }> {
  return request('/v2/daily/events', {
    method: 'post',
    data: info,
  });
}

export async function deleteEvent(start_time: string): Promise<{ success: boolean }> {
  return request('/v2/daily/events', {
    method: 'delete',
    params: { 'start_time': start_time },
  });
}

