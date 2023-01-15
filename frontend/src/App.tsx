
import React from 'react';
import "./App.css";
import { IconCheck, IconExternalLink, IconStar } from "@tabler/icons";

import {
  AppShell,
  Navbar,
  Container,
  Header,
  NavLink,
  Text,
} from "@mantine/core";
import {
  createBrowserRouter,
  RouterProvider,
  useNavigate
} from "react-router-dom";

import Home from 'src/modules/home/home';
import CreateCertification from 'src/modules/create-certification/create-certification';
import Root from './Root';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path:"certifications/create",
        element: <CreateCertification/>
      }
    ],
  },
 
]);

function App() {

  return (
    <div className="App">
      <RouterProvider router={router} />
    </div>
    )
}


export default App;
