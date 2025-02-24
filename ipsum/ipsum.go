// Package ipsum provides some convenience methods around creating strings of Lorem Ipsum.
package ipsum

// Paragraph represents one constant paragraph of Lorem Ipsum text.
var Paragraph = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque imperdiet libero eu neque facilisis, ac pretium nisi dignissim. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.\n"

// Generate returns paragraphs of 'Lorem ipsum' text.
func Generate(paragraphs ...int) []byte {
	count := 1
	if len(paragraphs) > 0 {
		count = paragraphs[0]
	}

	output := Paragraph
	for i := 0; i < count; i++ {
		output += Paragraph
	}
	return []byte(output)
}
