// hooks_cursor_posttodo.go contains the PostTodo hook handler for Cursor.
// This reuses the same incremental checkpoint logic as Claude Code's PostTodo handler.
package cli

// handleCursorPostTodo handles the PostToolUse[TodoWrite] hook for Cursor.
// Reuses the same incremental checkpoint logic as Claude Code.
func handleCursorPostTodo() error {
	return handleClaudeCodePostTodo()
}
