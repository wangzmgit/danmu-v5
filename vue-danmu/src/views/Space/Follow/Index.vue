<template>
    <p class="follow-title">{{ `我的${following ? '关注' : '粉丝'}` }}</p>
    <follow-list :following="following" :uid="userInfo.uid"></follow-list>
</template>

<script>
import { useRoute } from "vue-router";
import storage from "@/utils/stored-data.js";
import FollowList from "@/components/FollowList";
import { ref, computed, onBeforeMount } from "vue";

export default {

    setup() {
        const route = useRoute();
        const following = ref(true);//是否为关注列表

        const userInfo = computed(() => {
            return storage.get("userInfo");
        })

        onBeforeMount(() => {
            if (route.name == "Followers") {
                following.value = false;
            }
        })

        return {
            following,
            userInfo
        }
    },
    components: {
        FollowList,
    }
};
</script>

<style scoped>
.follow-title {
    font-size: 18px;
    margin-top: 20px;
}
</style>