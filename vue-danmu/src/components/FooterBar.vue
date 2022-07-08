<template>
    <footer class="footer">
        <div class="links link-list">
            <span v-if="config.mobile">
                <a :href="config.mobile">移动端</a>
            </span>
            <span>|</span>
            <span>
                <a href="http://kuukaa.fun" target="_blank">使用文档</a>
            </span>
            <span>|</span>
            <span>
                Powered by <a href="https://github.com/wangzmgit/danmu-v5" target="_blank">wzmgit/danmu-v5</a>
            </span>
        </div>
        <div class="copyright link-list">
            <span v-show="config.icp">
                <a href="https://beian.miit.gov.cn/#/Integrated/index" target="_blank">{{ config.icp }}</a>
            </span>
            <span v-show="config.icp">|</span>
            <span v-show="config.security" class="security">
                <img class="security" src="@/assets/filing.png" alt="公网安备" />
                <a href="https://www.beian.gov.cn/portal/registerSystemInfo" target="_blank">{{ config.security }}</a>
            </span>
        </div>
    </footer>
</template>

<script>
import config from '@/config'
export default {
    name: 'FooterBar',
    setup() {
        return {
            config
        }
    }
}
</script>

<style lang="less" scoped>
.footer {
    color: #666;
    width: 100%;
    min-width: 1200px;
    height: 80px;
    text-align: center;
    background-color: #f6f7f8;

    .links {
        padding-top: 16px;
    }

    .copyright {
        padding-top: 8px;

        .security {
            img {
                width: 12px;
                padding: 0 3px;
            }
        }
    }

    .link-list {
        span {
            margin: 0 3px;

            a {
                color: #666;
                text-decoration: none;

                &:hover {
                    text-decoration: underline;
                }
            }
        }
    }
}
</style>