<template>
    <div class="player-container">
        <w-player class="player" :key="playerKey" :options="options"></w-player>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, defineProps, onBeforeMount, watch, withDefaults } from 'vue';
import useHls from '@/hooks/hls';
import { WPlayer } from 'vue-wplayer';
import 'vue-wplayer/dist/style.css';
import { optionsType, qualityType } from '@/types/player';
import { danmakuType } from "@/types/danmaku";
import { getDanmakuAPI, sendDanmakuAPI } from '@/api/danmaku';

const props = withDefaults(defineProps<{
    vid: number,
    resources: Array<any>,
    part: number
}>(), {
    part: 1
})

const { loadHls } = useHls();
const playerKey = ref(0);

const options = reactive<optionsType>({
    resource: [],
    type: "hls",
    mobile: true,
    customType: (player, src) => {
        loadHls(src, player);
    },
    theme: "#18a058",
    danmaku: {
        open: true,
        data: [],
        send: (danmaku) => {
            sendDanmaku(danmaku);
        }
    }
});

const loadPart = async (part: number) => {
    loadResource(part);
    await getDanmaku(part);

    playerKey.value = Date.now();
}

const loadResource = (part: number) => {
    let tmpResource: qualityType | string = {};
    if (props.resources[part - 1]["original"]) {
        tmpResource = props.resources[part - 1]["original"];
    } else {
        for (let key of Object.keys(props.resources[part - 1])) {
            const reg = /res(\d+)/g;
            const matchArr = reg.exec(key);
            if (matchArr && matchArr[1]) {
                console.log('matchArr', matchArr)
                tmpResource[parseInt(matchArr[1])] = {
                    name: `${matchArr[1]}P`,
                    url: props.resources[part - 1][key],
                }
            }
        }
    }
    options.resource = tmpResource;
}


const sendDanmaku = (danmakuForm: danmakuType) => {
    danmakuForm.vid = props.vid;
    danmakuForm.part = props.part;
    sendDanmakuAPI(danmakuForm);
}

const getDanmaku = async (part: number) => {
    if (options.danmaku?.data) options.danmaku.data = [];
    const res = await getDanmakuAPI(props.vid, part);
    if (res.data.code === 2000) {
        const list = res.data.data.danmaku;
        if (list) {
            options.danmaku!.data = list;
        }
    }
}

watch(() => props.part, (val) => {
    loadPart(val);
});

onBeforeMount(() => {
    loadPart(props.part);
});
</script>

<style lang="less" scoped>
.player-container {
    height: 0;
    width: 100%;
    padding-bottom: 56.25%;
    position: relative;
    margin-bottom: 40px;

    .player {
        width: 100%;
        height: 100%;
        position: absolute;
        background-color: black;
    }
}
</style>