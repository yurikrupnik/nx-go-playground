// import logo from './logo.svg';
import styles from './App.module.css';
import { Router, Route, Routes, Link } from 'solid-app-router';

const Predictions = () => {
  return <div>
    Predictions
  </div>
}

const About = () => {
  return <div>
    About
  </div>
}

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
      <header class={styles.header}>
        {/*<img src={logo} class={styles.logo} alt="logo" />*/}
        <div>
          hello here
        </div>
      </header>
    </div>
  );
}

export default App;
