import { IconCheck, IconExternalLink, IconStar } from "@tabler/icons";

import "./App.css";
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
import { ICertification, ICourse } from "./model/interfaces";

const fetcher = (url: string) => axios.get(url).then((res) => res.data);

function App() {
  const getCertifications = () => {
    const { data, error, isLoading } = useSWR(
      `http://localhost:8081/api/certifications`,
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
          {certifications?.map((cert: ICertification, certIndex: number) => (
            <Card
              key={`cert-${certIndex}`}
              shadow="sm"
              p="lg"
              radius="md"
              withBorder
            >
              <Card.Section>
                <Image
                  height={160}
                  radius="md"
                  src={`http://localhost:8081/${cert.Image}`}
                  alt={cert.Name}
                />
              </Card.Section>
              <Group position="center" mt="md" mb="xs">
                <Title order={4}>{cert.Name}</Title>
                <Badge color="blue" variant="dark">
                  223 Ratings
                </Badge>
                <Text>{cert.Notes}</Text>
              </Group>
              <Table striped mt="md" mb="xs">
                <thead>
                  <th>Provider</th>
                  <th>Link</th>
                  <th>Rating</th>
                </thead>
                <tbody>
                  {cert.Courses?.map((c: ICourse, index: number) => (
                    <tr key={`course-${index}`}>
                      <td>
                        <Text
                          component="a"
                          target="_blank"
                          href={c.Provider?.Link}
                        >
                          {c.Provider?.Name}
                        </Text>
                      </td>
                      <td>
                        <Button
                          component="a"
                          variant="light"
                          target="_blank"
                          href={c.Link}
                          color="blue"
                          fullWidth
                          mt="md"
                          leftIcon={<IconExternalLink size={14} />}
                          radius="md"
                        >
                          {c.Name}
                        </Button>
                      </td>
                      <td></td>
                    </tr>
                  ))}
                </tbody>
              </Table>
              {cert.Skills?.map((s: string, index: number) => (
                <Badge
                  key={`skill-${index}`}
                  variant="gradient"
                  sx={{ marginRight: 5, marginLeft: 5 }}
                  gradient={{ from: "indigo", to: "cyan" }}
                >
                  {s}
                </Badge>
              ))}
              <Button
                variant="gradient"
                gradient={{ from: "orange", to: "red" }}
                color="white"
                fullWidth
                mt="md"
                leftIcon={<IconStar size={14} />}
                radius="md"
              >
                Leave A Rating
              </Button>
            </Card>
          ))}
        </Container>
      </AppShell>
    </div>
  );
}

export default App;
