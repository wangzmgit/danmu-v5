import request from '@/utils/request';

//获取收藏夹列表
export const getCollectionListAPI = () => {
    return request.get('v1/archive/collection/list');
}

//创建收藏夹
export const addCollectionAPI = (name) => {
    return request.post('v1/archive/collection/add', { name });
}

//修改收藏夹
export const modifyCollectionAPI = ({ id, cover, name, desc, open }) => {
    return request.post('v1/archive/collection/modify', { id, cover, name, desc, open });
}

//获取收藏夹信息(非所有者仅获取公开的)
export const getCollectionInfoAPI = (id) => {
    return request.get(`v1/archive/collection/info?id=${id}`);
}

//删除收藏夹
export const deleteCollectionAPI = (id) => {
    return request.post('v1/archive/collection/delete', { id });
}

//获取收藏状态
export const getCollectInfoAPI = (vid) => {
    return request.get(`v1/archive/collect/status?vid=${vid}`);
}

//收藏
export const collectAPI = (vid, addList, cancelList) => {
    return request.post('v1/archive/collect/add', { vid, addList, cancelList });
}

//获取收藏夹内的视频
export const getCollectVideoAPI = (id, page, page_size) => {
    return request.get(`v1/archive/collect/get?id=${id}&page=${page}&page_size=${page_size}`);
}


