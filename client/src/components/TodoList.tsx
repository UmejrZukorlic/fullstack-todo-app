import { Box, Stack, Text, Button, Flex } from "@chakra-ui/react";
import { useEffect, useState } from "react";
export interface TodoListProps {
  // Define any props if needed in the future
  _id?: string;
  body?: string;
  completed?: boolean;
}

export default function TodoList() {
  const [todos, setTodos] = useState<TodoListProps[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const fetchTodos = async () => {
    setIsLoading(true);
    try {
      const response = await fetch("http://localhost:5000/api/todos");
      const data = await response.json();
      setTodos(data);
    } catch (error) {
      console.error("Error fetching todos:", error);
    } finally {
      setIsLoading(false);
    }
  };
  useEffect(() => {
    fetchTodos();
  }, []);
  const handleDelete = async (id: string | undefined) => {
    try {
      await fetch(`http://localhost:5000/api/todos/${id}`, {
        method: "DELETE",
      });
      setTodos(todos.filter((todo) => todo._id !== id));
    } catch (error) {
      console.error("Error deleting todo:", error);
    }
  };

  // Update function to toggle completion status without clearing the body field----------------
  // Don't forget to update that list is refresing after updating and adding a new task

  const handleUpdate = async (id: string | undefined) => {
    try {
      const response = await fetch(`http://localhost:5000/api/todos/${id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) throw new Error("Failed to update todo");
      fetchTodos();
    } catch (error) {
      console.error("Error updating todo:", error);
    }
  };

  return (
    <Box>
      <Text fontSize={"lg"} fontWeight={500}>
        Your Tasks
      </Text>
      {isLoading && <Text>Loading...</Text>}
      <Stack gap={4} mt={4}>
        {todos.map((todo, i) => (
          <Box key={i} p={4} bg={"gray.200"} borderRadius={5}>
            <Flex justify="space-between" align="center" gap="4">
              <Text>{todo.body}</Text>
              <Text>{todo.completed ? "Completed" : "Pending"}</Text>
              <Flex gap={2}>
                <Button
                  size="sm"
                  colorScheme="blue"
                  onClick={() => handleUpdate(todo._id)}>
                  Update
                </Button>
                <Button
                  size="sm"
                  colorScheme="red"
                  onClick={() => handleDelete(todo._id)}>
                  Delete
                </Button>
              </Flex>
            </Flex>
          </Box>
        ))}
      </Stack>
    </Box>
  );
}
