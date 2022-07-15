<template>
    <n-slider v-model:value="current" :vertical="vertical" @update:value="changeValue"></n-slider>
</template>

<script lang="ts">
import { NSlider } from 'naive-ui'
import { defineComponent, SetupContext } from 'vue';

export default defineComponent({
    emits: ['changeValue'],
    props: {
        val: {
            type: Number
        },
        vertical: {
            type: Boolean,
            default: false
        }
    },
    setup(props: any, ctx: SetupContext) {
        const changeValue = (val: number) => {
            ctx.emit('changeValue', val);
        }

        return {
            current: props.val,
            changeValue
        }
    },
    components: {
        NSlider
    }
});
</Script>