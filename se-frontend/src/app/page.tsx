'use client';

import { NextPage } from 'next';
import { ReactElement } from 'react';
import { Row } from '@/components/ui/Row';
import { Col } from '@/components/ui/Col';
import { SignInForm } from '@/components/layout/form/SignInForm';
import dynamic from 'next/dynamic';
import { LoadingIcon } from '@/components/ui/LoadingIcon';
import { BrandIntro } from '@/app/components/Home/BrandIntro';
import { BrandIconContainer } from '@/app/components/Home/styles';
import { Card } from '@/components/ui/Card';

const Home: NextPage = (): ReactElement => {
    const UserRegistrationForm = dynamic(
        () =>
            import('@/components/layout/form/UserRegistrationForm').then(
                (components) => components.UserRegistrationForm,
            ),
        {
            ssr: false,
            loading: () => <LoadingIcon />,
        },
    );

    return (
        <main>
            <Row gutter={[0, 48]} justify="center" align="middle" style={{ minHeight: '100vh' }}>
                <Col
                    xs={22}
                    flex={1}
                    style={{
                        height: '100vh',
                        display: 'flex',
                        flexDirection: 'column',
                        justifyContent: 'center',
                        alignItems: 'center',
                    }}
                >
                    <BrandIconContainer>
                        <BrandIntro />
                    </BrandIconContainer>
                </Col>

                <Col xs={22} lg={12} xl={8} xxl={6}>
                    <Card
                        tabList={[
                            { key: 'Entrar', label: 'Entrar', children: <SignInForm /> },
                            { key: 'Cadastrar', label: 'Cadastrar', children: <UserRegistrationForm /> },
                        ]}
                        tabProps={{ type: 'card' }}
                        style={{ minHeight: '100vh' }}
                    />
                </Col>
            </Row>
        </main>
    );
};

export default Home;
