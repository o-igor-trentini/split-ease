import styled, { keyframes } from 'styled-components';
import { CircleNotch } from '@phosphor-icons/react';

const spin = keyframes`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`;

export const SpinningIcon = styled(CircleNotch)`
    animation: ${spin} 2s linear infinite;
`;
