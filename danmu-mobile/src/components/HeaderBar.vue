<template>
  <div class="header-box">
    <common-avatar :url="userInfo.avatar" :size="36" :iconsize="20" @click="headerRouter('Space')"></common-avatar>
    <div class="search-box">
      <n-input v-model:value="keywords" round placeholder="搜索关键词~" @keydown.enter="search">
        <template #suffix>
          <n-icon @click="search">
            <search />
          </n-icon>
        </template>
      </n-input>
    </div>
    <n-icon class="msg" @click="headerRouter('Message')">
      <mail-outline />
    </n-icon>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import storage from '@/utils/stored-data';
import CommonAvatar from "./CommonAvatar.vue";
import { userInfoType } from "@/types/user";
import { NIcon, NInput } from "naive-ui";
import { MailOutline, Search } from '@vicons/ionicons5'

export default defineComponent({
  setup() {
    const route = useRoute();
    const router = useRouter();

    const userInfo = computed((): userInfoType => {

      const info: userInfoType = {
        uid: 0,
        name: "",
        avatar: "",
        gender: 0,
        sign: "",
        birthday: ""
      }

      if (storage.get("userInfo")) {
        return storage.get("userInfo") as userInfoType;
      } else {
        return info;
      }
    })

    const headerRouter = (page: string) => {
      const path = route.name;
      switch (page) {
        case 'Space':
          if (path == "SpaceInfo")
            page = "Home";
          break;
        case 'Message':
          if (path == "Message")
            page = "Home";
          break;
      }
      router.push({ name: page });
    }

    // 搜索
    const keywords = ref("");
    const search = () => {
      router.push({ name: "Search", query: { keywords: keywords.value } });
    }

    return {
      userInfo,
      keywords,
      search,
      headerRouter
    }
  },
  components: {
    NIcon,
    NInput,
    Search,
    CommonAvatar,
    MailOutline,
  }
});
</script>

<style lang="less" scoped>
.header-box {
  width: calc(100% - 20px);
  height: 50px;
  display: flex;
  align-items: center;
  padding: 0 10px;
  background-color: #36ad6a;
  justify-content: space-between;
  -webkit-box-shadow: 0px 0px 3px #c8c8c8;
  -moz-box-shadow: 0px 0px 3px #c8c8c8;
  box-shadow: 0px 0px 3px #c8c8c8;

  .search-box {
    width: calc(100% - 130px);
  }

  .msg {
    color: #fff;
    font-size: 26px;
  }
}
</style>