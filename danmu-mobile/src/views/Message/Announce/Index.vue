<template>
    <div v-title data-title="系统公告">
        <header-bar></header-bar>
        <div v-if="loading" class="loading">
            <n-spin type="loading" />
        </div>
        <div v-else class="announce-box">
            <div class="announce-item" v-for="(item, index) in announceList" :key="index">
                <div class="title">
                    <p class="item-title">{{ item.title }}</p>
                    <n-time class="item-time" :time="new Date(item.created_at)"></n-time>
                </div>
                <span>{{ item.content }}</span>
                <span class="link" v-if="item.url">
                    <span @click="linkTo(item.url)">网页链接</span>
                </span>
            </div>
            <div class="more">
                <n-button v-if="!noMore" @click="getAnnounceList(++page)">加载更多</n-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { NSpin, NTime,NButton } from 'naive-ui';
import { getAnnounceAPI } from "@/api/announce";
import { announceType } from '@/types/announce';
import HeaderBar from "@/components/HeaderBar.vue";
import { defineComponent, onBeforeMount, ref } from "vue";


export default defineComponent({
    setup() {
        const page = ref(1);
        const pageSize = 10;
        const noMore = ref(false);
        const loading = ref(false);
        const announceList = ref<Array<announceType>>([]);

        const getAnnounceList = (page: number) => {
            getAnnounceAPI(page, pageSize).then((res) => {
                if (res.data.code === 2000) {
                    const resAnnounce = res.data.data.announces;
                    if (resAnnounce.length < pageSize) {
                        noMore.value = true;
                    }

                    announceList.value.push(...resAnnounce);
                    loading.value = false;
                }
            });
        }

        const linkTo = (url: string) => {
            window.open(url, "_blank");
        }

        onBeforeMount(() => {
            getAnnounceList(page.value);
        })

        return {
            page,
            noMore,
            loading,
            announceList,
            linkTo,
            getAnnounceList
        }
    },
    components: {
        NSpin,
        NTime,
        NButton,
        HeaderBar,
    },
});
</script>

<style lang="less" scoped>
.loading {
    font-size: 30px;
    text-align: center;
}

.announce-box {
    min-height: 600px;

    .announce-item {
        padding: 8px 6px;
        border-bottom: 1px solid #e1e2e3;

        .title {
            display: flex;
            align-items: center;
            justify-content: space-between;
            margin: 5px 0;

            .item-title {
                font-size: 16px;
                font-weight: 500;
                width: calc(100% - 160px);
                margin: 0;
            }

            .item-time {
                color: #999;
            }
        }

        .link {
            margin-left: 10px;
            color: #1890ff;
            cursor: pointer;
        }
    }

    .more {
        text-align: center;
        margin-top: 10px;
    }
}
</style>