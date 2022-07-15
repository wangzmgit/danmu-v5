import axios, { AxiosInstance, AxiosRequestConfig } from "axios";
import config from "../config";
import storage from "@/utils/stored-data";

const URL = config.url + "api/";
export const MsgSocketURL = URL + 'v1/message/ws';

const service: AxiosInstance = axios.create({
    baseURL: URL,
    timeout: 5000,
    headers: {},
});

service.interceptors.request.use((config): AxiosRequestConfig => {
    Object.assign(config.headers ? config.headers : {}, { Authorization: `Bearer ${storage.get('token')}` });
    return config;
}), (error: any) => {
    return Promise.reject(error);
}

export default service;