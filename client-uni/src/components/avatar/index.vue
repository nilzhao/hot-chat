<script lang="ts">
const DEFAULT_WIDTH: number = 80;
</script>

<script lang="ts" setup>
import { computed } from 'vue';
import { AvatarProps } from './types';

const props = withDefaults(defineProps<AvatarProps>(), {
  width: DEFAULT_WIDTH,
});

const style = computed(() => {
  return {
    width: props.width + 'px',
    height: props.width + 'px',
  };
});

const draw = (text?: string): string => {
  if (!text) return '';
  const colors = [
    '#748DA6',
    '#9CB4CC',
    '#D3CEDF',
    '#F2D7D9',
    '#8879B0',
    '#FBA1A1',
    '#FBC5C5',
    '#377D71',
    '#0078AA',
    '#3AB4F2',
    '#F2DF3A',
  ];
  const cvs = document.createElement('canvas');
  cvs.setAttribute('width', props.width + '');
  cvs.setAttribute('height', props.width + '');
  const ctx = cvs.getContext('2d');
  if (!ctx) return '';
  ctx.fillStyle = colors[Math.floor(Math.random() * colors.length)];
  ctx.fillRect(0, 0, props.width, props.width);
  ctx.fillStyle = '#FEFBF6';
  ctx.font = props.width * 0.6 + 'px Arial';
  ctx.textBaseline = 'middle';
  ctx.textAlign = 'center';
  ctx.fillText(text.substring(1, 0), props.width / 2, props.width / 2);

  return cvs.toDataURL('image/jpeg', 1);
};
</script>

<template>
  <view
    :style="style"
    :class="[
      props.className,
      'avatar',
      {
        round: props.round,
      },
    ]"
  >
    <image :style="style" :src="props.src || draw(props.text)" />
  </view>
</template>

<style lang="scss" scoped>
.avatar {
  overflow: hidden;
  border-radius: $uni-border-radius-base;
  &.round {
    border-radius: 100%;
  }
}
</style>
