<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-notification-provider>
      <router-view></router-view>
    </n-notification-provider>
  </n-config-provider>
</template>

<script lang="ts">
import {
  zhCN, dateZhCN,
  NNotificationProvider, NConfigProvider
} from 'naive-ui';
import { defineComponent } from 'vue';
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
    NNotificationProvider
  }
});
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: system-ui, -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
}
</style>
