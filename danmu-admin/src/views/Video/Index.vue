<template>
    <n-card class="card">
        <div class="info-header">
            <search-box :keyword="keyword" @search="searchVideo"></search-box>
        </div>
        <n-table class="table" striped>
            <thead class="table-head">
                <tr>
                    <th>VID</th>
                    <th>封面</th>
                    <th>标题</th>
                    <th>简介</th>
                    <th>分区</th>
                    <th>作者UID</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in videoList">
                    <td>{{ item.vid }}</td>
                    <td>
                        <img class="cover" :src="item.cover" alt="视频封面">
                    </td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.desc }}</td>
                    <td>{{ item.partition }}</td>
                    <td>{{ item.uid }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td class="btn-list">
                        <n-button text @click="searchAuthor(item.uid)">查找作者</n-button>
                        <n-button text @click="deleteVideo(item.vid)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="6" @update-page="pageChange" />
        </div>
    </n-card>
</template>

<script>
import { onBeforeMount, ref } from 'vue';
import { useRouter } from 'vue-router';
import SearchBox from '@/components/SearchBox.vue';
import { getVideoListAPI, searchVideoAPI, deleteVideoAPI } from '@/api/video';
import { NTable, NButton, NCard, NTime, NPagination, useNotification } from 'naive-ui';

export default {
    setup() {
        const page = ref(1);
        const count = ref(0);
        const videoList = ref([]);
        const notification = useNotification();//通知

        const getVideoList = () => {
            getVideoListAPI(page.value, 6).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    videoList.value = res.data.data.videos;
                }
            }).catch((err) => {
                notification.error({
                    title: '加载视频失败',
                    duration: 5000,
                })
            });
        }

        //页码改变
        const pageChange = (target) => {
            page.value = target;
            getVideoList();
        }

        //查找作者
        const router = useRouter();
        const searchAuthor = (uid) => {
            router.push({ name: 'User', query: { uid: uid } });
        }

        //搜索
        const keyword = ref('');
        const searchVideo = (keyword) => {
            page.value = 1;
            count.value = 0;
            if (!keyword) {
                getVideoList();
                return;
            }
            searchVideoAPI(keyword).then((res) => {
                if (res.data.code === 2000) {
                    videoList.value = res.data.data.videos;
                }
            })
        }

        //删除
        const deleteVideo = (id) => {
            deleteVideoAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    getVideoList();
                }
            }).catch((err) => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        onBeforeMount(() => {
            getVideoList();
        })

        return {
            page,
            count,
            keyword,
            videoList,
            pageChange,
            searchAuthor,
            searchVideo,
            deleteVideo
        }
    },
    components: {
        NTable,
        NCard,
        NTime,
        NButton,
        SearchBox,
        NPagination,
    }
}
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        float: right;
        height: 30px;
        width: 300px;
        margin-bottom: 20px;
    }

    .table {

        .table-head {
            text-align: center;
        }

        .table-body {
            text-align: center;

            .cover {
                height: 60px;
                width: 100px;
            }

            .btn-list {
                button {
                    margin: 0 6px;
                }
            }
        }
    }
}

.pagination {
    margin-top: 20px;
}
</style>