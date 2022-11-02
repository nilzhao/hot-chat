<script setup lang="ts">
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app';
import './styles/index.scss';
import useAuthStore from './stores/auth';
import useWsStore from './stores/ws';
import { watch } from 'vue';

onLaunch(() => {
  const authStore = useAuthStore();
  const wsStore = useWsStore();

  authStore.getCurrentUser();

  watch(
    () => authStore.isLogin,
    (isLogin) => {
      if (!isLogin) return;
      wsStore.init();
    }
  );

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
