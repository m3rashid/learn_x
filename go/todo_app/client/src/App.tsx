import React from "react";
import { Box, List } from "@mantine/core";
import useSWR from "swr";

import AddTodo from "./addTodo";
import TodoItem from "./todoItem";

export interface ITodo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

export const ENDPOINT = "http://localhost:5000";
const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

const App = () => {
  const { data, mutate } = useSWR<ITodo[]>("api/todos", fetcher);
  return (
    <Box
      sx={() => ({
        padding: "2rem",
        width: "100%",
        maxWidth: "40rem",
        margin: "0 auto",
      })}
    >
      <List mb={20}>
        {data?.map((todo) => {
          return (
            <TodoItem
              key={`todo-list__${todo.id}`}
              mutate={mutate}
              todo={todo}
            />
          );
        })}
      </List>
      <AddTodo mutate={mutate} />
    </Box>
  );
};

export default App;
