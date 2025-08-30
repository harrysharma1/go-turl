package urlshortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const uid = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGen(t *testing.T) {
	initialLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := ShortLinkGeneration(initialLink1, uid)

	initialLink2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink2 := ShortLinkGeneration(initialLink2, uid)

	initialLink3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink3 := ShortLinkGeneration(initialLink3, uid)
	assert.Equal(t, shortLink1, "2xdutAWF")
	assert.Equal(t, shortLink2, "oc52ubJ4")
	assert.Equal(t, shortLink3, "GY6jpQ7N")
}
