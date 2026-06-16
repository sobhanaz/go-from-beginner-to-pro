package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"taskflow/internal/models"
	"taskflow/internal/store"
)

const testSecret = "test-secret"

// newTestServer spins up a Server backed by a real, isolated SQLite DB.
func newTestServer(t *testing.T) *Server {
	t.Helper()
	dbPath := filepath.Join(t.TempDir(), "test.db")
	db, err := store.Open(dbPath)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return NewServer(store.NewTaskStore(db), store.NewUserStore(db), testSecret)
}

// do sends a request, optionally with a Bearer token, and returns the recorder.
func do(t *testing.T, srv *Server, method, path, body, token string) *httptest.ResponseRecorder {
	t.Helper()
	var r *http.Request
	if body == "" {
		r = httptest.NewRequest(method, path, nil)
	} else {
		r = httptest.NewRequest(method, path, strings.NewReader(body))
	}
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	srv.Routes().ServeHTTP(rec, r)
	return rec
}

// registerAndToken creates a user and returns their auth token.
func registerAndToken(t *testing.T, srv *Server, email string) string {
	t.Helper()
	rec := do(t, srv, "POST", "/auth/register",
		`{"email":"`+email+`","password":"secret123"}`, "")
	if rec.Code != http.StatusCreated {
		t.Fatalf("register status = %d; body=%s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Token string `json:"token"`
	}
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Token == "" {
		t.Fatal("no token returned")
	}
	return resp.Token
}

func TestHealthIsPublic(t *testing.T) {
	srv := newTestServer(t)
	if rec := do(t, srv, "GET", "/health", "", ""); rec.Code != http.StatusOK {
		t.Fatalf("status = %d; want 200", rec.Code)
	}
}

func TestTasksRequireAuth(t *testing.T) {
	srv := newTestServer(t)
	// No token -> 401.
	if rec := do(t, srv, "GET", "/tasks", "", ""); rec.Code != http.StatusUnauthorized {
		t.Errorf("no-token status = %d; want 401", rec.Code)
	}
	// Garbage token -> 401.
	if rec := do(t, srv, "GET", "/tasks", "", "not-a-real-token"); rec.Code != http.StatusUnauthorized {
		t.Errorf("bad-token status = %d; want 401", rec.Code)
	}
}

func TestTaskCRUDWithAuth(t *testing.T) {
	srv := newTestServer(t)
	token := registerAndToken(t, srv, "alice@example.com")

	// CREATE
	rec := do(t, srv, "POST", "/tasks", `{"title":"Learn Go"}`, token)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d; body=%s", rec.Code, rec.Body.String())
	}
	var created models.Task
	_ = json.Unmarshal(rec.Body.Bytes(), &created)
	if created.ID == 0 || created.Title != "Learn Go" {
		t.Fatalf("unexpected task: %+v", created)
	}

	// UPDATE
	rec = do(t, srv, "PUT", "/tasks/1", `{"title":"Learn Go well","done":true}`, token)
	if rec.Code != http.StatusOK {
		t.Fatalf("update status = %d", rec.Code)
	}

	// DELETE
	rec = do(t, srv, "DELETE", "/tasks/1", "", token)
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d; want 204", rec.Code)
	}
}

// Tasks must be isolated per user: Bob cannot see Alice's task.
func TestTasksAreUserScoped(t *testing.T) {
	srv := newTestServer(t)
	aliceToken := registerAndToken(t, srv, "alice@example.com")
	bobToken := registerAndToken(t, srv, "bob@example.com")

	// Alice creates a task (it gets id 1).
	do(t, srv, "POST", "/tasks", `{"title":"Alice's secret"}`, aliceToken)

	// Bob's list must be empty.
	rec := do(t, srv, "GET", "/tasks", "", bobToken)
	var bobTasks []models.Task
	_ = json.Unmarshal(rec.Body.Bytes(), &bobTasks)
	if len(bobTasks) != 0 {
		t.Errorf("bob sees %d tasks; want 0", len(bobTasks))
	}

	// Bob cannot fetch Alice's task by id -> 404.
	rec = do(t, srv, "GET", "/tasks/1", "", bobToken)
	if rec.Code != http.StatusNotFound {
		t.Errorf("bob get alice's task status = %d; want 404", rec.Code)
	}
}

func TestLoginFlow(t *testing.T) {
	srv := newTestServer(t)
	registerAndToken(t, srv, "carol@example.com")

	// Correct password -> 200 + token.
	rec := do(t, srv, "POST", "/auth/login",
		`{"email":"carol@example.com","password":"secret123"}`, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("login status = %d", rec.Code)
	}

	// Wrong password -> 401.
	rec = do(t, srv, "POST", "/auth/login",
		`{"email":"carol@example.com","password":"wrong"}`, "")
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("wrong-password status = %d; want 401", rec.Code)
	}
}

func TestDuplicateEmail(t *testing.T) {
	srv := newTestServer(t)
	registerAndToken(t, srv, "dave@example.com")
	// Registering the same email again -> 409 Conflict.
	rec := do(t, srv, "POST", "/auth/register",
		`{"email":"dave@example.com","password":"secret123"}`, "")
	if rec.Code != http.StatusConflict {
		t.Errorf("duplicate status = %d; want 409", rec.Code)
	}
}

// --- Priority feature tests ---

// A task created without a priority defaults to "medium".
func TestCreateTaskDefaultsPriority(t *testing.T) {
	srv := newTestServer(t)
	token := registerAndToken(t, srv, "p1@example.com")

	rec := do(t, srv, "POST", "/tasks", `{"title":"no priority given"}`, token)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d; body=%s", rec.Code, rec.Body.String())
	}
	var task models.Task
	_ = json.Unmarshal(rec.Body.Bytes(), &task)
	if task.Priority != models.PriorityMedium {
		t.Errorf("default priority = %q; want %q", task.Priority, models.PriorityMedium)
	}
}

// A valid priority is stored and returned.
func TestCreateTaskWithPriority(t *testing.T) {
	srv := newTestServer(t)
	token := registerAndToken(t, srv, "p2@example.com")

	rec := do(t, srv, "POST", "/tasks", `{"title":"urgent","priority":"high"}`, token)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d", rec.Code)
	}
	var task models.Task
	_ = json.Unmarshal(rec.Body.Bytes(), &task)
	if task.Priority != models.PriorityHigh {
		t.Errorf("priority = %q; want %q", task.Priority, models.PriorityHigh)
	}
}

// An invalid priority is rejected with 400.
func TestCreateTaskInvalidPriority(t *testing.T) {
	srv := newTestServer(t)
	token := registerAndToken(t, srv, "p3@example.com")

	rec := do(t, srv, "POST", "/tasks", `{"title":"x","priority":"urgent"}`, token)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("invalid-priority status = %d; want 400", rec.Code)
	}
}

// ?priority= and ?done= filters narrow the list correctly.
func TestListTaskFilters(t *testing.T) {
	srv := newTestServer(t)
	token := registerAndToken(t, srv, "p4@example.com")

	// Seed: one high (will be marked done), one low (stays not-done).
	do(t, srv, "POST", "/tasks", `{"title":"high task","priority":"high"}`, token)
	do(t, srv, "POST", "/tasks", `{"title":"low task","priority":"low"}`, token)
	// Mark task id 1 (high) as done.
	do(t, srv, "PUT", "/tasks/1", `{"title":"high task","priority":"high","done":true}`, token)

	listLen := func(query string) int {
		rec := do(t, srv, "GET", "/tasks"+query, "", token)
		if rec.Code != http.StatusOK {
			t.Fatalf("list%s status = %d", query, rec.Code)
		}
		var tasks []models.Task
		_ = json.Unmarshal(rec.Body.Bytes(), &tasks)
		return len(tasks)
	}

	if n := listLen(""); n != 2 {
		t.Errorf("no filter: got %d; want 2", n)
	}
	if n := listLen("?priority=high"); n != 1 {
		t.Errorf("priority=high: got %d; want 1", n)
	}
	if n := listLen("?done=true"); n != 1 {
		t.Errorf("done=true: got %d; want 1", n)
	}
	if n := listLen("?done=false"); n != 1 {
		t.Errorf("done=false: got %d; want 1", n)
	}
	if n := listLen("?priority=low&done=false"); n != 1 {
		t.Errorf("combined filter: got %d; want 1", n)
	}

	// A bad filter value is a 400.
	if rec := do(t, srv, "GET", "/tasks?priority=bogus", "", token); rec.Code != http.StatusBadRequest {
		t.Errorf("bad priority filter status = %d; want 400", rec.Code)
	}
}
