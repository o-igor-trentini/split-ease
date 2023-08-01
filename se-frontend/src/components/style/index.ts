import { createGlobalStyle } from 'styled-components';
import { ThemeType } from '@/components/style/types';

export const GlobalStyle = createGlobalStyle<ThemeType>`
  html, body {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

    background-color: ${({ defaultBgColor }) => defaultBgColor};
  }

  ::selection {
    color: ${({ colorPrimary }) => colorPrimary};
  }
`;
