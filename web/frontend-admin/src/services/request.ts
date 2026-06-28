import axios from 'axios';

// 1. 创建统一请求实例
export const request = axios.create({
  baseURL: '/api', // 后续可在环境变量中配置
  timeout: 10000,
});

// 2. 请求拦截器：注入鉴权与公共 Header
request.interceptors.request.use((config) => {
  // TODO: 从 session.ts 获取 Token 并注入 Authorization
  // TODO: 注入 App-Name, Timestamp 等
  return config;
});

// 3. 响应拦截器：解包与错误处理
request.interceptors.response.use(
  (response) => {
    // TODO: 判断 response.data.code，若不为 0 则抛出错误
    // 页面层默认只接收解包后的业务 data
    return response.data.data;
  },
  (error) => {
    // TODO: 统一错误提示、处理 401 登出逻辑
    return Promise.reject(error);
  }
);