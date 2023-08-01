'use client';

import { FC, ReactElement, ReactNode } from 'react';
import { ConfigProvider } from 'antd';
import { GlobalStyle } from '@/components/style';
import { ThemeType } from '@/components/style/types';
import { ThemeProvider } from 'styled-components';

interface LayoutProviderPros {
    children: ReactNode;
}

const themeValues: ThemeType = {
    colorPrimary: '#F2CD60',
    defaultBgColor: '#1A1A1A',
};

export const LayoutProvider: FC<LayoutProviderPros> = ({ children }): ReactElement => {
    return (
        <ThemeProvider theme={themeValues}>
            <GlobalStyle {...themeValues} />

            <ConfigProvider theme={{ token: { ...themeValues } }}>{children}</ConfigProvider>
        </ThemeProvider>
    );
};
