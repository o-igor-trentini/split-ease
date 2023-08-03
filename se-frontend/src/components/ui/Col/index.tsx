import { CSSProperties, FC, ReactElement, ReactNode } from 'react';
import { Col as AntdCol, ColProps as AntdColProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';

export interface ColProps {
    children?: ReactNode;
    id?: string;
    span?: AntdColProps['span'];
    xs?: AntdColProps['xs'];
    sm?: AntdColProps['sm'];
    md?: AntdColProps['md'];
    lg?: AntdColProps['lg'];
    xl?: AntdColProps['xl'];
    xxl?: AntdColProps['xxl'];
    flex?: AntdColProps['flex'];
    hidden?: boolean;
    style?: CSSProperties;
}

export const Col: FC<ColProps> = ({
    children,
    id,
    span,
    xs,
    sm,
    md,
    lg,
    xl,
    xxl,
    flex,
    hidden,
    style,
}): ReactElement => {
    return (
        <AntdCol
            id={makeComponentId('col', id)}
            test-id={makeTestId('col', id)}
            span={span}
            xs={xs}
            sm={sm}
            md={md}
            lg={lg}
            xl={xl}
            xxl={xxl}
            flex={flex}
            hidden={hidden}
            style={style}
        >
            {children}
        </AntdCol>
    );
};
