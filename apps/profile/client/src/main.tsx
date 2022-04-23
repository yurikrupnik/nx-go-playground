import * as ReactDOMClient from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import './styles.css';
import App from './app/app';

const root = ReactDOMClient.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <BrowserRouter>
    <App />
  </BrowserRouter>
);
