import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { ReactElement, ReactNode } from 'react';
import { StyledComponentsRegistry } from '@/lib/registry';
import { LayoutProvider } from '@/components/ui/ConfigProvider';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
    title: 'Split Ease',
    description: 'Dividindo de forma simples e eficiente.',
};

interface RootLayoutProps {
    children: ReactNode;
}

const RootLayout = ({ children }: RootLayoutProps): ReactElement => {
    return (
        <html lang="pt-BR">
            <body className={inter.className}>
                <StyledComponentsRegistry>
                    <LayoutProvider>{children}</LayoutProvider>
                </StyledComponentsRegistry>
            </body>
        </html>
    );
};

export default RootLayout;
