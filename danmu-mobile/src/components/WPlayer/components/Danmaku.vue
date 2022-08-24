<template>
  <div class="danmaku-container" ref="danmakuRef"></div>
</template>

<script lang="ts">
import { ref, reactive, defineComponent } from "vue";
import { danmakuItemType, drawDanmakuType } from "../types/danmaku";

export default defineComponent({
  props: {
    list: {
      type: Array as () => Array<danmakuItemType>
    }
  },
  setup(props) {
    const danmakuRef = ref<HTMLElement | null>(null);
    const paused = ref(false);//是否暂停
    const danmaku = ref<HTMLElement[]>([]);//当前弹幕
    const currentTime = ref(0);//当前时间
    const danmakuTunnel = reactive({
      row: [] as Array<number>, //轨道结束的时间
      top: [] as Array<number>,
      bottom: [] as Array<number>,
    })

    // 播放暂停
    const startOrPause = (start: boolean) => {
      const state = start ? 'running' : 'paused';
      for (let i = 0; i < danmaku.value.length; i++) {
        (danmaku.value[i] as HTMLElement).style.animationPlayState = state;
      }
      paused.value = !start;
    }

    //清除弹幕
    const clearDanmaku = () => {
      danmakuTunnel.row = [];
      danmakuTunnel.top = [];
      danmakuTunnel.bottom = [];
      danmakuRef.value!.innerHTML = "";
    }

    //设置弹幕不透明度
    const setOpacity = (opacity: number) => {
      if (danmakuRef.value) {
        danmakuRef.value.style.opacity = (opacity * 0.01).toString();
      }
    }

    //更新时间
    const timeUpdate = (time: number) => {
      if (Math.round(time) !== currentTime.value) {
        currentTime.value = Math.round(time);
        //绘制弹幕
        if (!props.list) {
          return;
        }
        const currentDanmaku = props.list.filter((item: danmakuItemType) => {
          return item.time === currentTime.value;
        })

        currentDanmaku.map((item: danmakuItemType) => {
          drawDanmaku(item, false);
        })
      }
    }

    //获取滚动弹道
    const getRowTunnel = (text: string) => {
      //当前弹幕结束时间
      let duration = 30 - text.length * 0.5;
      let width = danmakuRef.value!.offsetWidth;
      duration = Math.ceil((width + text.length * 20) / (3000 / duration)) + currentTime.value;
      //计算弹道数量
      let tunnel = Math.floor(danmakuRef.value!.offsetHeight / 26);
      //遍历轨道
      for (let i = 0; i < danmakuTunnel.row.length; i++) {
        //如果当前轨道结束时间小于新弹幕的结束时间
        //说明弹幕可以放在该弹道
        if (danmakuTunnel.row[i] < duration) {
          danmakuTunnel.row[i] = duration;
          return i;
        }
      }
      //如果没有则尝试新增加弹道
      if (danmakuTunnel.row.length < tunnel) {
        danmakuTunnel.row.push(duration);
        return danmakuTunnel.row.length - 1;
      }
      //如果不可以新增弹道，则使用随机弹道
      return Math.round(Math.random() * tunnel);
    }

    //获取固定弹道
    const getFixedTunnel = (type: number) => {
      //当前弹幕结束时间
      let duration = currentTime.value + 5;
      //计算弹道数量
      let tunnel = Math.floor(danmakuRef.value!.offsetHeight / 26);
      switch (type) {
        case 1:
          //遍历轨道
          for (let i = 0; i < danmakuTunnel.row.length; i++) {
            //如果当前轨道结束时间小于新弹幕的结束时间
            //说明弹幕可以放在该弹道
            if (danmakuTunnel.top[i] < duration) {
              danmakuTunnel.top[i] = duration;
              return i;
            }
          }
          //如果没有则尝试新增加弹道
          if (danmakuTunnel.top.length < tunnel) {
            danmakuTunnel.top.push(duration);
            return danmakuTunnel.top.length - 1;
          }
          break;
        case 2:
          //遍历底部弹幕轨道
          for (let i = 0; i < danmakuTunnel.bottom.length; i++) {
            //如果当前轨道结束时间小于新弹幕的结束时间
            //说明弹幕可以放在该弹道
            if (danmakuTunnel.bottom[i] < duration) {
              danmakuTunnel.bottom[i] = duration;
              return i;
            }
          }
          //如果没有则尝试新增加弹道
          if (danmakuTunnel.bottom.length < tunnel) {
            danmakuTunnel.bottom.push(duration);
            return danmakuTunnel.bottom.length - 1;
          }
          break;
      }
      //如果不可以新增弹道，则使用随机弹道
      return Math.round(Math.random() * tunnel);
    }

    const drawDanmaku = (draw: drawDanmakuType, send: boolean) => {
      let width = danmakuRef.value!.offsetWidth;
      var item = document.createElement("span");
      var content = document.createTextNode(draw.text);
      item.style.color = draw.color;
      item.appendChild(content);
      item.className = "danmaku-item";
      //滚动弹幕
      if (draw.type == 0) {
        //设置轨道
        let rowTunnel = getRowTunnel(draw.text);
        item.style.top = `${rowTunnel * 26}px`;
        item.classList.add("danmaku-row");
        item.style.transform = `translateX(-${width}px)`;
        danmaku.value[rowTunnel] = item;
        danmakuRef.value!.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(rowTunnel);
          danmakuRef.value!.removeChild(item);
        });
        if (send) {
          item.style.border = "1px solid red";
        }
        item.classList.add("danmaku-row-move");
      } else if (draw.type == 1) {
        let topTunnel = getFixedTunnel(draw.type);
        item.style.width = "100%";
        item.style.textAlign = "center";
        item.style.top = `${topTunnel * 26}px`;
        danmaku.value[topTunnel] = item;
        danmakuRef.value!.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(topTunnel);
          danmakuRef.value!.removeChild(item);
        });
        item.classList.add("danmaku-center-move");
      } else if (draw.type == 2) {
        let bottomTunnel = getFixedTunnel(draw.type);
        item.style.width = "100%";
        item.style.textAlign = "center";
        item.style.bottom = `${bottomTunnel * 26}px`;
        danmaku.value[bottomTunnel] = item;
        danmakuRef.value!.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(bottomTunnel);
          danmakuRef.value!.removeChild(item);
        });
        item.classList.add("danmaku-center-move");
      }
      if (paused.value) {
        item.style.animationPlayState = "paused";
      }
    }

    return {
      danmakuRef,
      startOrPause,
      clearDanmaku,
      timeUpdate,
      drawDanmaku,
      setOpacity
    }
  }
});
</script>

<style>
.danmaku-container {
  z-index: 1;
  position: absolute;
  overflow: hidden;
  width: 100%;
  height: 100%;
}

.danmaku-item {
  position: absolute;
  font-size: 22px;
  white-space: nowrap;
  text-shadow: 1px 1px 0 #000, -1px -1px 0 #000, -1px 1px 0 #000, 1px -1px 0 #000;
}

.danmaku-row {
  right: 0;
}

.danmaku-row-move {
  will-change: transform;
  animation: danmaku 5s linear;
  animation-play-state: running;
}

.danmaku-center-move {
  will-change: transform;
  animation: danmaku-center 5s linear;
  animation-play-state: running;
}

@keyframes danmaku {
  from {
    transform: none;
  }
}

@keyframes danmaku-center {
  from {
    visibility: visible;
  }

  to {
    visibility: visible;
  }
}
</style>

