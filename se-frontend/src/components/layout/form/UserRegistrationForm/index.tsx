'use client';

import { Form, useForm } from '@/components/ui/Form';
import { Input, InputProps } from '@/components/ui/Input';
import { FormItem, FormItemProps } from '@/components/ui/Form/Item';
import { Button } from '@/components/ui/Button';
import { FC, ReactElement, useMemo, useState } from 'react';
import { Row } from '@/components/ui/Row';
import { Col } from '@/components/ui/Col';
import { Rules } from '@/utils/validation/form/rules';
import { Validators } from '@/utils/validation/form/validator';

interface UserRegistrationFormFields {
    firstName: string;
    lastName: string;
    email: string;
    user: string;
    password: string;
    passwordConfirmation: string;
}

export const UserRegistrationForm: FC = (): ReactElement => {
    const [form] = useForm<UserRegistrationFormFields>();
    const [loading, setLoading] = useState<boolean>(false);

    const handleFinish = async (values: UserRegistrationFormFields): Promise<void> => {
        try {
            setLoading(true);

            console.log('### value', values);
        } catch (err: unknown) {
            alert(err);
        } finally {
            setLoading(false);
        }
    };

    const items = useMemo<{ formItem: FormItemProps; input: InputProps }[]>(
        () => [
            {
                formItem: {
                    name: 'firstName',
                    label: 'Nome',
                    rules: [Rules.required],
                },
                input: {
                    id: 'firstName',
                    placeholder: 'Nome',
                },
            },
            {
                formItem: {
                    name: 'lastName',
                    label: 'Sobrenome',
                    rules: [Rules.required],
                },
                input: {
                    id: 'lastName',
                    placeholder: 'Sobrenome',
                },
            },
            {
                input: {
                    id: 'email',
                    placeholder: 'email@exemplo.com',
                    type: 'email',
                },
                formItem: {
                    name: 'email',
                    label: 'E-mail',
                    rules: [Rules.required],
                },
            },
            {
                input: {
                    id: 'user',
                    placeholder: 'Usuário',
                },
                formItem: {
                    name: 'user',
                    label: 'Usuário',
                    rules: [Rules.required],
                },
            },
            {
                input: {
                    id: 'password',
                    placeholder: 'Senha',
                    type: 'password',
                },
                formItem: {
                    name: 'password',
                    label: 'Senha',
                    rules: [Rules.required, ...Rules.password],
                    hasFeedback: true,
                },
            },
            {
                input: {
                    id: 'passwordConfirmation',
                    placeholder: 'Confirme a senha',
                    type: 'password',
                },
                formItem: {
                    name: 'passwordConfirmation',
                    label: 'Confirme a senha',
                    hasFeedback: true,
                    rules: [Rules.required, Validators.password],
                },
            },
        ],
        [],
    );

    return (
        <Form id="login" form={form} onFinish={handleFinish}>
            <Row gutter={[0, 12]}>
                {items.map(({ formItem, input }) => (
                    <Col key={formItem.name} span={24}>
                        <FormItem {...formItem}>
                            <Input {...input} />
                        </FormItem>
                    </Col>
                ))}

                <Col span={24}>
                    <Button id="signin" variant="primary" type="submit" block loading={loading}>
                        Cadastrar
                    </Button>
                </Col>
            </Row>
        </Form>
    );
};
