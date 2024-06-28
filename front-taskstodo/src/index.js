import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import ErrorPage from './components/pages/ErrorPage';
import Home from './components/pages/Home';
import Login from './components/pages/Login';
import Tasks from './components/pages/Tasks';
import Task from './components/pages/Task';


const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Home /> },

      {
        path: "/tasks",
        element: <Tasks />,
      },
      {
        path: "/tasks/:id",
        element: <Task />,
      },
      {
        path: "/login",
        element: <Login />,
      },

    ]
  }
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
