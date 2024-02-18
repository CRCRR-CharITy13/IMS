import { Container, Title, Text, Button, Group, Space } from '@mantine/core';
import {Error404} from "tabler-icons-react";

const style = {marginTop: "18%",
               marginLeft: "25%",
               marginRight: "25%"
              }

const PageNotFound = () => {
  return (
    <div style={style}>
        <Space h="md" />
        <Space h="md" />
        <div>
          <Title order={1} align="center">404 - Page Not Found</Title>
          <Space h="md" />
          <Text c="dimmed" size="lg" ta="center">
            Page you are trying to open does not exist. You may have mistyped the address, or the
            page has been moved to another URL. If you think this is an error contact support.
          </Text>
          <Space h="md" />
          <Group position="center">
            <Button size="md" component="a" href="/">Take me back to home page</Button>
          </Group>
        </div>
    </div>
  );
}

export default PageNotFound;