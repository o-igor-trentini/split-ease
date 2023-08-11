import { backendApi } from '@/lib/axios';
import { UserRegistrationFormFields } from '@/components/layout/form/UserRegistrationForm';

export const createUser = async (values: UserRegistrationFormFields): Promise<void> =>
    await backendApi.post('/user', values);
