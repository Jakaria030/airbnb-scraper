package config

const BASE_URL = "https://www.airbnb.com"
const TIMEOUT = 60
const FILE_PATH = "data/properties.csv"
const CONNECTION_STRING = "postgres://scraperuser:scraperpass@localhost:5432/scraperdb"

var SEARCH_URLS = [...]string{"https://www.airbnb.com/s/Kuala-Lumpur/homes?place_id=ChIJ5-rvAcdJzDERfSgcL1uO2fQ&refinement_paths%5B%5D=%2Fhomes&flexible_trip_lengths%5B%5D=weekend_trip&date_picker_type=FLEXIBLE_DATES&search_type=HOMEPAGE_CAROUSEL_CLICK", "https://www.airbnb.com/s/Kuala-Lumpur/homes?place_id=ChIJ5-rvAcdJzDERfSgcL1uO2fQ&refinement_paths%5B%5D=%2Fhomes&flexible_trip_lengths%5B%5D=weekend_trip&date_picker_type=flexible_dates&query=Kuala%20Lumpur&monthly_start_date=2026-03-01&monthly_length=3&monthly_end_date=2026-06-01&price_filter_input_type=2&price_filter_num_nights=2&channel=EXPLORE&federated_search_session_id=da2d7532-742d-4e0e-9df7-3ce4563d7b4f&pagination_search=true&cursor=eyJzZWN0aW9uX29mZnNldCI6MCwiaXRlbXNfb2Zmc2V0IjoxOCwidmVyc2lvbiI6MX0%3D"}
