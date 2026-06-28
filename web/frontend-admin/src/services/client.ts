import type { AxiosRequestConfig } from 'axios';
import { request } from './request';

// 导出一个通用函数供 generated 代码调用
export const customClient = <T>(config: AxiosRequestConfig): Promise<T> => {
  return request(config);
};

export default customClient;