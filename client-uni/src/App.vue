<script setup lang="ts">
import { onLaunch, onShow, onHide } from "@dcloudio/uni-app";
import "./style/index.scss";
import { useAuthStore } from "./stores/auth";

onLaunch(() => {
  const authStore = useAuthStore();
  authStore.getCurrentUser();
});

onShow(() => {
  console.log("App Show");
});

onHide(() => {
  console.log("App Hide");
});
</script>
