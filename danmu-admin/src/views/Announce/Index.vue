<template>
    <n-card class="card">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">发布公告</n-button>
        </div>
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>ID</th>
                    <th>标题</th>
                    <th>内容</th>
                    <th>URL</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in announceList">
                    <td>{{ item.aid }}</td>
                    <td>{{ item.title }}</td>
                    <td>{{ item.content }}</td>
                    <td>{{ item.url }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>
                        <n-button text @click="deleteAnnounce(item.aid)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
        <div class="pagination">
            <n-pagination v-model:page="page" :item-count="count" :page-size="5" @update-page="pageChange" />
        </div>
    </n-card>
    <n-drawer v-model:show="showAdd" :width="500" placement="right">
        <n-drawer-content title="发布公告">
            <n-form class="info-form">
                <n-form-item label="标题">
                    <n-input v-model:value="announceForm.title" placeholder="请输入名称" maxlength="50" show-count />
                </n-form-item>
                <n-form-item label="内容">
                    <n-input v-model:value="announceForm.content" placeholder="请输入内容" maxlength="200" show-count
                        type="textarea" :autosize="{ minRows: 3, maxRows: 5 }" />
                </n-form-item>
                <n-form-item label="URL">
                    <n-input v-model:value="announceForm.url" placeholder="请输入URL" maxlength="100" show-count />
                </n-form-item>
                <n-button type="primary" @click="submitAnnounce">保存</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script lang="ts">
import {
    NTable, NCard, NTime, NDrawer, NInput, NPagination,
    NForm, NFormItem, NDrawerContent, NButton, useNotification
} from 'naive-ui';
import { getAnnounceListAPI, addAnnounceAPI, deleteAnnounceAPI } from '@/api/announce';
import { defineComponent, onBeforeMount, reactive, ref } from 'vue';
import { announceType } from '@/types/announce';

export default defineComponent({
    setup() {
        const page = ref(1);
        const count = ref(0);
        const showAdd = ref(false);//显示编辑抽屉
        const announceList = ref<Array<announceType>>([]);
        const notification = useNotification();//通知

        const getAnnounceList = () => {
            getAnnounceListAPI(page.value, 5).then((res) => {
                if (res.data.code === 2000) {
                    count.value = res.data.data.count;
                    announceList.value = res.data.data.announces;
                }
            }).catch(() => {
                notification.error({
                    title: '加载轮播图失败',
                    duration: 5000,
                })
            });
        }

        //页码改变
        const pageChange = (target: number) => {
            page.value = target;
            getAnnounceList();
        }


        //新增轮播图
        const announceForm = reactive({
            title: '',
            content: '',
            url: ''
        })

        //上传轮播图
        const submitAnnounce = () => {
            if (!announceForm.title || !announceForm.content) {
                notification.error({
                    title: '请输入标题和内容',
                    duration: 5000
                });
                return;
            }
            addAnnounceAPI(announceForm).then((res) => {
                if (res.data.code === 2000) {
                    getAnnounceList();
                    announceForm.url = "";
                    announceForm.title = "";
                    announceForm.content = "";
                    showAdd.value = false;
                }
            }).catch((err) => {
                notification.error({
                    title: '上传轮播图失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        const deleteAnnounce = (id: number) => {
            deleteAnnounceAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    getAnnounceList();
                }
            }).catch(() => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        onBeforeMount(() => {
            getAnnounceList();
        })

        return {
            page,
            count,
            showAdd,
            announceForm,
            announceList,
            pageChange,
            submitAnnounce,
            deleteAnnounce
        }
    },
    components: {
        NInput,
        NForm,
        NFormItem,
        NTable,
        NCard,
        NTime,
        NDrawer,
        NButton,
        NPagination,
        NDrawerContent,
    }
});
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;

    .info-header {
        height: 30px;
        margin-bottom: 20px;
    }

    .table-head {
        text-align: center;
    }

    .table-body {
        text-align: center;
    }

}


.pagination {
    margin-top: 20px;
}
</style>