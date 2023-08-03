import styled from 'styled-components';

export const BrandIconContainer = styled.div`
    width: fit-content;
    height: fit-content;

    padding: 3rem;

    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 26% 74% 29% 71% / 75% 68% 32% 25%;
    background: ${({ theme }) => theme.white};

    @media (max-width: ${({ theme }) => theme.xs}px) {
        border-radius: ${({ theme }) => theme.borderRadius}px;
    }
`;
