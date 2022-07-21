<template>
    <div class="announce-box">
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
    </div>
    <div class="pagination">
        <n-pagination v-model:page="page" :item-count="count" :page-size="8" @update-page="pageChange"></n-pagination>
    </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from 'vue';
import { NPagination, NTime } from 'naive-ui';
import { LinkOutline } from '@vicons/ionicons5';
import { getAnnounceAPI } from '@/api/announce';

export default defineComponent({
    setup() {
        const page = ref(1);
        const count = ref(0);
        const announceList = ref([]);

        const linkTo = (url: string) => {
            window.open(url, "_blank");
        }

        const getAnnounce = () => {
            getAnnounceAPI(page.value, 8).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    announceList.value = res.data.data.announces;
                }
            })
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getAnnounce();
        }

        onBeforeMount(() => {
            getAnnounce();
        })

        return {
            page,
            count,
            announceList,
            linkTo,
            pageChange
        }
    },
    components: {
        NTime,
        NPagination,
        LinkOutline
    }
});
</script>

<style lang="less" scoped>
.announce-box {
    min-height: 600px;
    margin: 16px 10px 10px;
}

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

.pagination {
    padding-bottom: 16px;
}
</style>