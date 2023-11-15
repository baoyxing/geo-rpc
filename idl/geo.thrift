namespace go geo

struct Location {
    1: string country
    2: string region
    3: string city
    4: string isp
}


service GeoService {
 Location GetLocation()
}