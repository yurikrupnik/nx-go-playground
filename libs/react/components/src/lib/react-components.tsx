import styles from './react-components.module.css';

/* eslint-disable-next-line */
export interface ReactComponentsProps {}

export function ReactComponents(props: ReactComponentsProps) {
  return (
    <h1 className="font-bold text-2xl text-blue-900 underline">
      React and Tailwind with webpack!
    </h1>
  );
}

export default ReactComponents;
