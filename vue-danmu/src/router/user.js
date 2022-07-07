const user = () => import("../views/User/Index");
const userVideo = () => import("../views/User/UserVideo/Index");
const userFollow = () => import("../views/User/UserFollow/Index");
const mentionUser = () => import("../views/User/MentionUser/Index");
const userRouter = [
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