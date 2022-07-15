import storage from '@/utils/stored-data';
import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router';

const home = () => import("../views/Home/Index.vue");
const login = () => import("../views/Login/Index.vue");
const register = () => import("../views/Register/Index.vue");
const search = () => import("../views/Search/Index.vue");
const video = () => import("../views/Video/Index.vue");
const spaceInfo = () => import("../views/Space/Info/Index.vue");
const editInfo = () => import("../views/Space/Edit/Index.vue");
const message = () => import("../views/Message/Index.vue");
const announce = () => import("../views/Message/Announce/Index.vue");
const messageDetails = () => import("../views/Message/Details/Index.vue");

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: home
  },
  {
    path: '/login',
    name: "Login",
    component: login
  },
  {
    path: '/register',
    name: "Register",
    component: register
  },
  {
    path: '/search',
    name: "Search",
    component: search
  },
  {
    path: '/video/VID:vid',
    name: "Video",
    component: video
  },
  //消息相关
  {
    path: '/message',
    name: 'Message',
    component: message,
  },
  {
    path: '/announce',
    name: "Announce",
    component: announce
  },
  {
    path: '/message/details',
    name: 'MsgDetails',
    meta: { auth: true },
    component: messageDetails
  },
  {
    path: '/space',
    name: 'Space',
    meta: { auth: true },
    redirect: '/space/info',
    children: [
      {
        path: '/space/info',
        name: "SpaceInfo",
        meta: { auth: true },
        component: spaceInfo
      },
      {
        path: '/space/edit',
        name: "EditInfo",
        meta: { auth: true },
        component: editInfo
      },
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

router.beforeEach((to) => {
  //是否需要登录
  if (to.meta.auth && !storage.get('token')) {
    router.push({ name: 'Login' });
  }
})

export default router
