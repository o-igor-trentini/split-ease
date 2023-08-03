import { ReactElement, ReactNode } from 'react';
import { Row } from '@/components/ui/Row';
import { Col } from '@/components/ui/Col';

interface RecoverPasswordLayoutProps {
    children: ReactNode;
}

const RecoverPasswordLayout = ({ children }: RecoverPasswordLayoutProps): ReactElement => {
    return (
        <Row justify="center" align="middle" style={{ height: '100vh' }}>
            <Col xs={22} lg={12} xl={8} xxl={6}>
                {children}
            </Col>
        </Row>
    );
};

export default RecoverPasswordLayout;
