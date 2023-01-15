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
   <div id="certifications">
    <h1>Certifications</h1>
          <Certifications certifications={certifications}/>
          </div>
   
  );
}
export default Home;
