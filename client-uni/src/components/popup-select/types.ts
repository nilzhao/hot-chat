export interface SelectOption {
  label: string | number;
  value: string | number;
  [key: string]: any;
}

export interface PopupSelectProps {
  visible: boolean;
  modelValue?: string | number;
  options: SelectOption[];
}
