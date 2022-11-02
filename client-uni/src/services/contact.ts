import { ContactDetail } from '@/types/contact';
import request from '@/utils/request';

export const reqContacts = () => request.get<ContactDetail[]>('/contacts/me');
