import last from 'lodash/last';
import axios, { AxiosResponse } from 'axios';

type Api<U> = {
  name: string;
  getById: (id: string) => Promise<U>;
  get: (
    // key: string,
    params: Partial<U> & {
      projections?: keyof U[];
      page?: number;
      limit?: number;
      // [Key in U]: any;
    }
  ) => Promise<U[]>;
  post: (body: Partial<U>) => Promise<U>;
  put: (body: Partial<U>) => Promise<U>;
  delete: (id: string) => Promise<string>;
};

function handleResArray<T>(res: AxiosResponse): Promise<T[]> {
  return res.data;
}

function handleResObject<T>(res: AxiosResponse): Promise<T> {
  return res.data;
}

function handleResId(res: AxiosResponse): Promise<string> {
  return res.data;
}

export function createApis<U>(url: string): Api<U> {
  return {
    name: last(url.split('/')) || '',
    get(
      params: Partial<U> & {
        projections?: keyof U[];
        page?: number;
        limit?: number;
      }
    ): Promise<U[]> {
      return axios.get<U>(url, { params }).then((res) => handleResArray(res));
    },
    getById(id: string): Promise<U> {
      return axios
        .get<U>(`${url}/${id}`)
        .then((res) => handleResObject<U>(res));
    },
    post(body: Partial<U>) {
      return axios.post(`${url}`, body).then((res) => handleResObject<U>(res));
    },
    put(body: Partial<U>) {
      return axios.put(`${url}`, body).then((res) => handleResObject<U>(res));
    },
    delete(id: string) {
      return axios.delete(`${url}/${id}`).then(handleResId);
    },
  };
}
