import axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import config from "../config";
import storage from "@/utils/stored-data";

const URL = config.url + "api/";
export const CoverUrl = URL + "v1/upload/cover";
export const VideoUrl = URL + "v1/upload/video";
export const AvatarUrl = URL + "v1/upload/avatar";
export const MsgSocketURL = URL + 'v1/message/ws';
export const OnlineSocketURL = URL + 'v1/video/online/ws';

const service: AxiosInstance = axios.create({
    baseURL: URL,
    timeout: 5000,
    headers: {},
});

//请求拦截器
service.interceptors.request.use((config): AxiosRequestConfig => {
    Object.assign(config.headers ? config.headers : {}, { Authorization: `Bearer ${storage.get('token')}` });
    return config;
}), (error: any) => {
    return Promise.reject(error);
}


//响应拦截器
service.interceptors.response.use((res) => {
    return res;
}, (error) => {
    return Promise.reject(error);
});

export default service;