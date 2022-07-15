<template>
    <n-carousel autoplay show-arrow>
        <div class="carousel" v-for="(item, index) in carousels" :key="index">
            <img class="carousel-img" :src="item.img" alt="轮播图" />
        </div>
    </n-carousel>
</template>

<script lang="ts">
import { NCarousel } from 'naive-ui';
import { ref, defineComponent, onBeforeMount } from 'vue';
import { getCarouselAPI } from '@/api/carousel';
import { carouselType } from '@/types/carousel';

export default defineComponent({
    setup() {

        const carousels = ref(Array<carouselType>());
        const getCarousel = () => {
            getCarouselAPI().then((res) => {
                if (res.data.code === 2000) {
                    carousels.value = res.data.data.carousels;
                }
            })
        }

        onBeforeMount(() => {
            getCarousel();
        })

        return {
            carousels
        }
    },
    components: {
        NCarousel
    }
});
</script>

<style lang="less" scoped>
.carousel {
    width: 100%;
    height: 100%;

    .carousel-img {
        width: 100%;
        height: 100%;
        object-fit: fill;
    }
}
</style>