import { ReactNode } from 'react';
import { message as antdMessage } from 'antd';

export type ToastType = 'success' | 'error' | 'info' | 'warning';

export interface toastProps {
    content: ReactNode;
    type?: ToastType;
}

export const toast = ({ content, type = 'info' }: toastProps): void => {
    const defaultProps: Record<ToastType, { duration?: number }> = {
        info: {},
        error: { duration: 4 },
        warning: {},
        success: {},
    };

    antdMessage[type]({ ...defaultProps, content });
};
