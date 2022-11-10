// 联系人

import { reqContacts } from '@/services/contact';
import { ContactDetail } from '@/types/contact';
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

const useContactStore = defineStore('contact', () => {
  const contacts = ref<ContactDetail[]>([]);

  const getContacts = async () => {
    const { ok, data } = await reqContacts();
    if (ok) {
      contacts.value = data || [];
    }
  };

  const contactsMap = computed(() => {
    const map = new Map<number, ContactDetail>();
    contacts.value.forEach((contact) => {
      map.set(contact.targetId, contact);
    });
    return map;
  });

  return {
    contactsMap,
    contacts,
    getContacts,
  };
});

export default useContactStore;
