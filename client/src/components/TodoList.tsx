import { Box, Stack, Text, Button, Flex } from "@chakra-ui/react";

export default function TodoList() {
  return (
    <Box>
      <Text fontSize={"lg"} fontWeight={500}>
        Your Tasks
      </Text>
      <Stack gap={4} mt={4}>
        {/* Task 1 */}
        <Box p={4} bg={"gray.200"} borderRadius={5}>
          <Flex justify="space-between" align="center" gap="4">
            <Text>Sample Task 1</Text>
            <Flex gap={2}>
              <Button size="sm" colorScheme="blue">
                Update
              </Button>
              <Button size="sm" colorScheme="red">
                Delete
              </Button>
            </Flex>
          </Flex>
        </Box>

        {/* Task 2 */}
        <Box p={4} bg={"gray.200"} borderRadius={5}>
          <Flex justify="space-between" align="center">
            <Text>Sample Task 2</Text>
            <Flex gap={2}>
              <Button size="sm" colorScheme="blue">
                Update
              </Button>
              <Button size="sm" colorScheme="red">
                Delete
              </Button>
            </Flex>
          </Flex>
        </Box>
      </Stack>
    </Box>
  );
}
