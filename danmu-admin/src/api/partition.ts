import { addPartitionType } from '@/types/partition';
import request from '@/utils/request';

//添加分区
export const addPartitionAPI = (partition: addPartitionType) => {
    return request.post('v1/partition/add', partition);
}

//删除分区
export const deletePartitionAPI = (id: number) => {
    return request.post('v1/partition/delete', { id });
}

//获取所有
export const getAllPartitionAPI = () => {
    return request.get('v1/partition/all');
}