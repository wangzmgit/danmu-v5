import request from '@/utils/request';

//添加分区
export const addPartitionAPI = ({ content, parentId }) => {
    return request.post('v1/partition/add', { content, parentId });
}

//删除分区
export const deletePartitionAPI = (id) => {
    return request.post('v1/partition/delete', { id });
}

//获取所有
export const getAllPartitionAPI = () => {
    return request.get('v1/partition/all');
}