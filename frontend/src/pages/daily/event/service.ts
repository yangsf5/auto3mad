import { request } from 'umi';
import type { RoutineInfo, EventInfo, EventAPI } from './data';

export async function queryRoutineList(date: string): Promise<{ data: RoutineInfo[] }> {
  return request('/v2/daily/routines', {
    params: { 'date': date }
  });
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

export async function queryEventList(date: string): Promise<{ data: EventAPI }> {
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

export async function deleteEvent(date: string, id: number): Promise<{ success: boolean }> {
  return request('/v2/daily/events', {
    method: 'delete',
    params: { 'id': id },
  });
}

