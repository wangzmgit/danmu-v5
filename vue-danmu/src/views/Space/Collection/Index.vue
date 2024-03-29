<template>
    <div class="collection-list">
        <p class="create-title">收藏夹</p>
        <n-scrollbar style="max-height: 550px;">
            <div class="card-item" v-for="(item, index) in collections" :key="index">
                <div class="card-left">
                    <img v-if="item.cover" :src="item.cover" alt="收藏夹封面">
                    <div class="no-cover" v-else>
                        <img src="@/assets/logo.png" alt="默认封面">
                    </div>
                </div>
                <div class="card-center">
                    <p class="title" @click="goCollectionDetails(item.id)">{{ item.name }}</p>
                    <span class="desc">简介：{{ item.desc }}</span>
                    <span class="desc">创建于：<n-time :time="new Date(item.created_at!)"></n-time></span>
                </div>
                <div class="card-right">
                    <n-icon class="edit" size="20" @click="beforeEdit(item)">
                        <create-outline />
                    </n-icon>
                    <n-popconfirm negative-text="取消" positive-text="确认" @positive-click="deleteClick(item.id, index)">
                        <template #trigger>
                            <n-icon class="edit" size="20">
                                <trash-outline />
                            </n-icon>
                        </template>
                        是否删除这个收藏夹
                    </n-popconfirm>
                </div>
            </div>
        </n-scrollbar>
    </div>
    <n-drawer v-model:show="active" :width="500" placement="right">
        <n-drawer-content title="编辑收藏夹">
            <cover-upload :cover="collectionInfo.cover" @finish="finishUpload"></cover-upload>
            <n-form class="info-form">
                <n-form-item label="名称">
                    <n-input v-model:value="collectionInfo.name" placeholder="请输入名称" maxlength="20" show-count />
                </n-form-item>
                <n-form-item label="简介">
                    <n-input v-model:value="collectionInfo.desc" placeholder="收藏夹简介~" maxlength="150" show-count
                        type="textarea" :autosize="descSize" />
                </n-form-item>
                <n-form-item label="公开">
                    <n-switch v-model:value="collectionInfo.open" />
                </n-form-item>
                <div class="upload-next-btn">
                    <n-button type="primary" @click="modifyCollection">保存</n-button>
                </div>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script lang="ts">
import { useRouter } from 'vue-router';
import useCollection from '@/hooks/collection';
import { collectionType, modifyCollectionType} from '@/types/collect';
import CoverUpload from "@/components/CoverUpload.vue";
import { CreateOutline, TrashOutline } from '@vicons/ionicons5';
import { ref, reactive, onBeforeMount, defineComponent } from 'vue';
import { modifyCollectionAPI, deleteCollectionAPI } from '@/api/collect';
import {
    NTime, NIcon, NForm, NFormItem, NButton, NInput, NSwitch,
    NScrollbar, NPopconfirm, NDrawer, NDrawerContent, useNotification
} from 'naive-ui';

export default defineComponent({
    setup() {
        //简介输入框大小
        const descSize = {
            minRows: 3,
            maxRows: 3
        }
        const active = ref(false);//显示编辑抽屉
        const collectionInfo = reactive<modifyCollectionType>({
            id: 0,
            cover: "",
            name: "",
            desc: "",
            open: false,
        });
        const notification = useNotification();
        const { collections, getCollectionList } = useCollection();

        //删除收藏夹
        const deleteClick = (id:number, index:number) => {
            deleteCollectionAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    collections.value.splice(index, 1);
                }
            })
        }

        const beforeEdit = (item:collectionType) => {
            collectionInfo.id = item.id;
            collectionInfo.cover = item.cover!;
            collectionInfo.name = item.name!;
            collectionInfo.desc = item.desc!;
            collectionInfo.open = item.open!;
            active.value = true;
        }

        //封面上传完成
        const finishUpload = (cover:string) => {
            collectionInfo.cover = cover;
        }

        //修改收藏夹
        const modifyCollection = () => {
            if (!collectionInfo.name) {
                notification.error({
                    title: '请输入收藏夹名',
                });
                return;
            }
            modifyCollectionAPI(collectionInfo).then((res) => {
                if (res.data.code === 2000) {
                    getCollectionList();
                }
            }).catch((err) => {
                notification.error({
                    title: '修改失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        //前往收藏夹详情
        const router = useRouter();
        const goCollectionDetails = (id:number) => {
            router.push({ name: "CollectionDetails", params: { id: id } });
        }

        onBeforeMount(() => {
            getCollectionList();
        })

        return {
            active,
            descSize,
            collections,
            collectionInfo,
            beforeEdit,
            deleteClick,
            finishUpload,
            modifyCollection,
            getCollectionList,
            goCollectionDetails,
        }
    },
    components: {
        NTime,
        NIcon,
        NForm,
        NInput,
        NSwitch,
        NButton,
        NDrawer,
        NFormItem,
        NScrollbar,
        NPopconfirm,
        NDrawerContent,
        CoverUpload,
        TrashOutline,
        CreateOutline
    }
});
</script>

<style lang="less" scoped>
.collection-list {
    margin: 0 10px;

    .create-title {
        line-height: 30px;
        font-size: 20px;
        user-select: none;
    }

    .card-item {
        width: 100%;
        height: 80px;
        margin-bottom: 12px;
        border-bottom: 1px solid #e6e6e6;
        padding-bottom: 12px;
    }
}

.card-left {
    float: left;
    width: 120px;
    height: 80px;
    margin-right: 10px;

    img {
        width: 100%;
        height: 100%;
        border-radius: 2px;
    }

    .no-cover {
        width: 100%;
        height: 100%;
        background-color: #e6e6e6;
        border-radius: 2px;

        img {
            position: absolute;
            width: 50px;
            height: 50px;
            margin: 15px 35px;
            filter: grayscale(100%);
            opacity: 50%;
        }
    }
}

.card-center {
    float: left;
    width: calc(100% - 230px);

    .title {
        font-size: 14px;
        color: #212121;
        line-height: 18px;
        margin: 0 0 26px;
        cursor: pointer;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;

        &:hover {
            color: #36ad6a;
        }
    }

    .desc {
        font-size: 12px;
        color: #999;
        overflow: hidden;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 1;
        -webkit-box-orient: vertical;
    }
}

.card-right {
    float: left;
    width: 90px;
    height: 100%;
    text-align: center;


    .edit {
        color: #999;
        cursor: pointer;
        line-height: 90px;
        margin-right: 20px;

        &:hover {
            color: #36ad6a;
        }
    }
}
</style>
