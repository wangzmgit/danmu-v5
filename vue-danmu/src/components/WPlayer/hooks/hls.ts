import Hls from "hls.js";
import { onUnmounted, ref } from 'vue'
export default function useHls() {
    //hls
    const hls = ref<Hls | null>(null);

    //加载hls
    const loadHls = (src: string, player: HTMLVideoElement) => {
        hls.value = new Hls()
        hls.value.loadSource(src);
        hls.value.attachMedia(player);
        hls.value.on(Hls.Events.ERROR, () => {
            console.error("资源加载失败");
        });
    }

    onUnmounted(() => {
        if (hls.value) {
            hls.value.stopLoad();
            hls.value.destroy();
            hls.value = null;
        }
    })

    return {
        hls,
        loadHls,
    }
}