<template>
    <follow-list :following="following" :height="400" :uid="uid"></follow-list>
</template>

<script lang="ts">
import { useRoute } from "vue-router";
import { ref, onBeforeMount, defineComponent } from "vue";
import FollowList from "@/components/FollowList.vue";

export default defineComponent({
    setup() {
        const uid = ref(0);
        const following = ref(true);
        const route = useRoute();

        onBeforeMount(() => {
            uid.value = Number(route.params.uid);
            if (route.name === "UserFollowers") {
                following.value = false;
            }
        })

        return {
            uid,
            following
        }

    },
    components: {
        FollowList,
    },
});
</script>
