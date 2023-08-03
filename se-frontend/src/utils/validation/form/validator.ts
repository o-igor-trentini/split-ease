import { FormItemRuleRender } from '@/components/ui/Form/Item';

interface ValidatorsType {
    password: FormItemRuleRender;
}

export const Validators: ValidatorsType = {
    password: ({ getFieldValue }) => ({
        validator(_, value) {
            if (!value || getFieldValue('password') === value) return Promise.resolve();

            return Promise.reject('As senhas não correspondem!');
        },
    }),
};
