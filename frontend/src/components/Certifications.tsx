import {  IconExternalLink, IconStar } from "@tabler/icons";

import {
  Card,
  Text,
  Title,
  Badge,
  Image,
  Group,
  Table,
  Button,
} from "@mantine/core";
import { ICertification, ICourse } from "src/model/interfaces";
import { serverURL } from "../config/constants";

interface ICertificationsProps{
    certifications: ICertification[]
}

function Certifications({certifications}:ICertificationsProps) {
    return (
        <>
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
                  src={`${serverURL}/${cert.Image}`}
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
                    <tr>
                        <th>Provider</th>
                        <th>Link</th>
                        <th>Rating</th>
                    </tr>    
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
          </>
    )
}
export default Certifications;