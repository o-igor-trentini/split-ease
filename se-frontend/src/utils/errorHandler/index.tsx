import { FieldData, FormInstance } from '@/components/ui/Form';
import { isAxiosError } from '@/utils/errorHandler/utils';
import { toast } from '@/components/ui/toast';
import * as Sentry from '@sentry/browser';
import { capitalizeFirstLetter } from '@/utils/string';

interface Cause {
    field: string;
    message: string;
}

interface BackendApiError {
    code: number;
    message: string;
    error: string;
    causes: Cause[] | null;
}

export interface errorHandlerProps<T> {
    message: string;
    exception?: unknown;
    form?: FormInstance<T>;
}

export const errorHandler = <T,>({ message, exception, form }: errorHandlerProps<T>): void => {
    Sentry.captureException(exception);

    if (!isAxiosError(exception) || !exception?.response) {
        toast({ type: 'error', content: message });
        return;
    }

    const data: BackendApiError = exception.response.data as BackendApiError;

    if (data.code === 422) {
        const fields: FieldData[] = [];
        if (data?.causes)
            for (const item of data.causes)
                fields.push({
                    name: item.field,
                    errors: [capitalizeFirstLetter(item.message.replace(item.field, '').trim())],
                });
        if (form) form.setFields(fields);

        toast({ type: 'error', content: data.message });
        return;
    }

    toast({ type: 'error', content: message });
};
