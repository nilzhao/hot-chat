<template>
  <view
    :class="
      classNames(styles.wrapper, props.class, styles[props.layout], {
        [styles['is-focused']]: isFocused,
        [styles['is-empty']]: isEmpty,
        [styles['is-not-empty']]: !isEmpty,
      })
    "
  >
    <text :class="classNames(styles.label, props.labelClass)">
      {{ props.label }}
    </text>
    <input
      :class="classNames(styles.input, props.inputClass)"
      :type="props.type"
      :value="props.modelValue"
      :placeholder="props.placeholder"
      @input="handleInput"
      @focus="handleFocus"
      @blur="handleBlur"
    />
  </view>
</template>

<script lang="ts" setup>
import { InputProps } from './types';
import classNames from 'classnames';
import { computed, ref } from 'vue';
import styles from './index.module.scss';

const props = withDefaults(defineProps<InputProps>(), {
  layout: 'vertical',
});
const emit = defineEmits(['update:modelValue', 'focus', 'blur']);

const isFocused = ref(false);
const isEmpty = computed(() => {
  return [null, undefined, ''].includes(props.modelValue);
});

const handleInput = (e: any) => {
  let val = e.detail.value;

  if (props.type === 'number') {
    val = parseFloat(val);
  }
  emit('update:modelValue', val);
};

const handleFocus = (e: Event) => {
  isFocused.value = true;
  emit('focus', e);
};
const handleBlur = (e: Event) => {
  isFocused.value = false;
  emit('blur', e);
};
</script>
