import storage from "@/utils/stored-data";
import userRoutes from './user.js';//用户空间
import spaceRoutes from './space.js';//个人空间、
import uploadRoutes from "./upload.js";
import { createRouter, createWebHashHistory } from "vue-router";

const home = () => import("../views/Home/Index");
const video = () => import("../views/Video/Index");
const search = () => import("../views/Search/Index");
const collection = () => import("../views/Collection/Index");
const videoList = () => import("../views/VideoList/Index");
const pageNotFound = () => import("../views/Result/PageNotFound");

const baseRoutes = [
    { path: "/", redirect: "/home" },
    {
        path: "/home",
        name: "Home",
        component: home
    },
    {
        path: "/video/VID:vid",
        name: "Video",
        component: video
    },
    {
        path: "/collection/:id",
        name: "CollectionDetails",
        component: collection
    },
    {
        path: '/search/:keywords',
        name: 'Search',
        component: search
    },
    {
        path: '/video/list',
        name: 'VideoList',
        component: videoList
    },
    {
        path: '/404',
        name: '404',
        component: pageNotFound
    },
    {
        path: '/:catchAll(.*)',
        redirect: {
            name: "404"
        }
    },
]

const routes = baseRoutes.concat(userRoutes, spaceRoutes, uploadRoutes);

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

router.beforeEach((to, from) => {
    //是否需要登录
    if (to.meta.auth && !storage.get('token')) {
        router.push({ name: 'Home', params: { login: true } });
    }
})

export default router