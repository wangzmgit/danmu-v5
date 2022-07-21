import { ref } from 'vue';


export default function useFullScreen() {
    const playerRef = ref<any>(null);

    // 判断是否为全屏
    const isFullScreen = () => {
        return document.fullscreenElement ? true : false;
    }
    //全屏
    const fullScreen = () => {
        if (!isFullScreen()) {
            const doc = playerRef.value;
            if (doc?.requestFullscreen) {
                doc.requestFullscreen();
            } else if (doc?.mozRequestFullScreen) {
                doc?.mozRequestFullScreen();
            } else if (doc.webkitRequestFullScreen) {
                doc?.webkitRequestFullScreen();
            }
        } else {
            const exit: any = document;
            if (exit.exitFullscreen) {
                exit.exitFullscreen();
            } else if (exit.mozCancelFullScreen) {
                exit.mozCancelFullScreen();
            } else if (exit.webkitCancelFullScreen) {
                exit.webkitCancelFullScreen();
            }
        }
    }

    return {
        playerRef,
        fullScreen
    }
}

