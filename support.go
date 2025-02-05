package support

var Ipsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque imperdiet libero eu neque facilisis, ac pretium nisi dignissim. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.\n"

// GetIpsum returns the provided number of paragraphs of 'Lorem ipsum' text in byte form.
func GetIpsum(paragraphs int) []byte {
	output := Ipsum
	for i := 0; i < paragraphs; i++ {
		output += Ipsum
	}
	return []byte(output)
}
