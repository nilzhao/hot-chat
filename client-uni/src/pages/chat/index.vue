<!-- 单聊 -->
<script lang="ts" setup>
import Avatar from '@/components/avatar/index.vue';
import useAuthStore from '@/stores/auth';
import useWsStore from '@/stores/ws';
import { Message, MessageCmdEnum, MessageMediaEnum } from '@/types/chat';
import { nextTick, onMounted, ref, watch } from 'vue';
import { trim } from 'lodash';
import { onLoad } from '@dcloudio/uni-app';
import { User } from '@/types/user';
import useContactStore from '@/stores/contact';
import CacheChat from '@/utils/cache-chat';
import GoBack from '@/components/go-back/index.vue';
import classNames from 'classnames';
import { updateFile } from '@/services/upload';
import { getArrayValue } from '@/utils';

const authStore = useAuthStore();
const contactStore = useContactStore();
const wsStore = useWsStore();

const messageList = ref<Message[]>([]);
const content = ref('');
const targetInfo = ref<User | null>(null);
const scrollTop = ref(0);

onLoad((option) => {
  const redirectToHome = () => {
    uni.switchTab({
      url: '/pages/index/index',
    });
  };
  try {
    const targetParam = JSON.parse(decodeURIComponent(option.targetInfo || ''));
    if (!targetParam) {
      redirectToHome();
      return;
    }
    targetInfo.value = targetParam;
  } catch (err) {
    redirectToHome();
  }
});

onMounted(async () => {
  // 从缓存中获取聊天记录
  const cacheChat = new CacheChat(
    authStore.currentUser,
    contactStore.contactsMap
  );
  messageList.value = await cacheChat.getTargetMessageList(
    MessageCmdEnum.SINGLE,
    targetInfo.value!.id
  );
});

const scrollToBottom = () => {
  nextTick(() => {
    const query = uni.createSelectorQuery();
    query.select('#message-scroll-view').boundingClientRect();
    query.select('#message-list-view').boundingClientRect();
    query.exec((ret) => {
      if (ret[1].height > ret[0].height) {
        scrollTop.value = ret[1].height - ret[0].height;
      }
    });
  });
};

watch(
  () => wsStore.newMsg,
  (newMsg) => {
    if (!newMsg) return;

    messageList.value.push({ ...newMsg });
    wsStore.newMsg = null;
  }
);

watch(
  () => messageList.value,
  () => {
    scrollToBottom();
  },
  {
    deep: true,
  }
);

const generateMessageDomId = (messageData: Message | undefined) => {
  if (!messageData) return '';
  return `message-${messageData.id}`;
};

const send = () => {
  const val = trim(content.value);
  if (!val) return;

  wsStore.sendText({
    targetId: targetInfo.value!.id,
    content: val,
    cmd: MessageCmdEnum.SINGLE,
  });
  content.value = '';
};

const handleAttachmentClick = () => {
  uni.chooseImage({
    count: 1,
    success: async (res) => {
      console.log(res);

      const { data } = await updateFile(getArrayValue(res.tempFilePaths));
      if (data) {
        wsStore.sendImg({
          targetId: targetInfo.value!.id,
          content: data.path,
          cmd: MessageCmdEnum.SINGLE,
          width: data.width,
          height: data.height,
        });
      }
    },
  });
};
</script>

<template>
  <view class="wrapper">
    <view class="header">
      <GoBack />
      <view>
        {{ targetInfo?.name }}
      </view>
    </view>

    <view class="body">
      <scroll-view
        scroll-y
        scroll-with-animation
        scroll-anchoring
        class="list"
        id="message-scroll-view"
        :scroll-top="scrollTop"
      >
        <view id="message-list-view">
          <template v-for="messageData in messageList" :key="messageData.id">
            <view
              :id="generateMessageDomId(messageData)"
              :class="[
                'message',
                messageData.userId === authStore.currentUser.id &&
                  'message-self',
              ]"
            >
              <view class="message-container">
                <view class="message-avatar">
                  <Avatar
                    :width="40"
                    text="她"
                    :src="
                      messageData.userId === authStore.currentUser.id
                        ? authStore.currentUser.avatar
                        : targetInfo?.avatar
                    "
                  />
                </view>
                <view class="message-card">
                  <view class="message-blank" />
                  <view class="message-content">
                    <image
                      v-if="messageData.media === MessageMediaEnum.IMG"
                      :src="`/api/v1/${messageData.content}`"
                      :style="`width: 200px;`"
                      mode="widthFix"
                    />
                    <text v-else>
                      {{ messageData.content }}
                    </text>
                  </view>
                </view>
              </view>
            </view>
          </template>
        </view>
      </scroll-view>
    </view>

    <view class="footer">
      <view
        class="item iconfont icon-attachment"
        @click="handleAttachmentClick"
      />
      <view class="item iconfont icon-emoji" />
      <textarea
        auto-height
        class="input mx-xss"
        placeholder="文明交流"
        v-model="content"
        @comfirm="send"
      />
      <view class="item iconfont icon-voice" />
      <view class="item send-more">
        <view
          :class="
            classNames('iconfont icon-send transition-fast text-primary', {
              'fade slide-right': !content,
            })
          "
          @click="send"
        />
        <view
          :class="
            classNames('iconfont icon-plus transition-fast', {
              'fade slide-left': !!content,
            })
          "
        />
      </view>
    </view>
  </view>
</template>

<style lang="scss">
@use '~@/styles/_var.scss' as *;

.wrapper {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  height: 40px;
  padding: 10px;
  text-align: center;
}

.body {
  position: relative;
  flex: 1;
  .list {
    position: absolute;
    left: 0px;
    top: 0px;
    right: 0px;
    bottom: 0px;
  }
  #message-list-view {
    padding: 20px 10px;
  }
}

$prefix: message;

.#{$prefix} {
  flex: 1;
  display: flex;
  justify-content: flex-start;
  margin-bottom: $space-md;
}
.#{$prefix}-container {
  display: flex;
  max-width: 75%;
}
.#{$prefix}-avatar {
  $size: 40px;
  width: $size;
  min-width: $size;
  height: $size;
  border-radius: 100%;
  background: grey;
  position: relative;
  z-index: 10;
  display: flex;
  overflow: hidden;
}
.#{$prefix}-card {
  display: flex;
  flex-wrap: nowrap;
}
.#{$prefix}-blank {
  $width: 16px;

  position: relative;
  width: $width;
  background: #fe2042;
  z-index: 10;
  &::after {
    content: '';
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background: #fff;
    border-radius: 0 $width 0 0;
  }
}
.#{$prefix}-content {
  flex: 1;
  background: transparent;
  color: #fff;
  padding: 10px 12px;
  line-height: 20px;
  border-radius: 0 $uni-border-radius-base $uni-border-radius-base
    $uni-border-radius-base;
  background: linear-gradient(291deg, #ff9509 0%, #fe2042 100%);
  word-break: break-all;
}

.#{$prefix}-self {
  justify-content: flex-end;
  .#{$prefix}-container {
    flex-direction: row-reverse;
  }
  .#{$prefix}-card {
    flex-direction: row-reverse;
  }
  .#{$prefix}-blank {
    $width: 16px;
    background: #e0e2e7;

    &::after {
      border-radius: $width 0 0 0;
    }
  }
  .#{$prefix}-content {
    color: #0d0e15;
    border-radius: $uni-border-radius-base 0 $uni-border-radius-base
      $uni-border-radius-base;
    background: #e0e2e7;
  }
}

.footer {
  $height: 36px;
  display: flex;
  flex-wrap: nowrap;
  align-items: flex-start;
  padding: $space-xs $space-sm;
  background: #ffffff;
  box-shadow: 0px 0px 20px 0px rgba(53, 73, 93, 0.2);
  margin: 0 -$space-xss;
  .item {
    height: $height;
    line-height: $height;
    color: #9da4b3;
    font-size: 20px;
    margin: 0 $space-xss;
  }
  .input {
    flex: 1;
    color: #0d0e15;
    line-height: $height - $space-xs * 2;
    max-height: 100px;
    overflow: auto;
    font-size: 14px;
    border-radius: $uni-border-radius-lg;
    background: #f3f4f6;
    padding: $space-xs;
    .uni-textarea-textarea {
      height: 100%;
    }
  }
  .send-more {
    position: relative;
    width: 16px;
    & > view {
      position: absolute;
      left: 0;
      top: 0;
    }
  }
}
</style>
