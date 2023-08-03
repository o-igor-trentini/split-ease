'use client';

import { FC, ReactElement } from 'react';
import { Input as AntdInput, InputProps as AntdInputProps } from 'antd';

export interface InputProps {
    id: string;
    type?: AntdInputProps['type'];
    value?: AntdInputProps['value'];
    defaultValue?: AntdInputProps['defaultValue'];
    placeholder?: AntdInputProps['placeholder'];
    addonAfter?: AntdInputProps['addonAfter'];
    addonBefore?: AntdInputProps['addonBefore'];
    min?: AntdInputProps['min'];
    max?: AntdInputProps['max'];
    minLength?: AntdInputProps['minLength'];
    maxLength?: AntdInputProps['maxLength'];
    status?: AntdInputProps['status'];
    disabled?: AntdInputProps['disabled'];
    hidden?: AntdInputProps['hidden'];
    bordered?: AntdInputProps['bordered'];
    onBlur?: AntdInputProps['onBlur'];
    onFocus?: AntdInputProps['onFocus'];
    onChange?: AntdInputProps['onChange'];
    style?: AntdInputProps['style'];
}

export const Input: FC<InputProps> = (props): ReactElement => {
    const p: AntdInputProps = { ...props, autoComplete: 'off' };

    const Component = p.type === 'password' ? AntdInput.Password : AntdInput;

    return <Component test-id={p.id} {...p} />;
};
