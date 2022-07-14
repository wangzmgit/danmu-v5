<template>
  <div class="danmaku-container" ref="danmakuRef"></div>
</template>

<script>
import { ref, reactive } from "vue";

export default {
  props: {
    list: {
      type: Array
    }
  },
  setup(props) {
    const danmakuRef = ref(null);
    const paused = ref(false);//是否暂停
    const danmaku = ref([]);//当前弹幕
    const currentTime = ref(0);//当前时间
    const danmakuTunnel = reactive({
      row: [], //轨道结束的时间
      top: [],
      bottom: [],
    })

    // 播放暂停
    const startOrPause = (start) => {
      const state = start ? 'running' : 'paused';
      for (let i = 0; i < danmaku.value.length; i++) {
        danmaku.value[i].style.animationPlayState = state;
      }
      paused.value = !start;
    }

    //清除弹幕
    const clearDanmaku = () => {
      danmakuTunnel.row = [];
      danmakuTunnel.top = [];
      danmakuTunnel.bottom = [];
      danmakuRef.value.innerHTML = "";
    }

    //设置弹幕不透明度
    const setOpacity = (opacity) => {
      if (danmakuRef.value) {
        danmakuRef.value.style.opacity = parseFloat(opacity) * 0.01;
      }
    }

    //更新时间
    const timeUpdate = (time) => {
      if (Math.round(time) !== currentTime.value) {
        currentTime.value = Math.round(time);
        //绘制弹幕
        if (!props.list) {
          return;
        }
        const currentDanmaku = props.list.filter((item) => {
          return item.time === currentTime.value;
        })

        currentDanmaku.map((item) => {
          drawDanmaku(item);
        })
      }
    }

    //获取滚动弹道
    const getRowTunnel = (text) => {
      //当前弹幕结束时间
      let duration = 30 - text.length * 0.5;
      let width = danmakuRef.value.offsetWidth;
      duration = Math.ceil((width + text.length * 20) / (3000 / duration)) + currentTime.value;
      //计算弹道数量
      let tunnnel = Math.floor(danmakuRef.value.offsetHeight / 26);
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
      if (danmakuTunnel.row.length < tunnnel) {
        danmakuTunnel.row.push(duration);
        return danmakuTunnel.row.length - 1;
      }
      //如果不可以新增弹道，则使用随机弹道
      return Math.round(Math.random() * tunnnel);
    }

    //获取固定弹道
    const getFixedTunnel = (type) => {
      //当前弹幕结束时间
      let duration = currentTime.value + 5;
      //计算弹道数量
      let tunnnel = Math.floor(danmakuRef.value.offsetHeight / 26);
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
          if (danmakuTunnel.top.length < tunnnel) {
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
          if (danmakuTunnel.bottom.length < tunnnel) {
            danmakuTunnel.bottom.push(duration);
            return danmakuTunnel.bottom.length - 1;
          }
          break;
      }
      //如果不可以新增弹道，则使用随机弹道
      return Math.round(Math.random() * tunnnel);
    }

    const drawDanmaku = ({ text, color, type }, send = false) => {
      let width = danmakuRef.value.offsetWidth;
      var item = document.createElement("span");
      var content = document.createTextNode(text);
      item.style.color = color;
      item.appendChild(content);
      item.className = "danmaku-item";
      //滚动弹幕
      if (type == 0) {
        //设置轨道
        item.style.top = String(getRowTunnel(text) * 26) + "px";
        item.classList.add("danmaku-row");
        item.style.transform = `translateX(-${width}px)`;
        danmaku.value.push(item);
        danmakuRef.value.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(item);
          danmakuRef.value.removeChild(item);
        });
        if (send) {
          item.style.border = "1px solid red";
        }
        item.classList.add("danmaku-row-move");
      } else if (type == 1) {
        item.style.width = "100%";
        item.style.textAlign = "center";
        item.style.top = String(getFixedTunnel(1) * 26) + "px";
        danmaku.value.push(item);
        danmakuRef.value.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(item);
          danmakuRef.value.removeChild(item);
        });
        item.classList.add("danmaku-center-move");
      } else if (type == 2) {
        item.style.width = "100%";
        item.style.textAlign = "center";
        item.style.bottom = String(getFixedTunnel(2) * 26) + "px";
        danmaku.value.push(item);
        danmakuRef.value.appendChild(item);
        item.addEventListener("animationend", () => {
          danmaku.value.splice(item);
          danmakuRef.value.removeChild(item);
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
};
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
    transform: translateX(100%);
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

