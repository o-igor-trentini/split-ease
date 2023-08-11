import { FC, ReactElement, ReactNode } from 'react';
import { Form as AntdForm, FormInstance as AntdFormInstance, FormProps as AntdFormProps } from 'antd';
import { makeComponentId, makeTestId } from '@/utils/test';
import { FieldData as AntdFieldData } from 'rc-field-form/lib/interface';

export const useForm = AntdForm.useForm;

export type FormInstance<T> = AntdFormInstance<T>;
export type FieldData = AntdFieldData;

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
