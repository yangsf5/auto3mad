import { request } from 'umi';
import type { UrlGroup } from './data';

export async function queryUrlList(): Promise<{ data: UrlGroup[] }> {
  return request('/v2/url/urls');
}
