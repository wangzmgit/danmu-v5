import { ref } from 'vue';

export default function useResolution() {
    const maxRes = ref(0);//最大分辨率

    //获取最大分辨率
    const getMaxRes = (resource) => {
        if (resource.original) return 0;
        if (resource.res1080) return 1080;
        if (resource.res720) return 720;
        if (resource.res480) return 480;
        if (resource.res360) return 360;
    }

    return {
        maxRes,
        getMaxRes
    }
}

