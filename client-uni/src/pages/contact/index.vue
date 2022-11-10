<script lang="ts" setup>
import { groupBy, orderBy, pick, sortBy } from 'lodash';
import { computed } from 'vue';
import { pinyin } from 'pinyin-pro';
import useContactStore from '@/stores/contact';

const { contacts } = useContactStore();

const contactData = computed(() => {
  const groupedContacts = groupBy(contacts, (contact) => {
    return pinyin(contact.targetUser.name)[0];
  });
  const letterContacts = Object.entries(groupedContacts)
    .map(([letter, contacs]) => {
      return contacs.map((contact) => {
        return {
          ...contact,
          letter,
        };
      });
    })
    .flat();
  return orderBy(letterContacts, ['letter', 'targetUser.name']);
});

const options = computed(() => {
  return Object.entries(groupBy(contactData.value, 'letter')).map(
    ([letter, list]) => {
      return {
        letter,
        data: list.map((contact) => contact.targetUser.name),
      };
    }
  );
});

const handleClick = ({ item }: { item: { itemIndex: number } }) => {
  const contact = contactData.value[item.itemIndex];
  if (!contact) return;
  uni.navigateTo({
    url: `/pages/chat/index?targetInfo=${encodeURIComponent(
      JSON.stringify(pick(contact.targetUser, ['id', 'name', 'avatar']))
    )}`,
  });
};
</script>
<template>
  <uni-indexed-list :options="options" @click="handleClick" />
</template>
