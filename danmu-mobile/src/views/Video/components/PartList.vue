<template>
    <div class="part-list">
        <div class="part-item" v-for="(item, index) in resources" :key="index" @click="changePart(index)">
            <span :class="{ 'active': current - 1 === index }">P{{ index + 1 }}</span>
        </div>
    </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";

export default defineComponent({
    emits: ['change'],
    props: {
        resources: {
            type: Array
        },
        active: {
            type: Number,
            default: 1
        }
    },
    setup(props, ctx) {
        const current = ref(props.active);
        const changePart = (part: number) => {
            current.value = part + 1;
            ctx.emit('change', part + 1)
        }

        return {
            current,
            changePart
        }
    }
});
</script>

<style lang="less" scoped>
.part-list {
    width: 100%;
    overflow: hidden;
    display: flex;
    align-items: center;
    overflow: auto;
    padding-left: 0;
    list-style: none;
    scrollbar-width: none; //隐藏滚动条(火狐)
    -ms-overflow-style: none; //隐藏滚动条(IE)

    &::-webkit-scrollbar {
        //隐藏滚动条(Chrome)
        display: none;
    }

    .part-item {
        text-align: center;
        font-size: 16px;
        padding: 10px 3px;
        margin: 3px;

        span {
            padding: 0 10px;
            overflow: hidden;
            height: 36px;
            line-height: 36px;
            width: 100px;
            display: block;
            border-radius: 3px;
            border: 1px solid #808080;
        }

        .active {
            color: #36ad6a;

            border: 1px solid #36ad6a;
        }
    }

}
</style>