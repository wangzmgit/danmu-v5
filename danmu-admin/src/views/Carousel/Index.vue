<template>
    <n-card class="card">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">上传轮播图</n-button>
        </div>
        <n-table striped>
            <thead class="table-head">
                <tr>
                    <th>ID</th>
                    <th>图片</th>
                    <th>链接</th>
                    <th>上传时间</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody class="table-body">
                <tr v-for="item in carouselList">
                    <td>{{ item.id }}</td>
                    <td>
                        <img class="cover" :src="item.img" alt="视频封面">
                    </td>
                    <td>{{ item.url }}</td>
                    <td>
                        <n-time :time="new Date(item.created_at)" />
                    </td>
                    <td>
                        <n-button text @click="deleteCarousel(item.id, index)">删除</n-button>
                    </td>
                </tr>
            </tbody>
        </n-table>
    </n-card>
    <n-drawer v-model:show="showAdd" :width="500" placement="right">
        <n-drawer-content title="添加轮播图">
            <carousel-upload :cover="carouselForm.img" @finish="finishUpload"></carousel-upload>
            <n-form class="info-form">
                <n-form-item label="名称">
                    <n-input v-model:value="carouselForm.url" placeholder="请输入名称" maxlength="100" show-count />
                </n-form-item>
                <n-button type="primary" @click="submitCarousel">保存</n-button>
            </n-form>
        </n-drawer-content>
    </n-drawer>
</template>

<script>
import {
    NTable, NCard, NTime, NDrawer, NInput, NForm,
    NFormItem, NDrawerContent, NButton, useNotification
} from 'naive-ui';
import { getCarouselAPI, addCarouselAPI, deleteCarouselAPI } from '@/api/carousel';
import { onBeforeMount, reactive, ref } from 'vue';
import CarouselUpload from '@/components/CarouselUpload.vue';

export default {
    setup() {

        const showAdd = ref(false);//显示编辑抽屉
        const carouselList = ref([]);
        const notification = useNotification();//通知

        const getCarousel = () => {
            getCarouselAPI().then((res) => {
                if (res.data.code === 2000) {
                    carouselList.value = res.data.data.carousels;
                }
            }).catch((err) => {
                notification.error({
                    title: '加载轮播图失败',
                    duration: 5000,
                })
            });
        }

        //新增轮播图
        const carouselForm = reactive({
            img: '',
            url: ''
        })

        //封面上传完成
        const finishUpload = (cover) => {
            carouselForm.img = cover;
        }

        //上传轮播图
        const submitCarousel = () => {
            if (!carouselForm.img) {
                notification.error({
                    title: '请上传先图片',
                    duration: 5000
                });
                return;
            }
            addCarouselAPI(carouselForm).then((res) => {
                if (res.data.code === 2000) {
                    getCarousel();
                    carouselForm.img = "";
                    carouselForm.url = "";
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

        const deleteCarousel = (id, index) => {
            deleteCarouselAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    carouselList.value.splice(index, 1);
                }
            }).catch((err) => {
                notification.error({
                    title: '删除失败',
                    duration: 5000,
                })
            });
        }

        onBeforeMount(() => {
            getCarousel();
        })

        return {
            showAdd,
            carouselForm,
            carouselList,
            finishUpload,
            submitCarousel,
            deleteCarousel
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
        NDrawerContent,
        CarouselUpload
    }
}
</script>

<style lang="less" scoped>
.card {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
}

.info-header {
    height: 30px;
    margin-bottom: 20px;
}

.table-head {
    text-align: center;
}

.table-body {
    text-align: center;

    .cover {
        height: 60px;
        width: 100px;
    }
}
</style>