import { AxiosError } from 'axios';

export const isAxiosError = (err: unknown): err is AxiosError => (err as AxiosError).response !== undefined;
