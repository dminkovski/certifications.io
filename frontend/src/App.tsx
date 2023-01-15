
import React from 'react';
import "./App.css";

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Home from 'src/modules/home/home';
import CreateCertification from 'src/modules/create-certification/create-certification';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home/>,
  },
  {
    path:"/certifications/create",
    element: <CreateCertification/>
  }
]);

function App() {

  return (
    <RouterProvider router={router} />
    )
}


export default App;
