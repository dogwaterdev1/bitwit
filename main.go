package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/cursor"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var (
// 	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
// 	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
// 	cursorStyle         = focusedStyle
// 	noStyle             = lipgloss.NewStyle()
// 	helpStyle           = blurredStyle
// 	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

// 	focusedButton = focusedStyle.Render("[ Submit ]")
// 	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
// )

// type model struct {
// 	focusIndex int
// 	inputs     []textinput.Model
// 	cursorMode cursor.Mode
// 	apiResult  string
// }

// func initialModel() model {
// 	m := model{
// 		inputs: make([]textinput.Model, 4),
// 	}

// 	var t textinput.Model
// 	for i := range m.inputs {
// 		t = textinput.New()
// 		t.Cursor.Style = cursorStyle
// 		t.CharLimit = 32

// 		switch i {
// 		case 0:
// 			t.Placeholder = "Username"
// 			t.Focus()
// 			t.PromptStyle = focusedStyle
// 			t.TextStyle = focusedStyle
// 		case 1:
// 			t.Placeholder = "Email"
// 			t.CharLimit = 64
// 		case 2:
// 			t.Placeholder = "Password"
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '•'
// 		case 3:
// 			t.Placeholder = "ASDF"
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '•'
// 		}

// 		m.inputs[i] = t
// 	}

// 	return m
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "esc":
// 			return m, tea.Quit

// 		// Change cursor mode
// 		case "ctrl+r":
// 			m.cursorMode++
// 			if m.cursorMode > cursor.CursorHide {
// 				m.cursorMode = cursor.CursorBlink
// 			}
// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := range m.inputs {
// 				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
// 			}
// 			return m, tea.Batch(cmds...)

// 		// Set focus to next input
// 		case "tab", "shift+tab", "enter", "up", "down":
// 			s := msg.String()

// 			// Did the user press enter while the submit button was focused?
// 			// If so, exit.
// 			if s == "enter" && m.focusIndex == len(m.inputs) {
// 				return m, tea.Quit
// 			}

// 			// Cycle indexes
// 			if s == "up" || s == "shift+tab" {
// 				m.focusIndex--
// 			} else {
// 				m.focusIndex++
// 			}

// 			if m.focusIndex > len(m.inputs) {
// 				m.focusIndex = 0
// 			} else if m.focusIndex < 0 {
// 				m.focusIndex = len(m.inputs)
// 			}

// 			cmds := make([]tea.Cmd, len(m.inputs))
// 			for i := 0; i <= len(m.inputs)-1; i++ {
// 				if i == m.focusIndex {
// 					// Set focused state
// 					cmds[i] = m.inputs[i].Focus()
// 					m.inputs[i].PromptStyle = focusedStyle
// 					m.inputs[i].TextStyle = focusedStyle
// 					continue
// 				}
// 				// Remove focused state
// 				m.inputs[i].Blur()
// 				m.inputs[i].PromptStyle = noStyle
// 				m.inputs[i].TextStyle = noStyle
// 			}

// 			return m, tea.Batch(cmds...)
// 		// execute api call
// 		case "0":
// 			s := msg.String()
// 			// URL to send GET request to (using a public API for demonstration)
// 			url := "https://api.ipify.org?format=json"

// 			// Perform the HTTP GET request
// 			resp, err := http.Get(url)
// 			if err != nil {
// 				s += fmt.Sprintf("Error making GET request: %v\n", err)
// 			}
// 			defer resp.Body.Close()

// 			// Read the response body
// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				s += fmt.Sprintf("Error reading response: %v\n", err)
// 			}

// 			//fmt.Printf("Response Status: %s\n", resp.Status)
// 			s += fmt.Sprintf("Response Status: %s\n", resp.Status)
// 			//fmt.Printf("Response Body: %s\n", string(body))
// 			s += fmt.Sprintf("Response Body: %v\n", string(body))
// 		}

// 	}

// 	// Handle character input and blinking
// 	cmd := m.updateInputs(msg)

// 	return m, cmd
// }

// func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
// 	cmds := make([]tea.Cmd, len(m.inputs))

// 	// Only text inputs with Focus() set will respond, so it's safe to simply
// 	// update all of them here without any further logic.
// 	for i := range m.inputs {
// 		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
// 	}

// 	return tea.Batch(cmds...)
// }

// func (m model) View() string {
// 	var b strings.Builder

// 	for i := range m.inputs {
// 		b.WriteString(m.inputs[i].View())
// 		if i < len(m.inputs)-1 {
// 			b.WriteRune('\n')
// 		}
// 	}

// 	button := &blurredButton
// 	if m.focusIndex == len(m.inputs) {
// 		button = &focusedButton
// 	}
// 	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

// 	b.WriteString(helpStyle.Render("cursor mode is "))
// 	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
// 	b.WriteString(helpStyle.Render(" (ctrl+r to change style)\n"))
// 	b.WriteString(helpStyle.Render(m.apiResult))

// 	return b.String()
// }

// func main() {
// 	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
// 		fmt.Printf("could not start program: %s\n", err)
// 		os.Exit(1)
// 	}
// }

func main() {

	// read secrets file for xai key
	fmt.Println("Hello")
	content, err := os.ReadFile("secrets.txt")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := strings.TrimSpace(string(content))

	// set url
	url := "https://api.x.ai/v1/chat/completions"

	// Create request payload
	payload := `{
		"messages": [
			{
				"role": "user",
				"content": "What is the latest crypto news today? And what is today's date? Summarize it all into about 1-2 paragraphs of information"
			}
		],
		"model": "grok-3-latest",
		"stream": false,
		"search_parameters": { 
			"mode":"on",
			"sources": [
  				{"type": "web"},
  				{"type": "x"},
 				{"type": "news"}
			]
		},
		"temperature": 0.7
	}`

	// create new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Print the response status and body
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", string(body))

	// parse data into a map
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
	}

	jsonPrettyString, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error pretty-printing JSON: %v\n", err)
		return
	}
	fmt.Printf("Map as pretty JSON string:\n%s\n", string(jsonPrettyString))
	fmt.Printf("\n\n\n%s\n", strings.Split(string(jsonPrettyString), "reasoning_content")[0])
	// old access to content field
	choices, ok := data["choices"]
	if !ok {
		fmt.Println("Error: 'choices' key not found in data")
		return
	}
	// val := strings.Split(string(choicesStr), "reasoning_content")
	fmt.Printf("\n\n\nContent: %s\n", choices)

}
