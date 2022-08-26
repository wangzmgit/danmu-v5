import storage from "@/utils/stored-data";
import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

const home = () => import("../views/Home/Index.vue");
const login = () => import("../views/Login/Index.vue");
const review = () => import("../views/Review/Index.vue");
const user = () => import("../views/User/Index.vue");
const video = () => import("../views/Video/Index.vue");
const partition = () => import("../views/Partition/Index.vue");
const carousel = () => import("../views/Carousel/Index.vue");
const announce = () => import("../views/Announce/Index.vue");
const dashboard = () => import("../views/Dashboard/Index.vue");
const about = () => import("../views/About/Index.vue");

const routes: Array<RouteRecordRaw> = [
    { path: "/", redirect: "/home" },
    {
        path: "/login",
        name: "Login",
        component: login
    },
    {
        path: "/home",
        name: "Home",
        meta: { admin: true },
        component: home,
        redirect: '/home/dashboard',
        children: [
            {
                path: "/home/dashboard",
                name: "Dashboard",
                meta: { admin: true },
                component: dashboard
            },
            {
                path: "/home/review",
                name: "Review",
                meta: { admin: true },
                component: review
            },
            {
                path: "/home/video",
                name: "Video",
                meta: { admin: true },
                component: video
            },
            {
                path: "/home/user",
                name: "User",
                meta: { admin: true },
                component: user
            },
            {
                path: "/home/announce",
                name: "Announce",
                meta: { admin: true },
                component: announce
            },
            {
                path: "/home/partition",
                name: "Partition",
                meta: { admin: true },
                component: partition
            },
            {
                path: "/home/carousel",
                name: "Carousel",
                meta: { admin: true },
                component: carousel
            },
            {
                path: "/home/about",
                name: "About",
                meta: { admin: true },
                component: about
            },
        ]
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

router.beforeEach((to, from) => {
    //是否需要登录
    if (to.meta.admin && !storage.get('admin_access_token') && !storage.get('admin_refresh_token')) {
        router.push({ name: 'Login' });
    }
})


export default router