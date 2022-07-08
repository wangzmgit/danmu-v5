<template>
    <div class="about">
        <n-card v-for="item in serverInfo" class="card-item" :title="item.title">
            <span>版本：{{ item.version?item.version:'未知' }}</span>
            <span>当前状态：{{ item.status ? '正常' : '异常' }}</span>
        </n-card>
        <n-card class="card-item" title="前端">
            <span>版本：{{ feInfo.version }}</span>
            <span>作者：{{ feInfo.author }}</span>
        </n-card>
    </div>
</template>

<script>
import { NCard } from "naive-ui";
import { getServerDataAPI } from '@/api/config';
import { onBeforeMount, ref } from "vue";

export default {
    setup() {
        const feInfo = ref({});
        const serverInfo = ref([]);
        const getServerData = () => {
            getServerDataAPI().then((res) => {
                if (res.data.code === 2000) {
                    const resData = res.data.data.info;
                    const infoMeaning = new Map([
                        ['version', '后端'],
                        ['ffmpeg', 'FFmpeg'],
                        ['redis', 'Redis'],
                        ['mysql', 'Mysql'],
                    ]);

                    feInfo.value = resData.web;
                    for (let key in resData) {
                        if (typeof (resData[key]) !== 'object') {
                            serverInfo.value.push({
                                title: infoMeaning.get(key),
                                version: resData[key],
                                status: resData[key] ? true : false
                            })
                        }
                    }
                }
            })
        }

        onBeforeMount(() => {
            getServerData();
        })

        return {
            feInfo,
            serverInfo
        }
    },
    components: {
        NCard
    }
}
</script>

<style lang="less" scoped>
.about {
    width: calc(100% - 40px);
    margin: 10px;
    display: grid;
    row-gap: 10px;
    column-gap: 10px;
    grid-template-columns: repeat(3, 33.333%);

    .card-item {
        span {
            display: block;
        }
    }
}
</style>