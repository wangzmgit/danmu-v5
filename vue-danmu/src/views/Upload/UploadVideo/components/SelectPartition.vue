<template>
    <div class="select-partition">
        <n-select class="select" placeholder="选择分区" label-field="content" value-field="id" remote :options="partitions"
            @update:value="partitionChange" />
        <n-select v-if="showSubpartition" class="select" label-field="content" value-field="id" remote
            placeholder="选择二级分区" :options="subpartitions" @update:value="selectedPartition" />
    </div>
</template>

<script lang="ts">
import { onBeforeMount, ref, defineComponent } from "vue";
import { getPartitionAPI } from '@/api/partition';
import { NSelect, useNotification } from 'naive-ui';
import { partitionType } from "@/types/partition";

export default defineComponent({
    emits: ["selected"],
    setup(_props, ctx) {
        const partitions = ref<Array<partitionType>>([]);//分区
        const subpartitions = ref<Array<partitionType>>([]);//二级分区
        const showSubpartition = ref(false);//是否显示二级分区
        const notification = useNotification();//

        //改变分区
        const partitionChange = (parentId: number) => {
            getPartitionList(parentId);
            showSubpartition.value = true;
        }

        //获取分区列表
        const getPartitionList = (parentId: number) => {
            getPartitionAPI(parentId).then((res) => {
                if (res.data.code === 2000) {
                    if (parentId === 0) {
                        partitions.value = res.data.data.partitions;
                    } else {
                        subpartitions.value = res.data.data.partitions;
                    }
                }
            }).catch((err) => {
                notification.error({
                    title: '获取分区失败',
                    content: "原因:" + err.response.data.msg,
                    duration: 5000,
                })
            });
        }

        const selectedPartition = (value: number) => {
            ctx.emit("selected", value);
        }

        onBeforeMount(() => {
            getPartitionList(0);
        })

        return {
            partitions,
            subpartitions,
            showSubpartition,
            partitionChange,
            getPartitionList,
            selectedPartition,
        }

    },
    components: {
        NSelect
    }
});
</script>

<style lang="less" scoped>
.select-partition {
    width: 330px;
    display: flex;
    justify-content: space-between;

    .select {
        width: 150px;
    }
}
</style>