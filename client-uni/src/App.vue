<script setup lang="ts">
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app';
import './styles/index.scss';
import { useAuthStore } from './stores/auth';

onLaunch(() => {
  const authStore = useAuthStore();
  authStore.getCurrentUser();
  uni.getNetworkType({
    success(res) {
      authStore.network.type = res.networkType;
    },
  });
  uni.onNetworkStatusChange((res) => {
    authStore.network.type = res.networkType;
    authStore.network.isConnected = res.isConnected;
  });
});

onShow(() => {
  console.log('App Show');
});

onHide(() => {
  console.log('App Hide');
});
</script>
