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
import { computed, defineComponent } from "vue";

const Icons: any = {
  play: play,
  pause: pause,
  replay: replay,
  volume: volume,
  fullScreen: fullScreen,
  setting: setting,
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