<template>
  <img :src="iconSrc" :alt="icon" />
</template>
<script>
import { computed } from "vue";
import play from "../icon/play.svg";
import pause from "../icon/pause.svg";
import volume from "../icon/volume.svg";
import replay from "../icon/replay.svg";
import setting from "../icon/setting.svg";
import fullScreen from "../icon/fullScreen.svg";

const Icons = {
  play: play,
  pause: pause,
  replay: replay,
  volume: volume,
  fullScreen: fullScreen,
  setting: setting,
};

export default {
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
};
</script>