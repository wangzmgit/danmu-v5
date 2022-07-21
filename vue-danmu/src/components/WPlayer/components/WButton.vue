<template>
  <button type="button" :class="['w-button', type]" :disabled="disabled">
    <slot></slot>
  </button>
</template>
<script lang="ts">
import { defineComponent, } from 'vue';

export default defineComponent({
  name: "WButton",
  props: {
    type: {
      type: String,
      default: "primary",
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  }
});
</script>

<style lang="less" scoped>
.w-button {
  font-weight: 400;
  font-size: 14px;
  height: 32px;
  outline: none;
  cursor: pointer;
  padding: 0 15px;
  border-radius: 4px;
  border: 1px solid #d9d9d9;
}

button[disabled] {
  opacity: 0.6;
  cursor: default;
}

.primary {
  color: #fff;
  background: #18a058;
  border-color: #18a058;

  &:hover {
    background: #36ad6a;
    border-color: #36ad6a;
  }

}

.text {
  color: #fff;
  background: transparent;
  border-color: transparent;
}
</style>
