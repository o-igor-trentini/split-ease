import { FC, ReactElement, ReactNode } from 'react';
import { Card as AntdCard, CardProps as AntdCardProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';

export interface CardProps {
    children?: ReactNode;
    id?: string;
    hoverable?: AntdCardProps['hoverable'];
    bordered?: AntdCardProps['bordered'];
    size?: AntdCardProps['size'];
    title?: AntdCardProps['title'];
    tabProps?: AntdCardProps['tabProps'];
    tabList?: AntdCardProps['tabList'];
    hidden?: AntdCardProps['hidden'];
    cover?: AntdCardProps['cover'];
    onClick?: AntdCardProps['onClick'];
    bodyStyle?: AntdCardProps['bodyStyle'];
    style?: AntdCardProps['style'];
}

export const Card: FC<CardProps> = ({
    children,
    id,
    hoverable,
    bordered,
    size,
    title,
    tabProps,
    tabList,
    hidden,
    cover,
    onClick,
    bodyStyle,
    style,
}): ReactElement => {
    return (
        <AntdCard
            id={makeComponentId('card', id)}
            test-id={makeTestId('card', id)}
            hoverable={hoverable}
            bordered={bordered}
            size={size}
            title={title}
            tabProps={tabProps}
            tabList={tabList}
            hidden={hidden}
            cover={cover}
            onClick={onClick}
            bodyStyle={bodyStyle}
            style={style}
        >
            {children}
        </AntdCard>
    );
};
