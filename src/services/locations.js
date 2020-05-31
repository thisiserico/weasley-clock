const locationsAddress = "/.netlify/functions/locations"

export const fetchAll = () => {
  return fetch(locationsAddress)
    .then(response => response.json())
}
