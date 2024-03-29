package text2image

import "testing"

func TestImageSynthesis(t *testing.T) {
	imageRequest := &ImageRequest{}
	imageRequest.Model = "stable-diffusion-xl"
	imageRequest.Input.Prompt = "a photograph of an astronaut riding a horse"
	imageRequest.Parameters.Size = "1024*1024"

	apiKey := "sk-xxxxxxx"

	resp, err := imageRequest.ImageSynthesis(apiKey)
	if err != nil {
		t.Errorf("image synthesis failed: %v", err)
	}

	t.Log(resp)
}
