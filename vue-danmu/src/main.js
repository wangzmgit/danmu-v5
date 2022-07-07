import { createApp } from 'vue';
import App from './App.vue';
import store from './store';
import router from './router';
import VueDOMPurifyHTML from 'vue-dompurify-html';

const app = createApp(App);
//动态标题
app.directive('title', {
    mounted(el) {
        document.title = el.dataset.title
    }
})

app.use(store);
app.use(router);
app.use(VueDOMPurifyHTML);//防止v-html的xss注入
app.mount('#app');