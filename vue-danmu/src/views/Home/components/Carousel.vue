<template>
    <n-carousel autoplay show-arrow>
        <div class="carousel" v-for="item in carousels">
            <img class="carousel-img" :src="item.img" alt="轮播图" />
        </div>
    </n-carousel>
</template>

<script>
import { NCarousel } from 'naive-ui';
import { ref, onBeforeMount } from 'vue';
import { getCarouselAPI } from '@/api/carousel';

export default {
    setup() {
        const carousels = ref([]);

        const getCarousel = () => {
            getCarouselAPI().then((res) => {
                if (res.data.code === 2000) {
                    carousels.value = res.data.data.carousels;
                }
            })
        }

        onBeforeMount(() => {
            getCarousel()
        })

        return {
            carousels
        }
    },
    components: {
        NCarousel
    }
}
</script>

<style lang="less" scoped>
.carousel {
    width: 100%;
    height: 280px;

    .carousel-img {
        width: 100%;
        height: 100%;
        object-fit: fill;
    }
}


</style>