import { ref } from 'vue';

export default function useConfig() {
    const getConfig = () => {
        let config = localStorage.getItem("player-config");
        if (!config) {
            return initConfig();
        }
        return JSON.parse(config);
    }

    const setConfig = (key, value) => {
        let config = JSON.parse(localStorage.getItem("player-config"));
        config[key] = value;
        localStorage.setItem("player-config", JSON.stringify(config));
    }

    const initConfig = () => {
        let config = {
            defaultRes: 720,//默认分辨率
            danmaku: true,
            volume: 80,
        }
        localStorage.setItem("player-config", JSON.stringify(config));
        return config;
    }

    return {
        getConfig,
        setConfig
    }
}

