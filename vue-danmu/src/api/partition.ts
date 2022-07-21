import request from '@/utils/request'

//获取分区
export const getPartitionAPI = (parentId: number) => {
    return request.get(`v1/partition/list?parent_id=${parentId}`);
}

//获取所有
export const getAllPartitionAPI = () => {
    return request.get('v1/partition/all');
}