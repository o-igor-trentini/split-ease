import { FormItemRule } from '@/components/ui/Form/Item';

interface RulesType {
    required: FormItemRule;
    password: FormItemRule[];
}

export const Rules: RulesType = {
    required: { required: true },
    password: [
        {
            min: 8,
            message: 'Mín. 8 caracteres',
        },
        {
            pattern: new RegExp(/^(?=.*[!@#$%^&*])/),
            message: 'Mín. 1 caractere especial',
        },
        {
            pattern: new RegExp(/^(?=.*[a-z])/),
            message: 'Mín. 1 letra minúscula',
        },
        {
            pattern: new RegExp(/^(?=.*[A-Z])/),
            message: 'Mín. 1 letra maiúscula',
        },
        {
            pattern: new RegExp(/^(?=.*[0-9])/),
            message: 'Mín. 1 número',
        },
    ],
};
