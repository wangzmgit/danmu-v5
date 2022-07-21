import { RouteRecordRaw } from "vue-router";
const user = () => import("../views/User/Index.vue");
const userVideo = () => import("../views/User/UserVideo/Index.vue");
const userFollow = () => import("../views/User/UserFollow/Index.vue");
const mentionUser = () => import("../views/User/MentionUser/Index.vue");

const userRouter: Array<RouteRecordRaw> = [
    {
        path: "/user/:uid",
        name: "User",
        component: user,
        redirect: { name: 'UserVideo' },
        children: [
            {
                path: '/user/:uid/video',
                name: 'UserVideo',
                component: userVideo,
            },
            {
                path: '/user/:uid/following',
                name: 'UserFollowing',
                component: userFollow,
            },
            {
                path: '/user/:uid/followers',
                name: 'UserFollowers',
                component: userFollow,
            },
        ]
    },
    {
        path: '/user/:name/mention',
        name: 'MentionUser',
        component: mentionUser,

    }
]

export default userRouter;