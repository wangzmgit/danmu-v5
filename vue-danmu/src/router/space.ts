import { RouteRecordRaw } from "vue-router";
const space = () => import("../views/Space/Index.vue");
const spaceFollow = () => import("../views/Space/Follow/Index.vue");
const spaceInfo = () => import("../views/Space/Info/Index.vue");
const setting = () => import("../views/Space/Setting/Index.vue");
const infoSetting = () => import("../views/Space/Setting/Info/Index.vue");
const securitySetting = () => import("../views/Space/Setting/Security/Index.vue");
const spaceCollection = () => import("../views/Space/Collection/Index.vue");
const spaceMessage = () => import("../views/Space/Message/Index.vue");
const spaceAnnounce = () => import("../views/Space/Announce/Index.vue");

const spaceRouter: Array<RouteRecordRaw> = [
    {
        path: "/space",
        name: "Space",
        meta: { auth: true },
        component: space,
        redirect: '/space/info',
        children: [
            {
                path: '/space/info',
                name: 'SpaceInfo',
                meta: { auth: true },
                component: spaceInfo,
            },
            {
                path: '/space/collection',
                name: 'Collection',
                meta: { auth: true },
                component: spaceCollection,
            },
            {
                path: '/space/message',
                name: 'Message',
                meta: { auth: true },
                component: spaceMessage,
            },
            {
                path: '/space/announce',
                name: 'Announce',
                meta: { auth: true },
                component: spaceAnnounce,
            },
            {
                path: '/space/setting',
                name: 'Setting',
                meta: { auth: true },
                component: setting,
                redirect: '/space/setting/info',
                children: [
                    {
                        path: '/space/setting/info',
                        name: 'InfoSetting',
                        meta: { auth: true },
                        component: infoSetting,
                    },
                    {
                        path: '/space/setting/security',
                        name: 'SecuritySetting',
                        meta: { auth: true },
                        component: securitySetting,
                    },
                ]
            },
            {
                path: '/space/following',
                name: 'Following',
                meta: { auth: true },
                component: spaceFollow,
            },
            {
                path: '/space/followers',
                name: 'Followers',
                meta: { auth: true },
                component: spaceFollow,
            },
        ]
    }
]

export default spaceRouter;