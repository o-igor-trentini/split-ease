/* eslint-disable*/
namespace NodeJS {
    interface ProcessEnv {
        NEXT_PUBLIC_BACKEND_API_URL: string;
        NEXT_PUBLIC_SENTRY_DSN: string;
        NEXT_PUBLIC_SENTRY_ENVIRONMENT: 'development' | 'staging' | 'production';
    }
}
