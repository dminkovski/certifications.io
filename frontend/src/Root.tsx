
import React from 'react';
import "./App.css";
import { IconCheck, IconExternalLink, IconHome, IconStar } from "@tabler/icons";

import {
  AppShell,
  Navbar,
  Container,
  Header,
  NavLink,
  Text
} from "@mantine/core";
import {
  createBrowserRouter,
  RouterProvider,
  useNavigate,
  Outlet
} from "react-router-dom";

import Home from 'src/modules/home/home';
import CreateCertification from 'src/modules/create-certification/create-certification';


function Root(){
    let navigate = useNavigate();


    return (
        <AppShell
          padding="sm"
          navbar={
            <Navbar p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
              <Navbar.Section>
                <NavLink onClick={()=>{
                  navigate(`/`)
                }}
                label="Dashboard"
                icon={<IconHome size={16} stroke={1.5} />}
                />
                <NavLink
                  onClick={()=>{
                    navigate(`/certifications/create`)
                  }}
                  label="Create Certification"
                  icon={<IconCheck size={16} stroke={1.5} />}
                />
              </Navbar.Section>
             
            </Navbar>
          }
          header={
            <Header height={60} p="xs">
              <Text>Header</Text>
            </Header>
          }
        >
          <Container size="xs" px="xs">
              <Outlet />
          </Container>
          </AppShell>
    )
}
export default Root;