<template>
    <div ref="playerRef" class="player-box uncheck" @click="setPlayVideo" @contextmenu.prevent="openMenu"
        @mouseleave="mouseLeave" @mousemove="mouseMove">
        <video id="player" class="player" ref="videoRef" @canplay="initVideo" @timeupdate="timeUpdate"
            @ended="videoEnd" />
        <!-- 控制栏 -->
        <div v-show="showControl" class="control">
            <control ref="controlRef" :left="offestLeft" @playChange="setPlayState" @resChange="setRes"
                @full="fullScreen" @progressChange="setProgress" @volumeChange="setVolume" @speedChange="setSpeed" />
        </div>
        <!-- 弹幕容器 -->
        <danmaku id="danmaku" ref="danmakuRef" v-if="showDanmaku" :list="danmakuList"></danmaku>
        <!-- 右键菜单 -->
        <context-menu ref="menuRef" @mirror="setMirror"></context-menu>
        <!-- 播放器消息 -->
        <player-msg v-show="showMsg" :msg="msg"></player-msg>
    </div>
    <!-- 发送弹幕 -->
    <send-danmaku :show="playerConfig.danmaku" :amount="amount" @changeShow="changeShowDanmaku" @showMsg="changeMsg"
        @send="sendDanmaku" @setOpacity="setOpacity">
    </send-danmaku>
</template>

<script>
import { onMounted, onBeforeMount, ref, onBeforeUnmount } from 'vue'
import useHls from './hooks/hls';
import useMsg from './hooks/msg';
import useConfig from './hooks/config';
import useDanmakuAPI from './hooks/danmaku';
import useResolution from './hooks/resolution';
import useFullScreen from './hooks/full-screen';
import Control from './components/Control';
import Danmaku from './components/Danmaku.vue';
import SendDanmaku from './components/SendDanmaku.vue';
import PlayerMsg from './components/PlayerMsg.vue';
import ContextMenu from './components/ContextMenu.vue';
export default {
    props: {
        vid: {
            type: Number,
            default: null,
        },
        part: {
            type: Number,
            default: 1,
        },
        resource: {
            type: Object,
            default: {},
        },
    },
    setup(props) {
        //弹幕接口
        const { sendDanmakuAPI, getDanmakuAPI } = useDanmakuAPI();
        //设置全屏
        const { playerRef, fullScreen } = useFullScreen();
        //hls
        const { loadHls } = useHls();
        //消息相关
        const { msg, showMsg, changeMsg } = useMsg();
        //分辨率
        const { maxRes, getMaxRes } = useResolution();
        // 播放器配置
        const { getConfig, setConfig } = useConfig();
        const playerConfig = getConfig();

        const videoRef = ref(null);
        const controlRef = ref(null);
        const offestLeft = ref(0);

        const initVideo = () => {
            let duration = videoRef.value.duration;
            maxRes.value = getMaxRes(props.resource.data);
            controlRef.value.initControl(duration, maxRes.value, playerRef);
        }

        //视频结束
        const videoEnd = () => {
            controlRef.value.videoEnd();
            if (danmakuRef.value) {
                danmakuRef.value.clearDanmaku();
            }
        }

        //更新进度时间
        const timeUpdate = () => {
            let loaded = 0;//加载信息
            const videoDOM = videoRef.value;
            if (videoDOM) {
                const currentTime = videoDOM.currentTime;
                if (videoDOM.buffered.length) {
                    loaded = videoDOM.buffered.end(videoDOM.buffered.length - 1)
                }
                controlRef.value.timeUpdate(currentTime, loaded);
                //更新弹幕
                if (danmakuRef.value) {
                    danmakuRef.value.timeUpdate(currentTime);
                }
            }
        }

        //设置视频状态
        const setPlayState = (play) => {
            if (play) {
                videoRef.value.play();
            } else {
                videoRef.value.pause();
            }
            if (danmakuRef.value) {
                danmakuRef.value.startOrPause(play);
            }
        }

        //设置分辨率
        const setRes = (res, currentTime, play) => {
            const resText = !res ? 'original' : 'res' + res;
            if (props.resource.type === "hls") {
                loadHls(props.resource.data[resText], videoRef.value);
            } else {
                videoRef.value.src = props.resource.data[resText];
            }
            //设置播放时间和状态
            videoRef.value.currentTime = currentTime;
            setPlayState(play);
        }

        //设置音量
        const setVolume = (volume) => {
            videoRef.value.volume = volume / 100;
        }

        //设置当前进度
        const setProgress = (currentTime) => {
            videoRef.value.currentTime = currentTime;
            if (danmakuRef.value) {
                danmakuRef.value.clearDanmaku();
            }
        }

        //设置倍速
        const setSpeed = (speed) => {
            videoRef.value.playbackRate = speed;
        }

        //弹幕
        const amount = ref(0);
        const danmakuList = ref([]);
        const danmakuRef = ref(null);
        const showDanmaku = ref(playerConfig.danmaku);
        const sendDanmaku = (danmakuForm) => {
            const tmpForm = {
                vid: props.vid,
                part: props.part,
                time: Math.round(videoRef.value.currentTime),
                text: danmakuForm.text,
                type: danmakuForm.type,
                color: `#${danmakuForm.color}`,
            }
            sendDanmakuAPI(tmpForm);
            //绘制弹幕
            danmakuRef.value.drawDanmaku(tmpForm, true);
        }

        //设置不透明度
        const setOpacity = (opacity) => {
            danmakuRef.value.setOpacity(opacity);
        }

        //设置弹幕显示
        const changeShowDanmaku = (val) => {
            if (!val) {
                danmakuRef.value.clearDanmaku();
            }
            showDanmaku.value = val;
            setConfig('danmaku', val);
        }

        //右键菜单
        const menuRef = ref(null);
        const setPlayVideo = (e) => {
            if (menuRef.value.menuIsShow()) {
                menuRef.value.closeMenu();
            } else {
                const target = e.target.id;
                if (target === 'danmaku' || target === 'player') {
                    controlRef.value.playOrPause();
                }
            }
        }

        //开启菜单
        const openMenu = (e) => {
            menuRef.value.openMenu(e, videoRef.value);
        }

        //设置镜像
        const setMirror = () => {
            videoRef.value.classList.toggle("player-mirror");
        }

        //快捷键
        const keyArray = ref([]);
        const keydown = ref("");
        const handleKeyDown = (e) => {
            if (keyArray.value.length > 0) {
                // a-z的按键 长按去重
                if (keyArray.value.indexOf(e.key.toLowerCase()) >= 0) {
                    return;
                }
            }
            keyArray.value.push(e.key.toLowerCase());
            keydown.value = keyArray.value.join("+");
            // 监听按键捕获
            switch (keydown.value) {
                case " ":
                    e.returnValue = false;
                    controlRef.value.playOrPause();
                    break;
                case "arrowright":
                    e.returnValue = false;
                    controlRef.value.fastForward(true, 10);
                    changeMsg("快进10秒");
                    break;
                case "arrowleft":
                    e.returnValue = false;
                    controlRef.value.fastForward(false, 10);
                    changeMsg("快退10秒");
                    break;
            }
        }

        const handleKeyUp = (e) => {
            keyArray.value.splice(keyArray.value.indexOf(e.key.toLowerCase()), 1);
            keydown.value = keyArray.value.join("+");
            e.preventDefault();
        }

        //显示/隐藏控制栏
        const showControl = ref(true);
        const mouseLeave = () => {
            if (!videoRef.value.paused) {
                controlRef.value.showMenu('');
                showControl.value = false;
            }
        }
        const mouseMove = () => {
            if (!showControl.value) {
                showControl.value = true;
                if (!videoRef.value.paused) {
                    setTimeout(() => {
                        controlRef.value.showMenu('');
                        showControl.value = false;
                    }, 3000);
                }
            }
        }

        onBeforeMount(() => {
            getDanmakuAPI(props.vid, props.part).then((res) => {
                if (res.data.code === 2000) {
                    const list = res.data.data.danmaku;
                    if (list) {
                        danmakuList.value = list;
                        amount.value = list.length;
                    }
                }
            })
        })

        onMounted(() => {
            //获取播放器到屏幕左侧距离，并监听窗口大小
            offestLeft.value = playerRef.value.offsetLeft;
            window.onresize = () => {
                offestLeft.value = playerRef.value.offsetLeft;
            }
            //监听快捷键
            document.addEventListener("keydown", handleKeyDown);
            document.addEventListener("keyup", handleKeyUp);
            //初始化视频
            const res = Math.min(playerConfig.defaultRes, maxRes.value);
            setRes(res, 0, playerConfig.autoPlay);
        })

        onBeforeUnmount(() => {
            document.removeEventListener("keydown", handleKeyDown);
            document.removeEventListener("keyup", handleKeyUp);
        })

        return {
            msg,
            showMsg,
            videoRef,
            playerRef,
            controlRef,
            offestLeft,
            playerConfig,
            changeMsg,
            setRes,
            loadHls,
            initVideo,
            timeUpdate,
            videoEnd,
            setSpeed,
            setProgress,
            setVolume,
            setPlayState,
            fullScreen,
            //弹幕相关
            amount,
            danmakuList,
            danmakuRef,
            showDanmaku,
            setOpacity,
            sendDanmaku,
            changeShowDanmaku,
            //右键菜单
            menuRef,
            openMenu,
            setPlayVideo,
            setMirror,
            //控制栏
            showControl,
            mouseLeave,
            mouseMove
        }
    },
    components: {
        Control,
        Danmaku,
        PlayerMsg,
        SendDanmaku,
        ContextMenu
    }
}
</script>

<style lang="less" scoped>
.uncheck {
    /**禁止文字选中 */
    -moz-user-select: -moz-none;
    -moz-user-select: none;
    -o-user-select: none;
    -khtml-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    user-select: none;
}

/**镜像翻转 */
.player-mirror {
    will-change: transform;
    transform: rotateY(180deg);
    -webkit-transform: rotateY(180deg);
    /* Safari and Chrome */
    -moz-transform: rotateY(180deg);
    /* Firefox */
}

.player-box {
    height: 0;
    width: 100%;
    position: relative;
    padding-bottom: 56.25%;

    .player {
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: black;
    }

    .control {
        z-index: 20;
        position: absolute;
        width: 100%;
        height: 50px;
        background: linear-gradient(rgba(0, 0, 0, 0), #000);
        bottom: 0;
        transition: opacity 1s;
        -moz-transition: opacity 1s;
        -webkit-transition: opacity 1s;
        -o-transition: opacity 1s;
    }
}
</style>