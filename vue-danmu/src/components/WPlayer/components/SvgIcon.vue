<template>
  <img :src="iconSrc" :alt="icon" />
</template>
<script lang="ts">
import play from "../icon/play.svg";
import pause from "../icon/pause.svg";
import volume from "../icon/volume.svg";
import replay from "../icon/replay.svg";
import setting from "../icon/setting.svg";
import fullScreen from "../icon/fullScreen.svg";
import rowOpen from '../icon/rowOpen.svg';
import rowClose from '../icon/rowClose.svg';
import topOpen from '../icon/topOpen.svg';
import topClose from '../icon/topClose.svg';
import bottomOpen from '../icon/bottomOpen.svg';
import bottomClose from '../icon/bottomClose.svg';
import colorOpen from '../icon/colorOpen.svg';
import colorClose from '../icon/colorClose.svg';

import { computed, defineComponent } from "vue";

const Icons: any = {
  play: play,
  pause: pause,
  replay: replay,
  volume: volume,
  fullScreen: fullScreen,
  setting: setting,
  rowOpen: rowOpen,
  rowClose: rowClose,
  topOpen: topOpen,
  topClose: topClose,
  bottomOpen: bottomOpen,
  bottomClose: bottomClose,
  colorOpen: colorOpen,
  colorClose: colorClose
};

export default defineComponent({
  props: {
    icon: {
      type: String,
      default: null,
    },
  },
  setup(props) {
    const iconSrc = computed(() => {
      return Icons[`${props.icon}`]
    })

    return {
      iconSrc
    }
  }
});
</script>