import axios from "axios";
import config from "../config";
import storage from "@/utils/stored-data.js";

const URL = config.url + "api/";
export const CarouselUrl = URL + "v1/upload/carousel";

const service = axios.create({
    baseURL: URL,
    timeout: 5000,
    headers: {},
});

service.interceptors.request.use((config) => {
    Object.assign(config.headers, { Authorization: `Bearer ${storage.get('admin')}` });
    return config;
}), (error) => {
    return Promise.reject(error);
}

export default service;