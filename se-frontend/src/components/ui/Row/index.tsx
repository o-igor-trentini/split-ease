import { CSSProperties, FC, ReactElement, ReactNode } from 'react';
import { Row as AntdRow, RowProps as AntdRowProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';

export interface RowProps {
    children?: ReactNode;
    id?: string;
    gutter?: AntdRowProps['gutter'];
    justify?: AntdRowProps['justify'];
    align?: AntdRowProps['align'];
    hidden?: boolean;
    style?: CSSProperties;
}

export const Row: FC<RowProps> = ({ children, id, gutter, justify, align, hidden, style }): ReactElement => {
    return (
        <AntdRow
            id={makeComponentId('row', id)}
            test-id={makeTestId('row', id)}
            gutter={gutter}
            justify={justify}
            align={align}
            hidden={hidden}
            style={style}
        >
            {children}
        </AntdRow>
    );
};
