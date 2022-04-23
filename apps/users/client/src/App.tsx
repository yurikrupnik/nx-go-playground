// import logo from './logo.svg';
import styles from './App.module.css';
import axios from 'axios';
import { Route, Routes, Link } from 'solid-app-router';
import { createEffect } from 'solid-js';

function getUsers() {
  return axios.get('http://localhost:3333/api/users').then((r) => r.data);
}

const Predictions = () => {
  createEffect(() => {
    getUsers().then((res) => {
      console.log('res', res);
    });
  });
  return <div>Predictions</div>;
};

const About = () => {
  return <div>About</div>;
};

function App() {
  return (
    <div class={styles.App}>
      <h1 class="underline font-bold">Welcome users-client</h1>
      <nav>
        <Link href="/about">About</Link>
        <Link href="/">Home</Link>
      </nav>
      <Routes>
        <Route path="/" element={<Predictions />} />
        <Route path="/about" element={<About />} />
        {/*<Route path="/*all" element={<NotFound />} />*/}
      </Routes>
    </div>
  );
}

export default App;
