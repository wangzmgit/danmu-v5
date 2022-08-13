export default function useConfig() {
    const getConfig = () => {
        const config = localStorage.getItem("player-config");
        if (!config) {
            return initConfig();
        }
        return JSON.parse(config);
    }

    const setConfig = (key: string, value: any) => {
        const readConfig = localStorage.getItem("player-config");
        const config = JSON.parse(readConfig ? readConfig : "{}");
        config[key] = value;
        localStorage.setItem("player-config", JSON.stringify(config));
    }

    const initConfig = () => {
        const config = {
            defaultRes: 720,//默认分辨率
            danmaku: true,
            volume: 80,
            disableType: {
                row: false,
                top: false,
                bottom: false,
                color: false
            }
        }
        localStorage.setItem("player-config", JSON.stringify(config));
        return config;
    }

    return {
        getConfig,
        setConfig
    }
}

