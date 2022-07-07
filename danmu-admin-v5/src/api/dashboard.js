import request from '@/utils/request';

//获取网站总体数据
export const getTotalDataAPI = () => {
  return request.get('v1/website/data/total');
}

//获取网站近期数据
export const getRecentDataAPI = () => {
  return request.get('v1/website/data/recent');
}
