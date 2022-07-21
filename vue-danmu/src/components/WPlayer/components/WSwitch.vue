<template>
  <div class="switch" @click="openSwitch()">
    <div class="switch-bg" :class="open ? 'open-bg' : ''">
      <span class="text" v-show="open">弹</span>
      <div class="switch-btn" :class="open ? 'open-btn' : ''"></div>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from 'vue';

export default defineComponent({
  emits: ['change'],
  props: {
    value: {
      type: Boolean,
      default: true,
    },
  },
  setup(props, ctx) {
    const open = ref(props.value);

    const openSwitch = () => {
      open.value = !open.value;
      ctx.emit("change", open.value);
    }
    return {
      open,
      openSwitch
    };
  }
});
</script>

<style lang="less" scoped>
.switch {
  position: relative;
  width: 44px;
  height: 22px;
  border-radius: 22px;
  border: 1px solid #e9e9e9;
  cursor: pointer;

  .switch-bg {
    background-color: #18a058;
    width: 20px;
    height: 22px;
    border-radius: 50px;
    transition: width 0.5s;
    -moz-transition: width 0.5s;
    -webkit-transition: width 0.5s;

    .switch-btn {
      position: absolute;
      top: -1px;
      left: -3px;
      width: 24px;
      height: 24px;
      border-radius: 50%;
      background-color: #fff;
      border: 1px solid #e9e9e9;
      transition: left 0.5s;
      -moz-transition: left 0.5s;
      /* Firefox 4 */
      -webkit-transition: left 0.5s;
      /* Safari and Chrome */
      -o-transition: left 0.5s;

      /* Opera */
    }

    .text {
      position: absolute;
      color: #fff;
      left: 6px;
    }



    .open-btn {
      left: 22px;
    }
  }

  .open-bg {
    width: 44px;
  }
}
</style>
