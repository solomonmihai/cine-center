#!/usr/local/bin/python3

from parsel import Selector
from datetime import datetime
import requests
import re
import json

EVENTBOOK_URL = "https://eventbook.ro"
# TODO: maybe the cinema name could also be retrieved from the page
CINEMAS = [
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
    },
    {
        "name": "Cinema Muzeul Taranului Roman",
        "url": "https://eventbook.ro/hall/cinema-muzeul-taranului-studio-horia-bernea"
    }
]


def get_page_selector(url: str):
    res = requests.get(url)
    return Selector(res.text)

# TODO: get the number from the last element because
# ... in the middle when many pages


def get_page_count(selector: Selector):
    pages_el = selector.css(".page-item").extract()
    count = len(pages_el[1:-1])

    # at least 1 page
    return max(count, 1)


def get_text_from_el(text: str):
    return " ".join([t.strip() for t in re.findall(r"<[^>]+>|[^<]+", text) if not "<" in t]).strip()


def get_film_details(perf: Selector):
    left_col = perf.css(".col-12.col-md-3")
    right_col = perf.css(".event-buy-tickets")

    title = right_col.css("h5::text").get().strip()

    date_el_text = left_col.css("div>h4").get().strip()
    date_text = " ".join(get_text_from_el(date_el_text).split())

    date = ""
    try:
        date = int(datetime.strptime(date_text, "%d %b %Y %H:%M").timestamp())
    except:
        return None

    link = EVENTBOOK_URL + left_col.css("a").attrib["href"]
    img_url = left_col.css("a>div>img").attrib["src"]

    price = "???"
    try:
        price_el_text = right_col.css(
            ".col-12.text-dark.text-center>.text-uppercase").get().strip()
        price = get_text_from_el(price_el_text).split(": ")[1]
    except:
        pass

    subtitles = right_col.css("h6").getall()
    sub_texts = [get_text_from_el(sub) for sub in subtitles]
    location = " // ".join(
        [t for t in sub_texts if "Add to cart" not in t and t.strip() != ""])

    return {
        "title": title,
        "date": date,
        "price": price,
        "link": link,
        "img_url": img_url,
        "location": location,
    }


def get_films_from_page(url: str):
    selector = get_page_selector(url)
    performances = selector.css("#performance")

    return [
        film for perf in performances
        if (film := get_film_details(perf)) is not None
    ]


def scrape_cinema(url: str):
    selector = get_page_selector(url)
    page_count = get_page_count(selector)

    films = []

    # TODO: films from the first page can be retrieved from the selector
    # used to get the page count
    # no need for another request
    for page_no in range(1, page_count + 1):
        page_url = f"{url}?page={page_no}"
        films.extend(get_films_from_page(page_url))

    return films


if __name__ == "__main__":
    data = {}

    for cinema in CINEMAS:
        print(' - ' + cinema["name"])
        films_data = scrape_cinema(cinema["url"])
        data[cinema["name"]] = {
            **cinema,
            "films": films_data
        }

    with open("../data.json", "w") as f:
        json.dump(data, f, indent=2)
