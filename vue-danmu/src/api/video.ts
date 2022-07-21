import { baseResourceType } from '@/types/resource';
import { modifyVideoType, uploadVideoType } from '@/types/video';
import request from '@/utils/request'

//上传视频信息
export const uploadVideoInfoAPI = (uploadVideo: uploadVideoType) => {
  return request.post('v1/video/upload/info', uploadVideo);
}

//修改视频信息
export const modifyVideoInfoAPI = (modifyVideo: modifyVideoType) => {
  return request.post('v1/video/modify/info', modifyVideo);
}

//修改资源标题
export const modifyTitleAPI = (resourceTitle: baseResourceType) => {
  return request.post('v1/video/resource/title/modify', resourceTitle);
}

//获取推荐视频
export const getRecommendVideoAPI = () => {
  return request.get('v1/video/recommend/get');
}

//获取视频信息
export const getVideoInfoAPI = (vid: number) => {
  return request.get(`v1/video/get?vid=${vid}`);
}

//获取我的视频
export const getMyVideoAPI = (page: number, page_size: number) => {
  return request.get(`v1/video/upload/get?page=${page}&page_size=${page_size}`);
}

//通过用户ID获取视频
export const getVideoListByUidAPI = (uid: number, page: number, page_size: number) => {
  return request.get(`v1/video/user/get?uid=${uid}&page=${page}&page_size=${page_size}`);
}

//获取视频列表
export const getVideoListAPI = (page: number, page_size: number, partition: number) => {
  return request.get(`v1/video/list/get?page=${page}&page_size=${page_size}&partition=${partition}`);
}

//删除视频
export const deleteVideoAPI = (id: number) => {
  return request.post('v1/video/delete', { id });
}

//删除资源
export const deleteResourceAPI = (id: number) => {
  return request.post('v1/video/resource/delete', { id });
}

//获取视频状态
export const getResourceListAPI = (vid: number) => {
  return request.get(`v1/video/resource/list?vid=${vid}`);
}
