import storage from "@/utils/stored-data";
import { createRouter, createWebHashHistory } from "vue-router";

const home = () => import("../views/Home/Index");
const login = () => import("../views/Login/Index");
const review = () => import("../views/Review/Index");
const user = () => import("../views/User/Index");
const video = () => import("../views/Video/Index");
const partition = () => import("../views/Partition/Index");
const carousel = () => import("../views/Carousel/Index");
const announce = () => import("../views/Announce/Index");
const dashboard = () => import("../views/Dashboard/Index");
const about = () => import("../views/About/Index");

const routes = [
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
    if (to.meta.admin && !storage.get('admin')) {
        router.push({ name: 'Login' });
    }
})


export default router