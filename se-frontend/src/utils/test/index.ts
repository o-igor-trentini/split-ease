export const makeComponentId = (prefix: string, id?: string): string | undefined =>
    id ? prefix + '-' + id : undefined;

export const makeTestId = (prefix: string, id?: string): string | undefined =>
    id ? 'test_id_' + makeComponentId(prefix, id) : undefined;
