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
export const setDatabaseAPI = ({ dbName, dbHost, dbPort, dbUser, dbPwd, cacheHost, cachePort, cachePwd }) => {
    return request.post('v1/config/database/set', { dbName, dbHost, dbPort, dbUser, dbPwd, cacheHost, cachePort, cachePwd });
}

//获取邮箱
export const getEmailAPI = () => {
    return request.get('v1/config/email/get');
}

//修改邮箱
export const setEmailAPI = ({ open, name, host, port, address, password }) => {
    return request.post('v1/config/email/set', { open, name, host, port, address, password });
}

//获取邮箱
export const getStorageAPI = () => {
    return request.get('v1/config/oss/get');
}

//修改邮箱
export const setStorageAPI = ({
    https, maxImgSize, maxVideoSize, storage, domain, bucket, endpoint, accesskeyId, accesskeySecret
}) => {
    return request.post('v1/config/oss/set', {
        https, maxImgSize, maxVideoSize, storage, domain, bucket, endpoint, accesskeyId, accesskeySecret
    });
}

//获取其他配置
export const getOtherAPI = () => {
    return request.get('v1/config/other/get');
}

//修改其他配置
export const setOtherAPI = ({ prefix, maxRes, fePath, email, password }) => {
    return request.post('v1/config/other/set', { prefix, maxRes, fePath, email, password });
}