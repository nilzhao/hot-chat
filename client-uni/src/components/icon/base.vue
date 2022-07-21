<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    :width="size"
    :height="size"
    viewBox="0 0 1024 1024"
    role="presentation"
  >
    <g :fill="color">
      <slot />
    </g>
  </svg>
</template>
<script lang="ts" setup>
import { IconProps } from './types';

withDefaults(defineProps<IconProps>(), {
  size: '1em',
  color: 'currentColor',
});
</script>
