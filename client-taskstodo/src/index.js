import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import App from './App';
import Home from './pages/Home';
import ErrorPage from './pages/ErrorPage';
import Login from './pages/Login';
import TaskList from './pages/TasksList';
import Task from './pages/Task';


// this router is passing for the router provider <RouterProvider/>
const router = createBrowserRouter([
  {
    path: "/",
    //and this router has a default element <App/> 
    element: <App />, //inside <App/> we have a <Outlet/> component
    errorElement: <ErrorPage />,
    children: [ // anything that is inside here
      // will use <App/> by default
      { index: true, element: <Home /> },
      {
        path: "/login",
        element: <Login />
      },
      {
        path: "/tasks",
        element: <TaskList />
      },
      {
        path: "/tasks/:id",
        element: <Task />
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
