<template>
    <n-card class="partition-page">
        <div class="info-header">
            <n-button type="primary" @click="showAdd = true">新建分区</n-button>
        </div>
        <div class="partitions">
            <div class="partition-item" v-for="(item, index) in allPartition" :key="index">
                <div class="item-left">
                    <p>{{ item.content }}</p>
                    <n-button text @click="deletePartition(item.id)">删除</n-button>
                </div>
                <div class="item-right">
                    <div v-for="(sub, i) in item.subpartition" :key="i">
                        <p>{{ sub.content }}</p>
                        <n-button text @click="deletePartition(sub.id)">删除</n-button>
                    </div>
                </div>
            </div>
        </div>
        <!--新增分区-->
        <n-drawer v-model:show="showAdd" :width="500" placement="right">
            <n-drawer-content title="新建分区">
                <n-form class="info-form">
                    <n-form-item label="分区名">
                        <n-input v-model:value="addForm.content" placeholder="请输入名称" maxlength="20" show-count />
                    </n-form-item>
                    <n-form-item label="所属分区">
                        <n-select v-model:value="addForm.parentId" remote :options="partitions" />
                    </n-form-item>
                    <n-button type="primary" @click="addPartition">保存</n-button>
                </n-form>
            </n-drawer-content>
        </n-drawer>
    </n-card>
</template>

<script lang="ts">
import { ref, reactive, onBeforeMount, defineComponent } from 'vue';
import {
    NButton, NDrawer, NDrawerContent, NInput, NCard,
    NSelect, NForm, NFormItem, useNotification
} from "naive-ui";
import { getAllPartitionAPI, addPartitionAPI, deletePartitionAPI } from "@/api/partition";
import { partitionType } from '@/types/partition';

export default defineComponent({
    setup() {
        const notification = useNotification();

        //获取所有分区
        const allPartition = ref<Array<partitionType>>([]);
        const partitions = ref<Array<{
            label: string,
            value: number
        }>>([]);//一级分区
        const getAllPartitionList = () => {
            getAllPartitionAPI().then((res) => {
                if (res.data.code === 2000) {
                    allPartition.value = res.data.data.partitions;
                    partitions.value = allPartition.value.map((item) => {
                        return {
                            label: item.content,
                            value: item.id
                        }
                    });
                    partitions.value.unshift({
                        label: '一级分区',
                        value: 0
                    });
                }
            }).catch((err) => {
                console.error(err);
                notification.error({
                    title: '分区加载失败',
                    duration: 5000
                })
            });
        }

        //删除分区
        const deletePartition = (id: number) => {
            deletePartitionAPI(id).then((res) => {
                if (res.data.code === 2000) {
                    getAllPartitionList();
                }
            }).catch((err) => {
                notification.error({
                    title: '删除失败',
                    duration: 5000
                })
            });
        }

        //新建分区
        const addForm = reactive({
            content: "",
            parentId: 0,
        });
        const showAdd = ref(false);
        const addPartition = () => {
            if (!addForm.content) {
                notification.error({
                    title: '请输入分区名称',
                    duration: 5000
                });
                return;
            }
            addPartitionAPI(addForm).then((res) => {
                if (res.data.code === 2000) {
                    showAdd.value = false;
                    getAllPartitionList();
                }
            }).catch(() => {
                notification.error({
                    title: '创建分区失败',
                    duration: 5000
                });
            });
        }

        onBeforeMount(() => {
            getAllPartitionList();
        })

        return {
            showAdd,
            addForm,
            partitions,
            allPartition,
            addPartition,
            deletePartition,
        }
    },
    components: {
        NCard,
        NForm,
        NInput,
        NButton,
        NDrawer,
        NSelect,
        NFormItem,
        NDrawerContent,
    }
});
</script>

<style lang="less" scoped>
.partition-page {
    margin: 10px;
    width: calc(100% - 20px);
    border-radius: 10px;
    background-color: #fff;
}

.info-header {
    height: 30px;
    margin-bottom: 20px;
}

.partitions {
    // margin: 10px;
    width: 700px;
    border-top: 1px solid #e8e8e8;
    border-left: 1px solid #e8e8e8;
    border-right: 1px solid #e8e8e8;
}

.partition-item {
    display: flex;
    font-size: 16px;

    .item-left {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 200px;
        border-right: 1px solid #e8e8e8;
        border-bottom: 1px solid #e8e8e8;

        p {
            margin: 0;
        }

        button {
            margin: 0 10px;
        }
    }

    .item-right {
        width: 500px;
        border-bottom: 1px solid #e8e8e8;
        align-items: center;

        div {
            height: 40px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            border-bottom: 1px solid #e8e8e8;

            p {
                margin: 0;
            }

            &:last-child {
                border-bottom: none;
            }
        }

        button {
            margin: 0 10px;
        }
    }
}
</style>