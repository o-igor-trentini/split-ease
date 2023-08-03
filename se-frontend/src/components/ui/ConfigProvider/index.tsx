'use client';

import { FC, ReactElement, ReactNode } from 'react';
import { ConfigProvider, theme as AntdTheme } from 'antd';
import { GlobalStyle } from '@/components/style';
import { ThemeType } from '@/components/style/types';
import { ThemeProvider } from 'styled-components';
import ptBRLocale from 'antd/lib/locale/pt_BR';
import { ThemeVariant } from '@/components/ui/ConfigProvider/types';

interface LayoutProviderPros {
    children: ReactNode;
}

const themeValues: ThemeType = {
    // antd
    colorPrimary: '#F2CD60',
    borderRadius: 6,

    // antd breakpoints
    xs: 576,
    sm: 576,
    md: 768,
    lg: 992,
    xl: 1200,
    xxl: 1600,

    // app
    defaultBgColor: '#1A1A1A',
    white: '#FFF',
};

export const AppTheme = AntdTheme;
const { defaultAlgorithm, darkAlgorithm } = AppTheme;

export const LayoutProvider: FC<LayoutProviderPros> = ({ children }): ReactElement => {
    const algorithm: Record<ThemeVariant, typeof defaultAlgorithm> = {
        light: defaultAlgorithm,
        dark: darkAlgorithm,
    };

    return (
        <ThemeProvider theme={themeValues}>
            <GlobalStyle {...themeValues} />

            <ConfigProvider
                locale={ptBRLocale}
                theme={{
                    token: { ...themeValues },
                    algorithm: algorithm['light'],
                    components: {
                        Button: {
                            colorLink: themeValues.colorPrimary,
                            colorLinkHover: themeValues.colorPrimary + '90',
                            colorLinkActive: themeValues.colorPrimary,
                        },
                    },
                }}
            >
                {children}
            </ConfigProvider>
        </ThemeProvider>
    );
};
