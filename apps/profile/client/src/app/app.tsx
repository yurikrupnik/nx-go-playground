// eslint-disable-next-line @typescript-eslint/no-unused-vars
// import styles from './app.module.css';
import { useState, useEffect } from 'react';
import { Route, Link, Switch, Router } from 'react-router-dom';
import { User, UserDocument } from '@nx-go-playground/api/users';
import { createApis } from '@nx-go-playground/api/helpers';

import {
  useQuery,
  useMutation,
  useQueryClient,
  QueryClient,
  QueryClientProvider,
} from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // refetchOnWindowFocus: false,
      // refetchOnmount: false,
      // refetchOnReconnect: false,
      // retry: false,
      // staleTime: 30000000,
    },
  },
});

// const api = createApis<User>('http://localhost:3333/api/users');

const Com = () => {
  // const queryClient = useQueryClient();
  //
  // console.log('queryClient', queryClient)
  // Queries
  const id = '6262eec5f33e22bf8e7dd487';
  const params = {
    email: 'a@a.com',
    projections: ['email'],
    page: 0,
    limit: 0,
  };
  // const query = useQuery(
  //   [
  //     api.name,
  //     params,
  //     // {
  //     //   email: 'a@a.com',
  //     //   projections: ['email'],
  //     //   page: 0,
  //     //   limit: 0,
  //     // },
  //   ],
  //   // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  //   // @ts-ignore
  //   () => api.get(params),
  //   {
  //     // cacheTime: 200000,
  //     staleTime: 30000000,
  //     // initialData: [],
  //   }
  // );
  // console.log('query', query.data);

  // Mutations
  // const mutation = useMutation(postTodo, {
  //   onSuccess: () => {
  //     // Invalidate and refetch
  //     queryClient.invalidateQueries('todos');
  //   },
  // });

  // const [data, setData] = useState<User[]>([]);
  // useEffect(() => {
  //   getUsers(1).then((res) => {
  //     console.log('res', res);
  //     setData(res);
  //   });
  // }, []);

  return (
    <div>
      <h1>data here</h1>
      {/*{Array.isArray(query.data) &&*/}
      {/*  query.data.map((v) => {*/}
      {/*    // eslint-disable-next-line @typescript-eslint/ban-ts-comment*/}
      {/*    return (*/}
      {/*      <div key={v._id}>*/}
      {/*        <div>{v.name}</div>*/}
      {/*        <div>{v._id}</div>*/}
      {/*      </div>*/}
      {/*    );*/}
      {/*  })}*/}
    </div>
  );
};

export function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <h1 className="underline">Welcome profile-client</h1>
      <div role="navigation">
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/page-2">Page 2 extra</Link>
          </li>
        </ul>
      </div>
      <div>
        <Route
          path="/"
          exact
          render={() => (
            <div>
              This is the generated root route.{' '}
              <Link to="/page-2">Click here for page 2.</Link>
            </div>
          )}
        />
        <Route path="/page-2" exact render={() => <Com />} />
      </div>
      {/* END: routes */}
      <ReactQueryDevtools />
    </QueryClientProvider>
  );
}

export default App;
