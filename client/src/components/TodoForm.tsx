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

  return (
    <Box as="form" onSubmit={handleSubmit}>
      <Flex gap={2}>
        <Input
          value={task}
          onChange={(e) => setTask(e.target.value)}
          placeholder="Add a new task"
        />
        <Button type="submit">Add</Button>
      </Flex>
    </Box>
  );
}
