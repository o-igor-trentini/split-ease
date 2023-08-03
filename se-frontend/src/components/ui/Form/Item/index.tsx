'use client';

import { FC, ReactElement } from 'react';
import { Form as AntdForm, FormItemProps as AntdFormItemProps } from 'antd';
import { Rule as AntdRule, RuleRender as AntdRuleRender } from 'rc-field-form/lib/interface';

export type FormItemRule = AntdRule;
export type FormItemRuleRender = AntdRuleRender;
export type FormItemProps = AntdFormItemProps;

export const FormItem: FC<FormItemProps> = (props): ReactElement => {
    return <AntdForm.Item {...props}>{props.children}</AntdForm.Item>;
};
