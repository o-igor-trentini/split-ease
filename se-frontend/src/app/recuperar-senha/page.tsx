'use client';

import { NextPage } from 'next';
import { ReactElement, useState } from 'react';
import { Card } from '@/components/ui/Card';
import { Form, useForm } from '@/components/ui/Form';
import { Row } from '@/components/ui/Row';
import { Col } from '@/components/ui/Col';
import { FormItem } from '@/components/ui/Form/Item';
import { Input } from '@/components/ui/Input';
import { Button } from '@/components/ui/Button';
import { Rules } from '@/utils/validation/form/rules';

interface RecoverPasswordFormFields {
    new: string;
    confirmation: string;
}

const RecoverPassword: NextPage = (): ReactElement => {
    const [form] = useForm();
    const [loading, setLoading] = useState<boolean>(false);

    const handleFinish = async (values: RecoverPasswordFormFields): Promise<void> => {
        try {
            setLoading(true);

            console.log('### value', values);
        } catch (err: unknown) {
            alert(err);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Card title="Recuperação de senha">
            <Form id="changePassword" form={form} onFinish={handleFinish}>
                <Row gutter={[0, 12]} justify="center" align="middle">
                    <Col span={24}>
                        <FormItem name="username" label="Usuário" rules={[Rules.required]}>
                            <Input id="username" placeholder="Usuário ou e-mail" />
                        </FormItem>
                    </Col>

                    <Col span={24}>
                        <Button id="submit" variant="primary" type="submit" block loading={loading}>
                            Enviar
                        </Button>
                    </Col>
                </Row>
            </Form>
        </Card>
    );
};

export default RecoverPassword;
