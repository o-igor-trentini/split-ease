import { FC, ReactElement, ReactNode } from 'react';
import { Form as AntdForm, FormProps as AntdFormProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';

export const useForm = AntdForm.useForm;

export interface FormProps {
    children?: ReactNode;
    id: string;
    form: AntdFormProps['form'];
    defaultValue?: AntdFormProps['defaultValue'];
    scrollToFirstError?: AntdFormProps['scrollToFirstError'];
    hidden?: AntdFormProps['hidden'];
    onFinish?: AntdFormProps['onFinish'];
    onChange?: AntdFormProps['onChange'];
    style?: AntdFormProps['style'];
}

export const Form: FC<FormProps> = ({
    children,
    id,
    form,
    defaultValue,
    scrollToFirstError,
    hidden,
    onFinish,
    onChange,
    style,
}): ReactElement => {
    return (
        <AntdForm
            id={makeComponentId('form', id)}
            test-id={makeTestId('form', id)}
            form={form}
            defaultValue={defaultValue}
            layout="vertical"
            scrollToFirstError={scrollToFirstError}
            hidden={hidden}
            onFinish={onFinish}
            onChange={onChange}
            style={style}
        >
            {children}
        </AntdForm>
    );
};
