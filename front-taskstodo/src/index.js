import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import ErrorPage from './pages/ErrorPage';
import Login from './pages/Login';
import Tasks from './pages/Tasks';
import Task from './pages/Task';


const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Tasks /> },


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
