from parsel import Selector
import requests
import re
import json

EVENTBOOK_URL = "https://eventbook.ro"

cinemas = [
  {
    "name": "Cinema Elvire Popesco",
    "url": "https://eventbook.ro/hall/cinema-elvire-popesco",
  },
  {
    "name": "Cinema Europa",
    "url": "https://eventbook.ro/hall/cinema-europa-bucuresti"
  },
  {
    "name": "Cinema Eforie",
    "url": "https://eventbook.ro/hall/cinema-eforie",
  },
  {
    "name": "Cinema Union",
    "url": "https://eventbook.ro/hall/cinema-union",
  }
]

def get_page_selector(url: str):
  res = requests.get(url)
  html = res.content
  text = html.decode()
  return Selector(text)

def get_page_count(selector: Selector):
  pages_el = selector.css(".page-item").extract()
  count = len(pages_el[1:-1])
  return count

def get_text_from_el(text: str):
  return " ".join([t.strip() for t in re.findall(r"<[^>]+>|[^<]+", text) if not "<" in t])

def get_film_details(perf: Selector):
  left_col = perf.css(".col-12.col-md-3");
  right_col = perf.css(".event-buy-tickets")

  title = right_col.css("h5::text").get().strip()

  date_el_text = left_col.css("div>h4").get().strip()
  date = get_text_from_el(date_el_text)

  link = EVENTBOOK_URL + left_col.css("a").attrib["href"]
  img_url = left_col.css("a>div>img").attrib["src"]

  price_el_text = right_col.css(".col-12.text-dark.text-center>.text-uppercase").get().strip()
  price = get_text_from_el(price_el_text).split(": ")[1]

  return {
    "title": title,
    "date": date,
    "price": price,
    "link": link,
    "img_url": img_url,
  }

def get_films_from_page(url: str):
  selector = get_page_selector(url)
  performances = selector.css("#performance")

  return [get_film_details(perf) for perf in performances]

def scrape(url: str):
  selector = get_page_selector(url)
  page_count = get_page_count(selector)

  films = []

  if page_count == 0:
    page_count = 1

  for page_no in range(1, page_count + 1):
    page_url = url + "?page=" + str(page_no)
    films.append(get_films_from_page(page_url))

  return films

if __name__ == "__main__":
  for cinema in cinemas:
    film_data = scrape(cinema["url"])
    cinema["films"] = film_data

  with open("data.json", "w") as f:
    json.dump(cinemas, f, indent=2)


