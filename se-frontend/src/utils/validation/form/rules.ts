import { FormItemRule } from '@/components/ui/Form/Item';

interface RulesType {
    required: FormItemRule;
    password: FormItemRule[];
    username: FormItemRule[];
}

export const Rules: RulesType = {
    required: { required: true },
    password: [
        {
            min: 8,
            message: 'Mín. 8 caracteres',
        },
        {
            max: 100,
            message: 'Máx. 100 caracteres',
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
    username: [
        {
            max: 50,
            message: 'Máx. 50 caracteres',
        },
        {
            pattern: new RegExp(/^[A-Za-z0-9\-_.]+$/),
            message: 'Aceita apenas letras maiúsculas e minúsculas, números e os caracteres "-", "_", e "."!',
        },
    ],
};
