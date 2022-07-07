<template>
    <div class="mask" @mouseover="over()" @mouseleave="leave()">
        <slot></slot>
        <span v-show="showMask" @click="handleClick" class="mask-action">
            <slot name="action"></slot>
        </span>
        <span :class="showMask ? 'mask-hide' : ''">
            <slot name="hide"></slot>
        </span>
    </div>
</template>

<script>
import { ref } from 'vue';
export default {
    setup(_props, ctx) {
        const showMask = ref(false);

        const handleClick = () => {
            ctx.emit('click')
        }

        const over = () => {
            showMask.value = true;
        }

        const leave = () => {
            showMask.value = false;
        }

        return {
            showMask,
            over,
            leave,
            handleClick,
        }
    },

}
</script>

<style>
.mask {
    width: 100%;
    overflow: hidden;
    position: relative;
    display: inline-block;
    vertical-align: middle;
}

.mask-action {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    cursor: pointer;
    opacity: 1;
    font-size: 20px;
    z-index: 1;
    background-color: rgba(0, 0, 0, 0.4);
    animation: activation .3s cubic-bezier(.42, 0, .58, 1);
}

.mask-hide {
    opacity: 0;
    transition: opacity .6s;
}

@keyframes activation {
    0% {
        transform: translateY(100%);
    }

    100% {
        transform: translateY(0);
    }
}
</style>