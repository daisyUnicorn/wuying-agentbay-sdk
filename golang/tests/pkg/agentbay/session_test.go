package agentbay_test

import (
	"fmt"
	"testing"

	"github.com/aliyun/wuying-agentbay-sdk/golang/pkg/agentbay"
	"github.com/aliyun/wuying-agentbay-sdk/golang/tests/pkg/agentbay/testutil"
)

func TestSession_ValidateLabels(t *testing.T) {
	// Initialize AgentBay client
	apiKey := testutil.GetTestAPIKey(t)
	agentBay, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	// Create a session
	fmt.Println("Creating a new session for label validation testing...")
	sessionResult, err := agentBay.Create(nil)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	session := sessionResult.Session
	t.Logf("Session created with ID: %s", session.SessionID)

	defer func() {
		// Clean up the session after test
		fmt.Println("Cleaning up: Deleting the session...")
		deleteResult, err := agentBay.Delete(session)
		if err != nil {
			t.Logf("Warning: Error deleting session: %v", err)
		} else {
			t.Logf("Session deleted (RequestID: %s)", deleteResult.RequestID)
		}
	}()

	// Test successful validation with valid labels
	t.Run("ValidLabels", func(t *testing.T) {
		labels := map[string]string{"key1": "value1", "key2": "value2"}
		errMsg := session.ValidateLabels(labels)
		if errMsg != "" {
			t.Errorf("Expected successful validation, got error: %s", errMsg)
		}
	})

	// Test validation with nil labels
	t.Run("NilLabels", func(t *testing.T) {
		var labels map[string]string
		errMsg := session.ValidateLabels(labels)
		if errMsg == "" {
			t.Error("Expected validation to fail with nil labels")
		}
		if errMsg != "Labels cannot be nil. Please provide a valid labels map." {
			t.Errorf("Expected specific error message, got: %s", errMsg)
		}
	})

	// Test validation with empty map
	t.Run("EmptyLabels", func(t *testing.T) {
		labels := map[string]string{}
		errMsg := session.ValidateLabels(labels)
		if errMsg == "" {
			t.Error("Expected validation to fail with empty labels")
		}
		if errMsg != "Labels cannot be empty. Please provide at least one label." {
			t.Errorf("Expected specific error message, got: %s", errMsg)
		}
	})

	// Test validation with empty key
	t.Run("EmptyKey", func(t *testing.T) {
		labels := map[string]string{"": "value1"}
		errMsg := session.ValidateLabels(labels)
		if errMsg == "" {
			t.Error("Expected validation to fail with empty key")
		}
		if errMsg != "Label keys cannot be empty. Please provide valid keys." {
			t.Errorf("Expected specific error message, got: %s", errMsg)
		}
	})

	// Test validation with empty value
	t.Run("EmptyValue", func(t *testing.T) {
		labels := map[string]string{"key1": ""}
		errMsg := session.ValidateLabels(labels)
		if errMsg == "" {
			t.Error("Expected validation to fail with empty value")
		}
		if errMsg != "Label values cannot be empty. Please provide valid values." {
			t.Errorf("Expected specific error message, got: %s", errMsg)
		}
	})
}

func TestSession_Properties(t *testing.T) {
	// Initialize AgentBay client
	apiKey := testutil.GetTestAPIKey(t)
	agentBay, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	// Create a session
	fmt.Println("Creating a new session for session testing...")
	sessionResult, err := agentBay.Create(nil)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	// Verify RequestID from Create operation
	if sessionResult.RequestID == "" {
		t.Errorf("Create method did not return RequestID")
	} else {
		t.Logf("Create method successfully returned RequestID: %s", sessionResult.RequestID)
	}

	session := sessionResult.Session
	t.Logf("Session created with ID: %s", session.SessionID)

	defer func() {
		// Clean up the session after test
		fmt.Println("Cleaning up: Deleting the session...")
		deleteResult, err := agentBay.Delete(session)
		if err != nil {
			t.Logf("Warning: Error deleting session: %v", err)
		} else {
			// Verify RequestID from Delete operation
			if deleteResult.RequestID == "" {
				t.Errorf("Delete method did not return RequestID")
			} else {
				t.Logf("Session deleted (RequestID: %s)", deleteResult.RequestID)
			}
		}
	}()

	// Test session properties
	if session.SessionID == "" {
		t.Errorf("Expected non-empty session ID")
	}
	if session.AgentBay != agentBay {
		t.Errorf("Expected AgentBay to be the same instance")
	}

	// Test GetSessionId method
	sessionID := session.GetSessionId()
	if sessionID != session.SessionID {
		t.Errorf("Expected GetSessionId to return '%s', got '%s'", session.SessionID, sessionID)
	}

	// Test GetAPIKey method
	apiKeyFromSession := session.GetAPIKey()
	if apiKeyFromSession != apiKey {
		t.Errorf("Expected GetAPIKey to return '%s', got '%s'", apiKey, apiKeyFromSession)
	}

	// Test GetClient method
	client := session.GetClient()
	if client == nil {
		t.Errorf("Expected GetClient to return a non-nil client")
	}

	// Test Info method and verify RequestID
	infoResult, err := session.Info()
	if err != nil {
		t.Errorf("Error getting session info: %v", err)
	} else {
		if infoResult.RequestID == "" {
			t.Errorf("Info method did not return RequestID")
		} else {
			t.Logf("Info method successfully returned RequestID: %s", infoResult.RequestID)
		}
	}
}

func TestSession_DeleteMethod(t *testing.T) {
	// Initialize AgentBay client
	apiKey := testutil.GetTestAPIKey(t)
	agentBay, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	// Create a session
	fmt.Println("Creating a new session for delete testing...")
	sessionResult, err := agentBay.Create(nil)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	// Verify RequestID from Create operation
	if sessionResult.RequestID == "" {
		t.Errorf("Create method did not return RequestID")
	} else {
		t.Logf("Create method successfully returned RequestID: %s", sessionResult.RequestID)
	}

	session := sessionResult.Session
	t.Logf("Session created with ID: %s", session.SessionID)

	// Test Delete method
	fmt.Println("Testing session.Delete method...")
	deleteResult, err := session.Delete()
	if err != nil {
		t.Fatalf("Error deleting session: %v", err)
	}

	// Verify RequestID from Delete operation
	if deleteResult.RequestID == "" {
		t.Errorf("Delete method did not return RequestID")
	} else {
		t.Logf("Session deleted (RequestID: %s)", deleteResult.RequestID)
	}

	// Verify the session was deleted by trying to list sessions
	listResult, err := agentBay.List()
	if err != nil {
		t.Fatalf("Error listing sessions: %v", err)
	}

	// Note: List method is a local operation and does not return RequestID
	t.Logf("Sessions listed (local operation, no RequestID)")

	// Check if the deleted session is not in the list
	for _, s := range listResult.Sessions {
		if s.SessionID == session.SessionID {
			t.Errorf("Session with ID %s still exists after deletion", session.SessionID)
		}
	}
}

func TestSession_GetLinkMethod(t *testing.T) {
	// Initialize AgentBay client
	apiKey := testutil.GetTestAPIKey(t)
	agentBay, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	// Create a session
	fmt.Println("Creating a new session for GetLink testing...")
	sessionParams := agentbay.NewCreateSessionParams().WithImageId("browser_latest")
	sessionResult, err := agentBay.Create(sessionParams)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	// Verify RequestID from Create operation
	if sessionResult.RequestID == "" {
		t.Errorf("Create method did not return RequestID")
	} else {
		t.Logf("Create method successfully returned RequestID: %s", sessionResult.RequestID)
	}

	session := sessionResult.Session
	t.Logf("Session created with ID: %s", session.SessionID)

	defer func() {
		// Clean up the session after test
		fmt.Println("Cleaning up: Deleting the session...")
		deleteResult, err := agentBay.Delete(session)
		if err != nil {
			t.Logf("Warning: Error deleting session: %v", err)
		} else {
			// Verify RequestID from Delete operation
			if deleteResult.RequestID == "" {
				t.Errorf("Delete method did not return RequestID")
			} else {
				t.Logf("Session deleted (RequestID: %s)", deleteResult.RequestID)
			}
		}
	}()

	// Test GetLink method with default parameters (no parameters)
	fmt.Println("Testing session.GetLink method with default parameters...")
	linkResult, err := session.GetLink(nil, nil)
	if err != nil {
		t.Fatalf("Error getting session link: %v", err)
	}

	// Verify RequestID from GetLink operation
	if linkResult.RequestID == "" {
		t.Errorf("GetLink method did not return RequestID")
	} else {
		t.Logf("Session link retrieved (RequestID: %s)", linkResult.RequestID)
	}

	// Verify the link
	if linkResult.Link == "" {
		t.Errorf("Expected non-empty link from GetLink")
	} else {
		t.Logf("Session link: %s", linkResult.Link)
	}

	// Test GetLink method with protocol_type parameter
	fmt.Println("Testing session.GetLink method with protocol_type parameter...")
	protocolType := "https"
	linkWithProtocolResult, err := session.GetLink(&protocolType, nil)
	if err != nil {
		// This error is expected behavior from the backend
		t.Logf("Expected error when using protocol type 'https': %v", err)
	} else {
		// If no error occurs, verify the result
		if linkWithProtocolResult.RequestID == "" {
			t.Errorf("GetLink with protocol type did not return RequestID")
		} else {
			t.Logf("Session link with protocol https retrieved (RequestID: %s)", linkWithProtocolResult.RequestID)
		}
		if linkWithProtocolResult.Link == "" {
			t.Errorf("Expected non-empty link from GetLink with protocol type")
		} else {
			t.Logf("Session link with protocol https: %s", linkWithProtocolResult.Link)
		}
	}

	// Test GetLink method with port parameter
	fmt.Println("Testing session.GetLink method with port parameter...")
	port := int32(8080)
	linkWithPortResult, err := session.GetLink(nil, &port)
	if err != nil {
		t.Errorf("Error getting session link with port: %v", err)
	} else {
		if linkWithPortResult.RequestID == "" {
			t.Errorf("GetLink with port did not return RequestID")
		} else {
			t.Logf("Session link with port 8080 retrieved (RequestID: %s)", linkWithPortResult.RequestID)
		}
		if linkWithPortResult.Link == "" {
			t.Errorf("Expected non-empty link from GetLink with port")
		} else {
			t.Logf("Session link with port 8080: %s", linkWithPortResult.Link)
		}
	}

	// Test GetLink method with both protocol_type and port parameters
	fmt.Println("Testing session.GetLink method with both protocol_type and port parameters...")
	protocolTypeHttps := "https"
	portHttps := int32(443)
	linkWithBothResult, err := session.GetLink(&protocolTypeHttps, &portHttps)
	if err != nil {
		t.Errorf("Error getting session link with protocol and port: %v", err)
	} else {
		if linkWithBothResult.RequestID == "" {
			t.Errorf("GetLink with protocol and port did not return RequestID")
		} else {
			t.Logf("Session link with protocol https and port 443 retrieved (RequestID: %s)", linkWithBothResult.RequestID)
		}
		if linkWithBothResult.Link == "" {
			t.Errorf("Expected non-empty link from GetLink with protocol and port")
		} else {
			t.Logf("Session link with protocol https and port 443: %s", linkWithBothResult.Link)
		}
	}
}

func TestSession_InfoMethod(t *testing.T) {
	// Initialize AgentBay client
	apiKey := testutil.GetTestAPIKey(t)
	agentBay, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	// Create a session
	fmt.Println("Creating a new session for info testing...")
	sessionResult, err := agentBay.Create(nil)
	if err != nil {
		t.Fatalf("Error creating session: %v", err)
	}

	// Verify RequestID from Create operation
	if sessionResult.RequestID == "" {
		t.Errorf("Create method did not return RequestID")
	} else {
		t.Logf("Create method successfully returned RequestID: %s", sessionResult.RequestID)
	}

	session := sessionResult.Session
	t.Logf("Session created with ID: %s", session.SessionID)

	defer func() {
		// Clean up the session after test
		fmt.Println("Cleaning up: Deleting the session...")
		deleteResult, err := agentBay.Delete(session)
		if err != nil {
			t.Logf("Warning: Error deleting session: %v", err)
		} else {
			// Verify RequestID from Delete operation
			if deleteResult.RequestID == "" {
				t.Errorf("Delete method did not return RequestID")
			} else {
				t.Logf("Session deleted (RequestID: %s)", deleteResult.RequestID)
			}
		}
	}()

	// Test Info method
	fmt.Println("Testing session.Info method...")
	infoResult, err := session.Info()
	if err != nil {
		t.Fatalf("Error getting session info: %v", err)
	}

	// Verify RequestID from Info operation
	if infoResult.RequestID == "" {
		t.Errorf("Info method did not return RequestID")
	} else {
		t.Logf("Session info retrieved (RequestID: %s)", infoResult.RequestID)
	}

	// Verify the session info
	if infoResult.Info == nil {
		t.Fatalf("Expected non-nil SessionInfo")
	}

	// Check SessionId field
	if infoResult.Info.SessionId == "" {
		t.Errorf("Expected non-empty SessionId in SessionInfo")
	}
	if infoResult.Info.SessionId != session.SessionID {
		t.Errorf("Expected SessionId to be '%s', got '%s'", session.SessionID, infoResult.Info.SessionId)
	}

	// Check ResourceUrl field
	if infoResult.Info.ResourceUrl == "" {
		t.Errorf("Expected non-empty ResourceUrl in SessionInfo")
	}
	t.Logf("Session ResourceUrl from Info: %s", infoResult.Info.ResourceUrl)

	// Log DesktopInfo fields (these may be empty depending on the API response)
	t.Logf("DesktopInfo - AppId: %s", infoResult.Info.AppId)
	t.Logf("DesktopInfo - AuthCode: %s", infoResult.Info.AuthCode)
	t.Logf("DesktopInfo - ConnectionProperties: %s", infoResult.Info.ConnectionProperties)
	t.Logf("DesktopInfo - ResourceId: %s", infoResult.Info.ResourceId)
	t.Logf("DesktopInfo - ResourceType: %s", infoResult.Info.ResourceType)
	t.Logf("DesktopInfo - Ticket: %s", infoResult.Info.Ticket)

	// Test Ticket field specifically
	if infoResult.Info.Ticket != "" {
		t.Logf("✅ Ticket field is present: %s", infoResult.Info.Ticket)
	} else {
		t.Errorf("ℹ️  Ticket field is empty (this may be normal depending on the API response)")
	}
}

func TestCreate_WithMcpPolicyId_Smoke(t *testing.T) {
	apiKey := testutil.GetTestAPIKey(t)
	client, err := agentbay.NewAgentBay(apiKey)
	if err != nil {
		t.Fatalf("Error initializing AgentBay client: %v", err)
	}

	params := agentbay.NewCreateSessionParams().WithMcpPolicyId("policy-abc")
	_, _ = client.Create(params) // Do not enforce external dependencies; smoke that call path works
}
