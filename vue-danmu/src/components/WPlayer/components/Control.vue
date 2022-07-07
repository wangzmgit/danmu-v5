<template>
    <!--时间轴-->
    <div class="time-line">
        <div id="slider" ref="timeSliderRef" class="player-bar" @click="clickSlider">
            <div id="loaded" class="player-loaded" :style="`width: ${schedule(videoInfo.loadedTime)}%`"></div>
            <div id="played" class="player-played" :style="`width: ${schedule(videoInfo.currentTime)}%`">
                <div ref="blockRef" class="player-block"></div>
            </div>
        </div>
    </div>
    <div class="control-box">
        <!--控制栏左-->
        <div class="control-left">
            <w-button type="text" @click="playOrPause">
                <svg-icon class="control-icon" :icon="videoInfo.state"></svg-icon>
            </w-button>
            <span class="time-text">{{ toTimeText(videoInfo.currentTime) }} / {{ toTimeText(videoInfo.duration)
            }}</span>
        </div>
        <!--控制栏中(全屏状态下的弹幕发送区)-->
        <div class="control-center"></div>
        <!--控制栏右-->
        <div class="control-right">
            <w-button class="res" type="text" @click="showMenu('res')">{{ videoInfo.resText }}</w-button>
            <div class="res-menu" v-show="menus.res">
                <div v-if="videoInfo.maxRes != 0">
                    <w-button v-if="videoInfo.maxRes >= 1080" type="text" @click="setRes(1080)">1080P</w-button>
                    <w-button v-if="videoInfo.maxRes >= 720" type="text" @click="setRes(720)">720P</w-button>
                    <w-button v-if="videoInfo.maxRes >= 480" type="text" @click="setRes(480)">480P</w-button>
                    <w-button type="text" @click="setRes(360)">360P</w-button>
                </div>
            </div>
            <w-button class="speed" type="text" @click="showMenu('speed')">{{ videoInfo.speedText }}</w-button>
            <div class="speed-menu" v-show="menus.speed">
                <w-button type="text" @click="setSpeed(0.5)">0.5x</w-button>
                <w-button type="text" @click="setSpeed(0.75)">0.75x</w-button>
                <w-button type="text" @click="setSpeed(1)">1.0x</w-button>
                <w-button type="text" @click="setSpeed(1.25)">1.25x</w-button>
                <w-button type="text" @click="setSpeed(1.5)">1.5x</w-button>
                <w-button type="text" @click="setSpeed(2)">2.0x</w-button>
            </div>
            <!-- 音量 -->
            <w-button class="right-icon" type="text" @click="showMenu('volume')">
                <svg-icon class="control-icon" icon="volume"></svg-icon>
            </w-button>
            <div class="volume" v-show="menus.volume">
                <slider class="slider" :value="videoInfo.volume" vertical @changeValue="setVolume"></slider>
            </div>
            <w-button class="right-icon" type="text" @click="fullScreen">
                <svg-icon class="control-icon" icon="fullScreen"></svg-icon>
            </w-button>
        </div>
    </div>
</template>

<script>
import Slider from './Slider.vue';
import useConfig from '../hooks/config';
import { ref, onMounted, reactive } from 'vue';
import SvgIcon from "../components/SvgIcon.vue";
import WButton from "../components/WButton.vue";


export default {
    emits: ['playChange', 'resChange', 'full', 'progressChange', 'volumeChange', 'speedChange'],
    props: {
        left: {
            type: Number,
        }
    },
    setup(props, ctx) {
        const blockRef = ref(null);
        const timeSliderRef = ref(null);
        const { getConfig, setConfig } = useConfig();
        const playerConfig = getConfig();

        //视频状态
        const videoState = {
            PLAY: "play",
            PAUSE: "pause",
            REPLAY: "replay"
        }

        const menus = reactive({
            speed: false,
            volume: false,
            res: false,
        })

        const videoInfo = reactive({
            duration: 0,//视频时长
            state: videoState.PLAY,//视频状态
            currentTime: 0,//当前时间
            loadedTime: 0,//当前加载时间
            maxRes: 0,//最大分辨率
            resText: "原始",//分辨率文本
            currentRes: 0,//当前分辨率
            speedText: "倍速", //视频倍速
            volume: 80, //音量
        })

        //初始化
        const initControl = (val, maxRes) => {
            videoInfo.duration = val;
            videoInfo.maxRes = maxRes;
        }

        //时间改变
        const timeUpdate = (current, loaded) => {
            videoInfo.currentTime = current;
            videoInfo.loadedTime = loaded;
        }

        //视频结束
        const videoEnd = () => {
            videoInfo.state = videoState.REPLAY;
        }

        //计算播放和加载进度
        const schedule = (time) => {
            return Math.round((time / videoInfo.duration) * 10000) / 100;
        }

        //播放和暂停
        const playOrPause = () => {
            let play = true;
            switch (videoInfo.state) {
                case videoState.PLAY:
                    videoInfo.state = videoState.PAUSE;
                    break;
                case videoState.PAUSE:
                    play = false;
                    videoInfo.state = videoState.PLAY;
                    break;
                case videoState.REPLAY:
                    videoInfo.state = videoState.PAUSE;
                    break;
            }
            ctx.emit('playChange', play);
        }

        //点击滑动条
        const clickSlider = (e) => {
            let id = e.target.id;
            let slider = timeSliderRef.value;
            if (id == "slider" || id == "loaded" || id == "played") {
                if (videoInfo.state === videoState.REPLAY) {
                    videoInfo.state = videoState.PLAY;
                }
                videoInfo.currentTime = (e.offsetX / slider.clientWidth) * videoInfo.duration;
                ctx.emit('progressChange', videoInfo.currentTime);
            }
        }

        // 滑动滑动条
        const slideTimeLine = () => {
            blockRef.value.onmousedown = function () {
                let width = timeSliderRef.value.offsetWidth;
                document.onmousemove = function (e) {
                    //计算新的百分比
                    let percentage = Math.round(((e.clientX - props.left) / width) * 100) / 100;
                    percentage = Math.max(0, percentage);
                    percentage = Math.min(percentage, 1);
                    videoInfo.currentTime = percentage * videoInfo.duration;
                    //更新视频进度
                    ctx.emit('progressChange', videoInfo.currentTime);
                };
                document.onmouseup = function () {
                    document.onmousemove = document.onmouseup = null;
                };
            };
        }

        //快进或快退 方向dir(true前进，false后退) 时间n
        const fastForward = (dir, n) => {
            let dur = videoInfo.duration;
            let current = videoInfo.currentTime;
            if (dir) {
                ctx.emit('progressChange', current + n > dur ? dur : current + n);
            } else {
                ctx.emit('progressChange', current - n < 0 ? 0 : current - n);
            }
        }

        //转为时间文本
        const toTimeText = (time) => {
            let m = parseInt(time / 60);
            let s = parseInt(time % 60);
            m = m < 10 ? "0" + m : m;
            s = s < 10 ? "0" + s : s;
            return m + ":" + s;
        }

        //打开或关闭菜单
        const showMenu = (name) => {
            //关闭除了name以外所有的菜单
            for (let key in menus) {
                if (key == name) {
                    menus[key] = !menus[key];
                    continue;
                }
                menus[key] = false;
            }
        }

        //设置分辨率
        const setRes = (res) => {
            if (res === 0) return;
            if (videoInfo.currentRes === res) return;
            //设置分辨率
            videoInfo.currentRes = res;
            videoInfo.resText = res + 'P';
            let play = videoInfo.state === videoState.PAUSE ? false : true;
            // 调用父组件设置视频资源和播放状态
            ctx.emit("resChange", res, currentTime, play);
        }

        //设置倍速
        const setSpeed = (speed) => {
            if (speed != 1) {
                videoInfo.speedText = speed + "x";
            } else {
                videoInfo.speedText = "倍速";
            }
            menus.speed = false;
            ctx.emit('speedChange', speed);
        }

        //改变音量
        const setVolume = (volume) => {
            setConfig('volume', volume);
            ctx.emit('volumeChange', volume);
        }

        //设置全屏
        const fullScreen = () => {
            ctx.emit('full')
        }

        onMounted(() => {
            slideTimeLine();
            videoInfo.volume = playerConfig.volume;
        })

        return {
            schedule,
            menus,
            videoInfo,
            blockRef,
            timeSliderRef,
            initControl,
            timeUpdate,
            videoEnd,
            playOrPause,
            toTimeText,
            showMenu,
            setRes,
            setSpeed,
            setVolume,
            fullScreen,
            clickSlider,
            fastForward
        }
    },
    components: {
        Slider,
        SvgIcon,
        WButton,
    }
}
</script>

<style lang="less" scoped>
.time-line {
    margin-left: 6px;
    width: calc(100% - 12px);

    .player-bar {
        position: relative;
        width: 100%;
        background: rgba(107, 107, 107, 0.6);
        height: 4px;
        cursor: pointer;
    }

    .player-loaded {
        position: absolute;
        background: #fff;
        height: 4px;
    }

    .player-played {
        position: absolute;
        background: #18a058;
        height: 4px;

        .player-block {
            position: absolute;
            top: -6px;
            right: -6px;
            width: 12px;
            height: 12px;
            border-radius: 50%;
            border: 2px solid #18a058;
            background: #fff;
            transition: 0.2s all;
            user-select: none;
        }
    }
}

.control-box {
    display: flex;

    //图标
    .control-icon {
        width: 30px;
        height: 30px;
        margin-top: 8px;
    }

    .control-left {
        width: 200px;
        display: flex;

        .time-text {
            color: #fff;
            font-size: 14px;
            margin-top: 15px;
        }
    }

    .control-center {
        width: calc(100% - 370px);
    }

    .control-right {
        display: flex;
        justify-content: space-between;
        width: 230px;
    }
}

// 分辨率菜单
.res {
    padding: 0 10px;
    margin-top: 8px;
}

.res-menu {
    width: 60px;
    bottom: 56px;
    right: 156px;
    position: absolute;
    text-align: center;
    border-color: transparent;
    border-radius: 5px;
    background: rgba(0, 0, 0, 0.5);
}

.speed {
    padding: 0 10px;
    margin-top: 8px;
}

.speed-menu {
    width: 60px;
    bottom: 56px;
    right: 116px;
    position: absolute;
    text-align: center;
    border-color: transparent;
    border-radius: 5px;
    background: rgba(0, 0, 0, 0.5);
}

.volume {
    width: 32px;
    height: 140px;
    right: 70px;
    bottom: 56px;
    border-radius: 5px;
    position: absolute;
    background: rgba(0, 0, 0, 0.7);

    .slider {
        top: 12px;
        left: 8px;
        height: 116px;
    }
}

.right-icon {
    padding: 0 10px;
}
</style>