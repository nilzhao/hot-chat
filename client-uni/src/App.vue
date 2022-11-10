<script setup lang="ts">
import { onLaunch, onShow, onHide } from '@dcloudio/uni-app';
import './styles/index.scss';
import useAuthStore from './stores/auth';
import useWsStore from './stores/ws';
import { watch } from 'vue';
import useContactStore from './stores/contact';

onLaunch(() => {
  const authStore = useAuthStore();
  const contactStore = useContactStore();
  const wsStore = useWsStore();

  authStore.getCurrentUser();

  // 登录后 连接 websocket
  watch(
    () => authStore.isLogin,
    async (isLogin) => {
      if (!isLogin) return;
      await contactStore.getContacts();
      wsStore.init({
        currentUser: authStore.currentUser,
        contactsMap: contactStore.contactsMap,
      });
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
