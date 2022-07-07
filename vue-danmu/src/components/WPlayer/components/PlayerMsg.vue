<template>
    <span class="msg">{{ msg }}</span>
</template>

<script>
export default {
    props: {
        msg: {
            type: String,
        },
    },
};
</script>

<style scoped>
.msg {
    z-index: 10;
    position: absolute;
    text-align: center;
    background: rgba(0, 0, 0, 0.6);
    width: 160px;
    height: 26px;
    left: 10px;
    bottom: 60px;
    line-height: 26px;
    border-radius: 2px;
    color: aliceblue;
}
</style>