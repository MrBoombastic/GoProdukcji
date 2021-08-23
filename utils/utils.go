package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BOOMfinity-Developers/bfcord/discord"
	"github.com/patrickmn/go-cache"
	"goprodukcji/config"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var requestsCache = cache.New(5*time.Minute, 10*time.Minute)

func FetchArticles(options string, caching bool) (Articles, error) {
	var articles Articles

	url := "https://naprodukcji.xyz/ghost/api/v3/content/posts/?key=" + config.GetConfig().GhostToken + options
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return Articles{}, err
	}
	err = json.NewDecoder(res.Body).Decode(&articles)
	if err != nil {
		log.Fatal(err)
		return Articles{}, err
	}
	if caching {
		requestsCache.Set("GetArticles", articles, cache.DefaultExpiration)
	}
	return articles, err
}

func GetArticles(options string, caching bool) (Articles, error) {
	cacheRes, found := requestsCache.Get("GetArticles")
	if options == "all" {
		options = "&limit=all&fields=id,title,url,primary_author,excerpt,published_at,feature_image&order=published_at%20desc&formats=plaintext&include=authors"
	} else if options == "id" {
		options = "&limit=1&fields=id,url"
	} else {
		return Articles{}, errors.New("bad options")
	}
	if found && caching {
		cachedArticlees := cacheRes.(Articles)
		if len(cachedArticlees.Posts) > 0 {
			return cachedArticlees, nil
		} else {
			articles, err := FetchArticles(options, caching)
			if err != nil {
				return Articles{}, err
			} else {
				return articles, nil
			}
		}
	} else {
		articles, err := FetchArticles(options, caching)
		if err != nil {
			return Articles{}, err
		} else {
			return articles, nil
		}
	}
}

func SearchArticle(query string) (Article, error) {
	articles, err := GetArticles("all", true)
	if err != nil {
		return Article{}, err
	}
	for i := range articles.Posts {
		if strings.Contains(strings.ToLower(articles.Posts[i].Title), strings.ToLower(query)) {
			return articles.Posts[i], nil
		}
	}
	return Article{}, errors.New(fmt.Sprintf("artykuł nie został znaleziony! Liczba artykułów w cache: %v", len(articles.Posts)))
}

func FormatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func GetMemory() (x uint64) {
	data, err := os.ReadFile("/proc/self/statm")
	if err != nil {
		return 0
	}
	d := bytes.Split(data, []byte(" "))

	r, _ := strconv.Atoi(string(d[1]))
	x += uint64(r)
	r, _ = strconv.Atoi(string(d[2]))
	x += uint64(r)
	x = x * 1024
	return
}

func MentionEmbed(config config.RunMode, guildIcon string) discord.MessageEmbed {
	return discord.MessageEmbed{
		Title: "Witaj!",
		Color: 0xff8000,
		Description: fmt.Sprintf("Jestem botem zaprojektowanym specjalnie dla serwera Na Produkcji!\n"+
			"Mój prefix to `%v`. Pomoc znajdziesz w `%vhelp`.\n"+
			"Kod źródłowy znajdziesz [tutaj](https://github.com/MrBoombastic/GoProdukcji).", config.Prefix, config.Prefix),
		Thumbnail: &discord.EmbedMedia{
			Url: guildIcon,
		},
		URL: "https://naprodukcji.xyz",
		Image: &discord.EmbedMedia{
			Url: "https://naprodukcji.xyz/content/images/2021/06/comment_1622802543Quw49Z60cINC7fttv0aBcp.jpg",
		},
	}
}
