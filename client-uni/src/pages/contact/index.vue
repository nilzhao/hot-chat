<script lang="ts" setup>
import { reqContacts } from '@/services/contact';
import { ContactDetail } from '@/types/contact';
import { groupBy, map, orderBy, pick, sortBy } from 'lodash';
import { computed, onMounted, ref } from 'vue';
import { pinyin } from 'pinyin-pro';

const contactList = ref<ContactDetail[]>([]);

const getContacts = async () => {
  const { ok, data } = await reqContacts();

  if (!ok) {
    contactList.value = [];
    return;
  }
  console.log(data);

  contactList.value = data;
};

onMounted(() => {
  getContacts();
});

const contactData = computed(() => {
  const groupedContacts = groupBy(contactList.value, (contact) => {
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
    ([letter, contacts]) => {
      return {
        letter,
        data: contacts.map((contact) => contact.targetUser.name),
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
