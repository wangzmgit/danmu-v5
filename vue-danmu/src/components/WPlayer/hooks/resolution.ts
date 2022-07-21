import { ref } from 'vue';
import { resourceType } from '../types/resource';

export default function useResolution() {
    const maxRes = ref(0);//最大分辨率

    //获取最大分辨率
    const getMaxRes = (resource: resourceType): number => {
        if (resource.res1080) return 1080;
        if (resource.res720) return 720;
        if (resource.res480) return 480;
        if (resource.res360) return 360;
        return 0;
    }

    return {
        maxRes,
        getMaxRes
    }
}

