import Axios from 'axios';

export const backendApi = Axios.create({
    baseURL: process.env.NEXT_PUBLIC_BACKEND_API_URL,
    withCredentials: true,
    headers: {
        Accept: 'application/json, application/pdf, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, text/plain',
        'Content-Type': 'application/json',
    },
});
