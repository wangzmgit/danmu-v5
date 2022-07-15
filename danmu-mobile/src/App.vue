<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-message-provider>
      <router-view></router-view>
    </n-message-provider>
  </n-config-provider>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { zhCN, dateZhCN, NMessageProvider, NConfigProvider } from 'naive-ui';

export default defineComponent({
  name: 'App',
  setup() {
    return {
      zhCN,
      dateZhCN
    }
  },
  components: {
    NConfigProvider,
    NMessageProvider
  }
})
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: system-ui, -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
</style>
