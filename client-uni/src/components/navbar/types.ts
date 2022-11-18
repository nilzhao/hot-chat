import { ExtractPropTypes, PropType } from 'vue';

export const navbarProps = {
  title: String,
  hideLeft: Boolean,
} as const;

export type NavbarProps = ExtractPropTypes<typeof navbarProps>;
