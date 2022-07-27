export interface InputProps {
  modelValue?: string | number;
  class?: string;
  labelClass?: string;
  inputClass?: string;
  label?: string;
  type?: 'number' | 'text' | 'password';
  placeholder?: string;
  layout?: 'vertical' | 'horizontal' | 'internal';
}
