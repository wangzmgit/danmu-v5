import request from '@/utils/request'

//获取我的视频
export const getMyVideoAPI = (page: number, page_size: number) => {
  return request.get(`v1/video/upload/get?page=${page}&page_size=${page_size}`);
}

//获取视频列表
export const getVideoListAPI = (page: number, page_size: number, partition: number) => {
  return request.get(`v1/video/list/get?page=${page}&page_size=${page_size}&partition=${partition}`);
}

//获取视频信息
export const getVideoInfoAPI = (vid: number) => {
  return request.get(`v1/video/get?vid=${vid}`);
}