import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import ErrorPage from './pages/ErrorPage';
import Home from './pages/Home'
import TaskList from './pages/TasksList'
import Login from './pages/Login'


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
        path: "/tasks",
        element: <TaskList />
      },
      {
        path: "/login",
        element: <Login />
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
