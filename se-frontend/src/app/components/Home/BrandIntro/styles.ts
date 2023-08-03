import styled from 'styled-components';
import { Subtract } from '@phosphor-icons/react';
import { Title, Text } from '@/components/ui/Text';

export const BrandIconContainer = styled.div`
    width: max-content;
    max-width: 100%;
    background: ${({ theme }) => theme.colorPrimary};
    border-radius: 50% 90% 70% 100% / 50% 100% 70% 90%;
`;

export const BrandIcon = styled(Subtract)`
    color: ${({ theme }) => theme.defaultBgColor};
`;

export const Container = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
`;

export const BrandTitle = styled(Title)`
    font-size: 4rem !important;
`;

export const BrandText = styled(Text)`
    font-size: 1.3rem;
`;
