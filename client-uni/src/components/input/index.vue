<template>
  <view class="form-item">
    <text class="label">{{ props.label }}</text>
    <input
      class="input text-right"
      :type="props.type"
      :value="props.modelValue"
      :placeholder="props.placeholder"
      @input="handleInput"
    />
  </view>
</template>

<script lang="ts" setup>
import { InputProps } from './types';

const props = defineProps<InputProps>();
const emit = defineEmits(['update:modelValue']);
const handleInput = (e: any) => {
  let val = e.detail.value;

  if (props.type === 'number') {
    val = parseFloat(val);
  }
  emit('update:modelValue', val);
};
</script>

<style lang="scss" scoped>
.form-item {
  display: flex;
  align-items: center;
  border-bottom: 1px solid $uni-border-color;
  margin-bottom: 20px;
  padding: 8px 0px;
  .label {
    margin-right: 12px;
  }
  .input {
    flex: 1;
    font-size: 14px;
  }
}
</style>
