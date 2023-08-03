import { FC, ReactElement } from 'react';
import { SpinningIcon } from '@/components/ui/LoadingIcon/styles';

export const LoadingIcon: FC = (): ReactElement => {
    return <SpinningIcon size={60} width="100%" />;
};
