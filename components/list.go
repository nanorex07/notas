package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	normalTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).
			Padding(0, 0, 0, 2) //nolint:mnd

	normalDesc = normalTitle.
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})

	normalTag = normalDesc.
			Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}).
			Italic(true)

	selectedTitle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}).
			Padding(0, 0, 0, 1)

	selectedDesc = selectedTitle.
			Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})

	selectedTag = selectedDesc.
			Foreground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
			Italic(true)
)

type ListItem struct {
	ID    string
	Title string
	Body  string
	Tags  []string
}

type ListComponent struct {
	selectedItem   int
	items          []*ListItem
	maxBodyLength  int
	maxTitleLength int
}

func NewListComponent() *ListComponent {
	return &ListComponent{
		selectedItem:   0,
		items:          []*ListItem{},
		maxBodyLength:  25,
		maxTitleLength: 10,
	}
}

func (item *ListItem) render(selected bool, maxBodyLength int, maxTitleLength int) string {
	tagString := "# " + strings.Join(item.Tags, ", ")
	body := item.Body
	if len(body) > maxBodyLength {
		body = body[:maxBodyLength-4] + " ..."
	}
	title := item.Title
	if len(title) > maxTitleLength {
		title = title[:maxTitleLength-4] + " ..."
	}
	if len(tagString) > maxBodyLength {
		tagString = tagString[:maxBodyLength-4] + " ..."
	}
	if selected {
		return selectedTitle.Render(title) + "\n" + selectedDesc.Render(body) + "\n" + selectedTag.Render(tagString)
	}
	return normalTitle.Render(title) + "\n" + normalDesc.Render(body) + "\n" + normalTag.Render(tagString)
}

func (c *ListComponent) AddItem(item *ListItem) {
	c.items = append(c.items, item)
}

func (c *ListComponent) Clear() {
	c.items = []*ListItem{}
}

func (c *ListComponent) SelectedItem() *ListItem {
	return c.items[c.selectedItem]
}

func (c *ListComponent) Up() {
	c.selectedItem = (c.selectedItem - 1 + len(c.items)) % len(c.items)
}

func (c *ListComponent) Down() {
	c.selectedItem = (c.selectedItem + 1) % len(c.items)
}

func (c *ListComponent) Len() int {
	return len(c.items)
}

func (c *ListComponent) Render() string {
	render := ""
	for i, item := range c.items {
		render += item.render(i == c.selectedItem, c.maxBodyLength, c.maxTitleLength)
		render += "\n\n"
	}
	return render
}
