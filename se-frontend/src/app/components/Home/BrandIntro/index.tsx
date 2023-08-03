import { FC, ReactElement } from 'react';
import { Row } from '@/components/ui/Row';
import {
    BrandIcon,
    BrandIconContainer,
    BrandText,
    BrandTitle,
    Container,
} from '@/app/components/Home/BrandIntro/styles';
import { Col } from '@/components/ui/Col';

export const BrandIntro: FC = (): ReactElement => {
    return (
        <Row>
            <Col span={24}>
                <Container>
                    <BrandIconContainer>
                        <BrandIcon size={180} width="100%" />
                    </BrandIconContainer>
                </Container>
            </Col>

            <Col span={24}>
                <Container>
                    <BrandTitle level={1}>Split Ease</BrandTitle>

                    <BrandText>Dividindo de forma simples e eficiente.</BrandText>
                </Container>
            </Col>
        </Row>
    );
};
