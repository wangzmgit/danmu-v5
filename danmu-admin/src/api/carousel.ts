import { addCarouselType, carouselType } from '@/types/carousel';
import request from '@/utils/request';

//获取轮播图
export const getCarouselAPI = () => {
  return request.get('v1/carousel/get');
}

//上传轮播图信息
export const addCarouselAPI = (carousel: addCarouselType) => {
  return request.post('v1/carousel/upload/info', carousel);
}

//删除轮播图
export const deleteCarouselAPI = (id: number) => {
  return request.post('v1/carousel/delete', { id });
}