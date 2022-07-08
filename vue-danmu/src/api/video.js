import request from '@/utils/request'

//上传视频信息
export const uploadVideoInfoAPI = ({ title, cover, desc, copyright, partition }) => {
  return request.post('v1/video/upload/info', { title, cover, desc, copyright, partition });
}

//修改视频信息
export const modifyVideoInfoAPI = ({ vid, title, cover, desc, copyright }) => {
  return request.post('v1/video/modify/info', { vid, title, cover, desc, copyright });
}

//修改资源标题
export const modifyTitleAPI = ({ id, title }) => {
  return request.post('v1/video/resource/title/modify', { id, title });
}

//获取推荐视频
export const getRecommendVideoAPI = () => {
  return request.get('v1/video/recommend/get');
}

//获取视频信息
export const getVideoInfoAPI = (vid) => {
  return request.get(`v1/video/get?vid=${vid}`);
}

//获取我的视频
export const getMyVideoAPI = (page, page_size) => {
  return request.get(`v1/video/upload/get?page=${page}&page_size=${page_size}`);
}

//通过用户ID获取视频
export const getVideoListByUidAPI = (uid, page, page_size) => {
  return request.get(`v1/video/user/get?uid=${uid}&page=${page}&page_size=${page_size}`);
}

//获取视频列表
export const getVideoListAPI = (page, page_size, partition) => {
  return request.get(`v1/video/list/get?page=${page}&page_size=${page_size}&partition=${partition}`);
}

//删除视频
export const deleteVideoAPI = (id) => {
  return request.post('v1/video/delete', { id });
}

//删除资源
export const deleteResourceAPI = (id) => {
  return request.post('v1/video/resource/delete', { id });
}

//获取视频状态
export const getResourceListAPI = (vid) => {
  return request.get(`v1/video/resource/list?vid=${vid}`);
}
