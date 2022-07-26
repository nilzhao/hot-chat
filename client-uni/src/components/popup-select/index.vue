<template>
  <uni-popup ref="popupRef" type="bottom" @change="handlePopupChange">
    <view class="text-center bg-grey type-selector">
      <view
        v-for="(option, index) in props.options"
        :key="option.label"
        :class="[
          'bg-white py-sm ',
          index !== options.length - 1 ? 'border-bottom' : 'mb-sm',
        ]"
        @click="handleChange(option)"
        >{{ option.label }}</view
      >
      <view @click="hide" class="bg-white py-sm">取消</view>
    </view>
  </uni-popup>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { PopupSelectProps, SelectOption } from './types';

const props = defineProps<PopupSelectProps>();
const emit = defineEmits(['update:modelValue', 'update:visible']);

const popupRef = ref<any>(null);

watch(
  () => props.visible,
  () => {
    if (props.visible) {
      show();
    }
  }
);

const show = () => {
  popupRef.value.open();
};

const hide = () => {
  popupRef.value.close();
};

const handlePopupChange = ({ show }: { show: boolean }) => {
  if (!show) {
    emit('update:visible', false);
  }
};

const handleChange = (option: SelectOption) => {
  emit('update:modelValue', option.value);
  hide();
};
</script>
