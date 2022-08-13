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
        <ul class="type-switch" @click="setType">
          <li class="type-item" name="0" :class="{ 'active': danmakuForm.type === 0 }">滚动</li>
          <li class="type-item" name="1" :class="{ 'active': danmakuForm.type === 1 }">顶部</li>
          <li class="type-item" name="2" :class="{ 'active': danmakuForm.type === 2 }">底部</li>
        </ul>
      </div>
      <p class="danmaku-menu-title">屏蔽弹幕类型</p>
      <div class="danmaku-filter" @click="setDisableType">
        <svg-icon name="row" class="filter-item" :icon="`row${disableType.row ? 'Close' : 'Open'}`"></svg-icon>
        <svg-icon name="top" class="filter-item" :icon="`top${disableType.top ? 'Close' : 'Open'}`"></svg-icon>
        <svg-icon name="bottom" class="filter-item" :icon="`bottom${disableType.bottom ? 'Close' : 'Open'}`"></svg-icon>
        <svg-icon name="color" class="filter-item" :icon="`color${disableType.color ? 'Close' : 'Open'}`"></svg-icon>
      </div>
      <p class="danmaku-menu-title">弹幕不透明度</p>
      <w-slider class="opacity" :value="opacity" @changeValue="setOpacity" />
    </div>
    <div class="danmaku-setting" @click="showMenu = !showMenu">
      <svg-icon class="danmaku-setting-icon" icon="setting"></svg-icon>
    </div>
    <input class="danmaku-input" v-model="danmakuForm.text" placeholder="发个友善的弹幕见证当下" />
    <w-button :disabled="!danmaku" class="send-btn" @click="sendDanmaku()">发送</w-button>
  </div>
</template>

<script lang="ts">
import { ref, reactive, defineComponent } from "vue";
import WSlider from "./WSlider.vue";
import useConfig from '../hooks/config';
import WSwitch from "../components/WSwitch.vue";
import SvgIcon from "../components/SvgIcon.vue";
import WButton from "../components/WButton.vue";
import { danmakuType } from "../types/danmaku";


export default defineComponent({
  emits: ['setOpacity', 'changeShow', 'showMsg', 'send', 'changeDisableType'],
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
    const danmakuForm = reactive<danmakuType>({
      vid: 0,
      time: 0,
      text: "",
      color: "fff",
      type: 0,
      part: 1,
    });
    const { getConfig } = useConfig();
    const danmaku = ref(props.show);
    const showMenu = ref(false);
    const opacity = ref(100);
    const disableType = reactive(getConfig().disableType);

    const setType = (e: Event) => {
      danmakuForm.type = parseInt((e.target as HTMLElement).getAttribute('name')!);
    }

    //设置弹幕不透明度
    const setOpacity = (val: number) => {
      ctx.emit('setOpacity', val);
    }

    //设置弹幕颜色
    const setColor = (color: string) => {
      danmakuForm.color = color;
    }

    //开启或关闭弹幕
    const setShow = (val: boolean) => {
      danmaku.value = val;
      ctx.emit('changeShow', val);
    }

    //设置屏蔽弹幕类型
    const setDisableType = (e: Event) => {
      const eElement = e.target as HTMLElement;
      if (eElement.className === 'filter-item') {
        switch (eElement.getAttribute("name")) {
          case 'row':
            disableType.row = !disableType.row;
            break;
          case 'top':
            disableType.top = !disableType.top;
            break;
          case 'bottom':
            disableType.bottom = !disableType.bottom;
            break;
          case 'color':
            disableType.color = !disableType.color;
            break;
          default:
            break;
        }
      }
      ctx.emit('changeDisableType', { ...disableType });
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
      disableType,
      danmakuForm,
      setShow,
      setType,
      setOpacity,
      setColor,
      setDisableType,
      sendDanmaku
    }
  },
  components: {
    WSlider,
    WSwitch,
    SvgIcon,
    WButton,
  },
});
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
  width: 250px;
  height: 300px;

  .danmaku-menu-top {
    display: flex;
    height: 46px;
    align-items: center;
    justify-content: space-between;
  }

  .danmaku-menu-title {
    color: #fff;
    margin: 8px 0 8px 10px;
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
    width: calc(100% - 20px);
    margin-left: 10px;
    display: flex;
    flex-wrap: nowrap;
    justify-content: space-around;

    div {
      width: 30px;
      height: 30px;
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

    .type-switch {
      padding: 0;
      margin: 0;
      display: flex;
      list-style: none;
      width: 218px;
      border: 1px solid #fff;
      border-radius: 6px;
      overflow: hidden;

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
    }
  }

  /**不透明度 */
  .opacity {
    width: 90% !important;
    margin: 0 auto;
  }

  /**弹幕屏蔽 */
  .danmaku-filter {
    display: flex;
    justify-content: space-around;

    .filter-item {
      width: 36px;
      height: 36px;
      cursor: pointer;
    }
  }
}
</style>