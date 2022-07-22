import axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import config from "../config";
import storage from "@/utils/stored-data";

const URL = config.url + "api/";
export const CarouselUrl = URL + "v1/upload/carousel";

const service: AxiosInstance = axios.create({
    baseURL: URL,
    timeout: 5000,
    headers: {},
});

//请求拦截器
service.interceptors.request.use((config): AxiosRequestConfig => {
    Object.assign(config.headers ? config.headers : {}, { Authorization: `Bearer ${storage.get('admin')}` });
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