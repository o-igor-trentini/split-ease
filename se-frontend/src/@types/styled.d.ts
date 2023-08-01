import 'styled-components';
import { ThemeType } from '@/components/style/types';

declare module 'styled-components' {
    // eslint-disable-next-line
    export interface DefaultTheme extends ThemeType {}
}
