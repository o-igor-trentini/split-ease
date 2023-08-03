export interface ThemeType {
    // antd
    colorPrimary: string;
    borderRadius: number;

    // antd breakpoints em 'px' (https://ant.design/components/grid#col)
    // menor que
    xs: 576;
    // maior que
    sm: 576;
    md: 768;
    lg: 992;
    xl: 1200;
    xxl: 1600;

    // app
    defaultBgColor: string;
    white: '#FFF';
}
