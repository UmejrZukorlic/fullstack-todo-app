import { Box, Flex, Input, Button } from "@chakra-ui/react";
import { useState } from "react";

export default function TodoForm() {
  const [task, setTask] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!task) return;
    // Add the new task
    setTask("");
  };
  const addTask = async () => {
    if (!task || task.trim() === "") return;
    try {
      const response = await fetch("http://localhost:5000/api/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ body: task }),
      });
      if (!response.ok) throw new Error("Failed to add todo");
    } catch (error) {
      console.error("Error adding todo:", error);
    }
  };

  return (
    <Box as="form" onSubmit={handleSubmit}>
      <Flex gap={2}>
        <Input
          value={task}
          onChange={(e) => setTask(e.target.value)}
          placeholder="Add a new task"
        />
        <Button type="submit" onClick={addTask}>
          Add
        </Button>
      </Flex>
    </Box>
  );
}
