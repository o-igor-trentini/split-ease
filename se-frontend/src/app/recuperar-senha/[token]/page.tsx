'use client';

import { NextPage } from 'next';
import { ReactElement, useState } from 'react';
import { Col } from '@/components/ui/Col';
import { Card } from '@/components/ui/Card';
import { Row } from '@/components/ui/Row';
import { useParams, useRouter } from 'next/navigation';
import { Form, useForm } from '@/components/ui/Form';
import { Input } from '@/components/ui/Input';
import { FormItem } from '@/components/ui/Form/Item';
import { Button } from '@/components/ui/Button';
import { Rules } from '@/utils/validation/form/rules';
import { Validators } from '@/utils/validation/form/validator';

interface ChangePasswordFormFields {
    password: string;
    passwordConfirmation: string;
}

const ChangePassword: NextPage = (): ReactElement => {
    const params = useParams();
    const router = useRouter();
    const [form] = useForm<ChangePasswordFormFields>();
    const [loading, setLoading] = useState<boolean>(false);

    console.log('### params', params);

    const handleFinish = async (values: ChangePasswordFormFields): Promise<void> => {
        try {
            setLoading(true);

            console.log('### value', values);

            router.replace('/');
        } catch (err: unknown) {
            alert(err);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Card title="Alteração de senha">
            <Form id="changePassword" form={form} onFinish={handleFinish}>
                <Row gutter={[0, 12]} justify="center" align="middle">
                    <Col span={24}>
                        <FormItem
                            name="password"
                            label="Nova senha"
                            hasFeedback
                            rules={[Rules.required, ...Rules.password]}
                        >
                            <Input id="password" type="password" placeholder="Nova senha" maxLength={100} />
                        </FormItem>
                    </Col>

                    <Col span={24}>
                        <FormItem
                            name="passwordConfirmation"
                            label="Confirme a senha"
                            hasFeedback
                            rules={[Rules.required, Validators.password]}
                        >
                            <Input
                                id="passwordConfirmation"
                                type="password"
                                placeholder="Confirme a senha"
                                maxLength={100}
                            />
                        </FormItem>
                    </Col>

                    <Col span={24}>
                        <Button id="change" variant="primary" type="submit" block loading={loading}>
                            Alterar
                        </Button>
                    </Col>
                </Row>
            </Form>
        </Card>
    );
};

export default ChangePassword;
