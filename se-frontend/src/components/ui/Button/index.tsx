import { FC, ReactElement, ReactNode } from 'react';
import { Button as AntdButton, ButtonProps as AntdButtonProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';

export interface ButtonProps {
    children?: ReactNode;
    id: string;
    variant?: AntdButtonProps['type'];
    type?: AntdButtonProps['htmlType'];
    hidden?: boolean;
    disabled?: boolean;
    block?: boolean;
    loading?: boolean;
    icon?: AntdButtonProps['icon'];
    size?: AntdButtonProps['size'];
    onClick?: AntdButtonProps['onClick'];
    style?: AntdButtonProps['style'];
}

export const Button: FC<ButtonProps> = ({
    children,
    id,
    variant = 'default',
    type = 'button',
    hidden,
    disabled,
    block,
    loading,
    icon,
    size,
    onClick,
    style,
}): ReactElement => {
    return (
        <AntdButton
            id={makeComponentId('btn', id)}
            test-id={makeTestId('btn', id)}
            type={variant}
            shape="default"
            htmlType={type}
            hidden={hidden}
            disabled={disabled}
            block={block}
            loading={loading}
            icon={icon}
            size={size}
            onClick={onClick}
            style={style}
        >
            {children}
        </AntdButton>
    );
};
