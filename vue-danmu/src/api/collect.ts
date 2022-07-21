import { collectType, modifyCollectionType } from '@/types/collect';
import request from '@/utils/request';

//获取收藏夹列表
export const getCollectionListAPI = () => {
    return request.get('v1/archive/collection/list');
}

//创建收藏夹
export const addCollectionAPI = (name: string) => {
    return request.post('v1/archive/collection/add', { name });
}

//修改收藏夹
export const modifyCollectionAPI = (collect: modifyCollectionType) => {
    return request.post('v1/archive/collection/modify', collect);
}

//获取收藏夹信息(非所有者仅获取公开的)
export const getCollectionInfoAPI = (id: number) => {
    return request.get(`v1/archive/collection/info?id=${id}`);
}

//删除收藏夹
export const deleteCollectionAPI = (id: number) => {
    return request.post('v1/archive/collection/delete', { id });
}

//获取收藏状态
export const getCollectInfoAPI = (vid: number) => {
    return request.get(`v1/archive/collect/status?vid=${vid}`);
}

//收藏
export const collectAPI = (collect: collectType) => {
    return request.post('v1/archive/collect/add', collect);
}

//获取收藏夹内的视频
export const getCollectVideoAPI = (id: number, page: number, page_size: number) => {
    return request.get(`v1/archive/collect/get?id=${id}&page=${page}&page_size=${page_size}`);
}


