import {
  Button,
  Card,
  Group,
  List,
  Text,
  ThemeIcon,
  useMantineTheme,
} from "@mantine/core";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import React from "react";
import { KeyedMutator } from "swr";

import { ITodo, ENDPOINT } from "./App";

interface IProps {
  todo: ITodo;
  mutate: KeyedMutator<ITodo[]>;
}

const TodoItem: React.FC<IProps> = ({ todo, mutate }) => {
  const theme = useMantineTheme();
  const secondaryColor =
    theme.colorScheme === "dark" ? theme.colors.dark[1] : theme.colors.gray[7];

  const markAsDone = async (id: number) => {
    const updated = await fetch(`${ENDPOINT}/api/todos/${id}/done`, {
      method: "PATCH",
    }).then((response) => response.json());

    mutate(updated);
  };
  const { id, title, body, done } = todo;

  return (
    <List.Item
      mb={10}
      key={`todo__${id}`}
      sx={() => ({
        listStyleType: "none",
      })}
    >
      <Card
        shadow="sm"
        p="lg"
        sx={() => ({
          width: "100%",
        })}
      >
        <Text weight={500}>{title}</Text>
        <div
          style={{
            float: "right",
            display: "flex",
          }}
        >
          <Group
            position="apart"
            style={{ marginBottom: 5, marginTop: theme.spacing.sm }}
          >
            <ThemeIcon color={done ? "green" : "red"} size={24} radius="xl">
              <CheckCircleFillIcon size={20} />
            </ThemeIcon>
            <Button onClick={() => markAsDone(id)}>
              {`Mark as ${done ? "Not Done" : "Done"}`}
            </Button>
          </Group>
        </div>
        <Text size="sm" style={{ color: secondaryColor, lineHeight: 1.5 }}>
          {body}
        </Text>
      </Card>
    </List.Item>
  );
};

export default TodoItem;
