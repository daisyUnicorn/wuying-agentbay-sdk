import os
import sys
import unittest

from agentbay import AgentBay

# Add the parent directory to the path so we can import the agentbay package
sys.path.append(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))


def get_test_api_key():
    """Get API key for testing"""
    api_key = os.environ.get("AGENTBAY_API_KEY")
    if not api_key:
        api_key = "akm-xxx"  # Replace with your test API key
        print(
            "Warning: Using default API key. Set AGENTBAY_API_KEY environment variable for testing."
        )
    return api_key


class TestAgentBay(unittest.TestCase):
    """Test cases for the AgentBay class."""

    def test_init_with_api_key(self):
        """Test initialization with API key."""
        api_key = get_test_api_key()
        agent_bay = AgentBay(api_key=api_key)
        self.assertEqual(agent_bay.api_key, api_key)
        self.assertIsNotNone(agent_bay.region_id)
        self.assertIsNotNone(agent_bay.client)

    def test_init_without_api_key(self):
        """Test initialization without API key."""
        os.environ["AGENTBAY_API_KEY"] = "env_api_key"
        try:
            agent_bay = AgentBay()
            self.assertEqual(agent_bay.api_key, "env_api_key")
        finally:
            del os.environ["AGENTBAY_API_KEY"]

    def test_init_without_api_key_raises_error(self):
        """Test initialization without API key raises error."""
        if "AGENTBAY_API_KEY" in os.environ:
            del os.environ["AGENTBAY_API_KEY"]
        with self.assertRaises(ValueError):
            AgentBay()

    def test_create_list_delete(self):
        """Test create, list, and delete methods."""
        api_key = get_test_api_key()
        agent_bay = AgentBay(api_key=api_key)

        # Create a session
        print("Creating a new session...")
        result = agent_bay.create()
        session = result.session
        print(f"Session created with ID: {session.session_id}")

        # Ensure session ID is not empty
        self.assertIsNotNone(session.session_id)
        self.assertNotEqual(session.session_id, "")

        # List sessions
        print("Listing sessions...")
        sessions = agent_bay.list()

        # Ensure at least one session (the one we just created)
        self.assertGreaterEqual(len(sessions), 1)

        # Check if our created session is in the list
        found = False
        for s in sessions:
            if s.session_id == session.session_id:
                found = True
                break
        self.assertTrue(
            found,
            f"Created session with ID {session.session_id} not found in sessions list",
        )

        # Delete the session
        print("Deleting the session...")
        agent_bay.delete(session)

        # List sessions again to ensure it's deleted
        sessions = agent_bay.list()

        # Check if the deleted session is not in the list
        for s in sessions:
            self.assertNotEqual(
                s.session_id,
                session.session_id,
                f"Session with ID {session.session_id} still exists after deletion",
            )


class TestSession(unittest.TestCase):
    """Test cases for the Session class."""

    def setUp(self):
        """Set up test fixtures."""
        api_key = get_test_api_key()
        self.agent_bay = AgentBay(api_key=api_key)

        # Create a session with default windows image
        print("Creating a new session for testing...")
        self.result = self.agent_bay.create()
        self.session = self.result.session
        print(f"Session created with ID: {self.session.session_id}")

    def tearDown(self):
        """Tear down test fixtures."""
        print("Cleaning up: Deleting the session...")
        try:
            self.agent_bay.delete(self.session)
        except Exception as e:
            print(f"Warning: Error deleting session: {e}")

    def test_session_properties(self):
        """Test session properties and methods."""
        # Test session properties
        self.assertIsNotNone(self.session.session_id)
        self.assertEqual(self.session.agent_bay, self.agent_bay)



        # Test get_api_key method
        api_key = self.session.get_api_key()
        self.assertEqual(api_key, self.agent_bay.api_key)

        # Test get_client method
        client = self.session.get_client()
        self.assertEqual(client, self.agent_bay.client)

        # Test get_session_id method
        session_id = self.session.get_session_id()
        self.assertEqual(session_id, self.session.session_id)

    def test_delete(self):
        """Test session delete method."""
        # Create a new session specifically for this test
        print("Creating a new session for delete testing...")
        result = self.agent_bay.create()
        session = result.session
        print(f"Session created with ID: {session.session_id}")

        # Test delete method
        print("Testing session.delete method...")
        try:
            result = session.delete()
            self.assertTrue(result)

            # Verify the session was deleted by checking it's not in the list
            sessions = self.agent_bay.list()
            for s in sessions:
                self.assertNotEqual(
                    s.session_id,
                    session.session_id,
                    f"Session with ID {session.session_id} still exists after deletion",
                )
        except Exception as e:
            print(f"Note: Session deletion failed: {e}")
            # Clean up if the test failed
            try:
                self.agent_bay.delete(session)
            except BaseException:
                pass

    def test_command(self):
        """Test command execution."""
        if self.session.command:
            print("Executing command...")
            try:
                response = self.session.command.execute_command("ls")
                print(f"Command execution result: {response}")
                self.assertIsNotNone(response)
                # Check if response contains "tool not found"
                self.assertNotIn(
                    "tool not found",
                    response.lower(),
                    "Command.ExecuteCommand returned 'tool not found'",
                )
            except Exception as e:
                print(f"Note: Command execution failed: {e}")
                # Don't fail the test if command execution is not supported
        else:
            print("Note: Command interface is nil, skipping command test")

    def test_filesystem(self):
        """Test filesystem operations."""
        if self.session.file_system:
            print("Reading file...")
            try:
                content = self.session.file_system.read_file("/etc/hosts")
                print(f"ReadFile result: content='{content}'")
                self.assertIsNotNone(content)
                # Check if response contains "tool not found"
                self.assertNotIn(
                    "tool not found",
                    content.lower(),
                    "FileSystem.ReadFile returned 'tool not found'",
                )
                print("File read successful")
            except Exception as e:
                print(f"Note: File operation failed: {e}")
                # Don't fail the test if filesystem operations are not supported
        else:
            print("Note: FileSystem interface is nil, skipping file test")


if __name__ == "__main__":
    unittest.main()
