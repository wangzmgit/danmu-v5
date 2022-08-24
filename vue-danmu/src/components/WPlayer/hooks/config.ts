export default function useConfig() {
    const configItems = ["defaultRes", "danmaku", "volume", "disableType"];
    const getConfig = () => {
        const config = localStorage.getItem("player-config");
        if (!config) {
            return initConfig();
        }

        const objConfig = JSON.parse(config);
        // 检查配置项是否存在
        const configKeys = Object.keys(objConfig);
        for (const item of configItems) {
            if (!configKeys.includes(item)) {
                initConfig();
            }
        }
        
        return objConfig;
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

