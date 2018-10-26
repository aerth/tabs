package tabs

// TabJSON is returned by the tab service we provide
type TabJSON struct {
	Data struct {
		TabView struct {
			WikiTab struct {
				Content string `json:"content"`
			} `json:"wiki_tab"`
		}
	}
}

// String returns only the tab as string
func (t TabJSON) String() string {
	return t.Data.TabView.WikiTab.Content
}

// Bytes returns only the tab as bytes
func (t TabJSON) Bytes() []byte {
	return []byte(t.String())
}
