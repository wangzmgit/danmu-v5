<template>
  <div class="bottom-bar">
    <span class="danmaku-amount">{{ amount }}条弹幕</span>
    <div class="danmaku-switch">
      <w-switch :value="danmaku" @change="setShow"></w-switch>
    </div>
    <!--弹幕设置-->
    <div v-show="showMenu" class="danmaku-menu">
      <div class="danmaku-menu-top">
        <p class="danmaku-menu-title">弹幕颜色</p>
        <div class="customize-color">
          <span style="color: #fff">#</span>
          <input type="text" v-model="danmakuForm.color" maxlength="6">
          <div :style="`background-color: #${danmakuForm.color}`"></div>
        </div>
      </div>
      <div class="color-btn">
        <div @click="setColor('fff')"></div>
        <div @click="setColor('e54256')"></div>
        <div @click="setColor('ffe133')"></div>
        <div @click="setColor('64dd17')"></div>
        <div @click="setColor('39ccff')"></div>
        <div @click="setColor('d500f9')"></div>
      </div>
      <p class="danmaku-menu-title">弹幕类型</p>
      <div class="danmaku-type">
        <ul class="type-switch">
          <li class="type-item" :class="{ 'active': danmakuForm.type === 0 }" @click="setType(0)">滚动</li>
          <li class="type-item" :class="{ 'active': danmakuForm.type === 1 }" @click="setType(1)">顶部</li>
          <li class="type-item" :class="{ 'active': danmakuForm.type === 2 }" @click="setType(2)">底部</li>
        </ul>
      </div>
      <p class="danmaku-menu-title">弹幕不透明度</p>
      <slider class="opacity" :value="opacity" @changeValue="setOpacity" />
    </div>
    <div class="danmaku-setting" @click="showMenu = !showMenu">
      <svg-icon class="danmaku-setting-icon" icon="setting"></svg-icon>
    </div>
    <input class="danmaku-input" v-model="danmakuForm.text" placeholder="发个友善的弹幕见证当下" />
    <w-button :disabled="!danmaku" class="send-btn" @click="sendDanmaku()">发送</w-button>
  </div>
</template>

<script>
import { ref, reactive } from "vue";
import Slider from "./Slider.vue";
import WSwitch from "../components/WSwitch.vue";
import SvgIcon from "../components/SvgIcon.vue";
import WButton from "../components/WButton.vue";
export default {
  emits: ['setOpacity', 'changeShow', 'showMsg', 'send'],
  props: {
    amount: {
      type: Number,
      default: 0
    },
    show: {
      type: Boolean,
      default: true
    },
  },
  setup(props, ctx) {
    const danmakuForm = reactive({
      text: "",
      color: "fff",
      type: 0,
    });
    const danmaku = ref(props.show);
    const showMenu = ref(false);
    const opacity = ref(100);

    const setType = (type) => {
      danmakuForm.type = type;
    }

    //设置弹幕不透明度
    const setOpacity = (val) => {
      ctx.emit('setOpacity', val);
    }

    //设置弹幕颜色
    const setColor = (color) => {
      danmakuForm.color = color;
    }

    //开启或关闭弹幕
    const setShow = (val) => {
      danmaku.value = val;
      ctx.emit('changeShow', val);
    }

    //发送弹幕
    const sendDanmaku = () => {
      if (danmakuForm.text == "") {
        ctx.emit('showMsg', "弹幕内容不能为空");
        return;
      }
      //验证颜色
      let reg = new RegExp("^([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$");
      if (danmakuForm.color.match(reg) == null) {
        return;
      }

      ctx.emit('send', danmakuForm);
      danmakuForm.text = "";
    }

    return {
      danmaku,
      showMenu,
      opacity,
      danmakuForm,
      setShow,
      setType,
      setOpacity,
      setColor,
      sendDanmaku
    }
  },
  components: {
    Slider,
    WSwitch,
    SvgIcon,
    WButton,

  },
};
</script>

<style lang="less" scoped>
.bottom-bar {
  position: relative;
  background: #fff;
  display: flex;
  width: 100%;
  height: 40px;
  align-items: center;
  border-bottom: 1px solid #ebebeb;
}

.danmaku-amount {
  padding-left: 12px;
  text-align: left;
  font-size: 12px;
  color: #999999;
  line-height: 40px;
  width: 120px;
}

.danmaku-switch {
  width: 46px;
}

.danmaku-setting {
  width: 30px;
  height: 30px;
  margin: 0 16px;
  cursor: pointer;
}

.danmaku-setting-icon {
  width: 30px;
  height: 30px;
}

.danmaku-input {
  height: 26px;
  width: calc(100% - 200px);
  border: 0;
  outline: none;
  padding: 0 2px;
  border-radius: 6px;
  background: #ebebeb;
}

.send-btn {
  width: 72px;
  height: 26px;
  margin-left: 10px;
}

.danmaku-menu {
  position: absolute;
  z-index: 20;
  bottom: 40px;
  background: rgba(12, 12, 12, 0.8);
  width: 240px;
  height: 240px;

  .danmaku-menu-top {
    display: flex;
    height: 46px;
    align-items: center;
    justify-content: space-between;
  }

  .danmaku-menu-title {
    color: #fff;
    margin: 12px 0 12px 10px;
  }

  .customize-color {
    display: flex;
    align-items: center;

    input {
      width: 80px;
      height: 24px;
      margin: 0 8px 0 2px;
      background-color: transparent;
      color: #fff;
      border: 1px solid rgba(161, 161, 161, 0.2);

      &:focus {
        border: 1px solid rgba(161, 161, 161, 0.2);
      }
    }

    div {
      width: 26px;
      height: 26px;
      border-radius: 50%;
      margin-right: 8px;
    }
  }

  .color-btn {
    display: flex;
    flex-wrap: nowrap;

    div {
      width: 30px;
      height: 30px;
      margin: 0 0 5px 8px;
      border-radius: 50%;
      cursor: pointer;

      &:nth-child(1) {
        background-color: white;
      }

      &:nth-child(2) {
        background-color: #e54256;
      }

      &:nth-child(3) {
        background-color: #ffe133;
      }

      &:nth-child(4) {
        background-color: #64dd17;
      }

      &:nth-child(5) {
        background-color: #39ccff;
      }

      &:nth-child(6) {
        background-color: #d500f9;
      }
    }
  }

  /**切换弹幕类型 */
  .danmaku-type {
    margin-left: 16px;
  }

  .type-switch {
    padding: 0;
    margin: 0;
    display: flex;
    list-style: none;
    width: 200px;
    border: 1px solid #fff;
    border-radius: 6px;
    overflow: hidden;
  }

  .type-item {
    flex: 1;
    color: #fff;
    padding: 10px;
    text-align: center;
    padding: 6px 6px;
  }

  .active {
    transition: all .3s;
    background-color: #18a058;
  }

  .opacity {
    width: 90% !important;
    margin: 0 auto;
  }
}
</style>