<template>
    <n-slider v-model:value="value" :vertical="vertical" @update:value="changeValue"></n-slider>
</template>

<script>
import { NSlider } from 'naive-ui'

export default {
    emits: ['changeValue'],
    props: {
        value: {
            type: Number
        },
        vertical: {
            type: Boolean,
            default: false
        }
    },
    setup(_props, ctx) {
        const changeValue = (val) => {
            ctx.emit('changeValue', val);
        }

        return {
            changeValue
        }
    },
    components: {
        NSlider
    }
}
</Script>