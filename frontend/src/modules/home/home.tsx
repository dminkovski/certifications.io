import { IconCheck, IconExternalLink, IconStar } from "@tabler/icons";

import {
  AppShell,
  Card,
  Navbar,
  Container,
  Header,
  NavLink,
  Text,
  Title,
  Badge,
  Image,
  Group,
  Table,
  Button,
} from "@mantine/core";
import useSWR from "swr";

import axios from "axios";
import { ICertification, ICourse } from "src/model/interfaces";
import { serverURL } from "src/config/constants";
import Certifications from "src/components/Certifications";
import { useNavigate } from "react-router-dom";

const fetcher = (url: string) => axios.get(url).then((res) => res.data);

function Home() {
    let navigate = useNavigate();

  const getCertifications = () => {
    const { data, error, isLoading } = useSWR(
      `${serverURL}/api/certifications`,
      fetcher
    );

    return {
      certifications: data,
      isLoading,
      isError: error,
    };
  };

  const { certifications, isLoading, isError } = getCertifications();

  return (
    <div className="App">
      <AppShell
        padding="sm"
        navbar={
          <Navbar p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
            <NavLink
              onClick={()=>{
                navigate(`/certifications/create`)
              }}
              label="Create Certification"
              icon={<IconCheck size={16} stroke={1.5} />}
            />
          </Navbar>
        }
        header={
          <Header height={60} p="xs">
            <Text>Header</Text>
          </Header>
        }
      >
        <Container size="xs" px="xs">
          <Certifications certifications={certifications}/>
        </Container>
      </AppShell>
    </div>
  );
}
export default Home;
