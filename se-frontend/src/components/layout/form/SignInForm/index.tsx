'use client';

import { Form, useForm } from '@/components/ui/Form';
import { Input } from '@/components/ui/Input';
import { FormItem } from '@/components/ui/Form/Item';
import { Button } from '@/components/ui/Button';
import { FC, ReactElement, useRef, useState } from 'react';
import { Row } from '@/components/ui/Row';
import { Col } from '@/components/ui/Col';
import Link from 'next/link';
import { Rules } from '@/utils/validation/form/rules';

interface SignInFormFields {
    username: string;
    password: string;
}

export const SignInForm: FC = (): ReactElement => {
    const [form] = useForm<SignInFormFields>();
    const [loading, setLoading] = useState<boolean>(false);
    const refLink = useRef<HTMLAnchorElement>(null);

    const handleFinish = async (values: SignInFormFields): Promise<void> => {
        try {
            setLoading(true);

            console.log('### value', values);
        } catch (err: unknown) {
            alert(err);
        } finally {
            setLoading(false);
        }
    };

    const handleRecoveryPassword = () => refLink.current?.click();

    return (
        <Form id="login" form={form} onFinish={handleFinish}>
            <Row gutter={[0, 12]}>
                <Col span={24}>
                    <FormItem name="username" label="Usuário" rules={[Rules.required]}>
                        <Input id="username" placeholder="Usuário ou e-mail" />
                    </FormItem>
                </Col>

                <Col span={24}>
                    <FormItem name="password" label="Senha" rules={[Rules.required]}>
                        <Input id="password" placeholder="Senha" type="password" maxLength={100} />
                    </FormItem>
                </Col>

                <Col span={24}>
                    <Button id="signin" variant="link" disabled={loading} onClick={handleRecoveryPassword}>
                        <Link ref={refLink} href="/recuperar-senha" prefetch>
                            Esqueci minha senha
                        </Link>
                    </Button>
                </Col>

                <Col span={24}>
                    <Button id="signin" variant="primary" type="submit" block loading={loading}>
                        Entrar
                    </Button>
                </Col>
            </Row>
        </Form>
    );
};
