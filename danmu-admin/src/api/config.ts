import { databaseType, emailType, otherType, storageType } from '@/types/config';
import request from '@/utils/request';

//获取服务数据
export const getServerDataAPI = () => {
    return request.get('v1/config/server/info');
}

//获取数据库
export const getDatabaseAPI = () => {
    return request.get('v1/config/database/get');
}

//修改数据库
export const setDatabaseAPI = (db: databaseType) => {
    return request.post('v1/config/database/set', db);
}

//获取邮箱
export const getEmailAPI = () => {
    return request.get('v1/config/email/get');
}

//修改邮箱
export const setEmailAPI = (email: emailType) => {
    return request.post('v1/config/email/set', email);
}

//获取邮箱
export const getStorageAPI = () => {
    return request.get('v1/config/storage/get');
}

//修改邮箱
export const setStorageAPI = (storage: storageType) => {
    return request.post('v1/config/storage/set', storage);
}

//获取其他配置
export const getOtherAPI = () => {
    return request.get('v1/config/other/get');
}

//修改其他配置
export const setOtherAPI = (other: otherType) => {
    return request.post('v1/config/other/set', other);
}