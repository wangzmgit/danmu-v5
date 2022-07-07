import request from '@/utils/request'

//获取轮播图
export const getCarouselAPI = () => {
    return request.get('v1/carousel/get');
}
